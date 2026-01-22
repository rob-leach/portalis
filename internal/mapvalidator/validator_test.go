package mapvalidator

import (
	"testing"
)

func TestDirectionDeltas(t *testing.T) {
	// Verify all directions have deltas defined
	expectedDirs := []string{
		"north", "south", "east", "west",
		"up", "down",
		"northeast", "northwest", "southeast", "southwest",
	}

	for _, dir := range expectedDirs {
		if _, ok := DirectionDelta[dir]; !ok {
			t.Errorf("Missing delta for direction: %s", dir)
		}
		if _, ok := InverseDirection[dir]; !ok {
			t.Errorf("Missing inverse for direction: %s", dir)
		}
	}
}

func TestCoordAddition(t *testing.T) {
	origin := Coord{0, 0, 0}

	tests := []struct {
		name     string
		delta    Coord
		expected Coord
	}{
		{"north", DirectionDelta["north"], Coord{0, 1, 0}},
		{"south", DirectionDelta["south"], Coord{0, -1, 0}},
		{"east", DirectionDelta["east"], Coord{1, 0, 0}},
		{"west", DirectionDelta["west"], Coord{-1, 0, 0}},
		{"northeast", DirectionDelta["northeast"], Coord{1, 1, 0}},
		{"southwest", DirectionDelta["southwest"], Coord{-1, -1, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := origin.Add(tc.delta)
			if !result.Equals(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestBidirectionalValidation(t *testing.T) {
	// Create a simple valid map
	rooms := map[int]*RoomInfo{
		100: {RoomID: 100, Zone: "Test", Exits: map[string]int{"north": 101, "east": 102}},
		101: {RoomID: 101, Zone: "Test", Exits: map[string]int{"south": 100}},
		102: {RoomID: 102, Zone: "Test", Exits: map[string]int{"west": 100}},
	}

	v := NewValidator(rooms)
	result := v.Validate()

	if len(result.Errors) > 0 {
		t.Errorf("Expected no errors for valid bidirectional map, got: %v", result.Errors)
	}
}

func TestMissingInverseExit(t *testing.T) {
	// Room 101 has north->102, but 102 has no exit back
	rooms := map[int]*RoomInfo{
		100: {RoomID: 100, Zone: "Test", Exits: map[string]int{"north": 101}},
		101: {RoomID: 101, Zone: "Test", Exits: map[string]int{"south": 100, "north": 102}},
		102: {RoomID: 102, Zone: "Test", Exits: map[string]int{}}, // Missing south exit to 101
	}

	v := NewValidator(rooms)
	result := v.Validate()

	// Should have an error about missing inverse exit
	foundMissingInverse := false
	for _, err := range result.Errors {
		if err.RoomID == 101 && containsString(err.Message, "missing inverse") {
			foundMissingInverse = true
			break
		}
	}

	if !foundMissingInverse {
		t.Errorf("Expected error about missing inverse exit from 102 to 101")
	}
}

func TestCoordinateCollision(t *testing.T) {
	// Create a map where room 102 ends up at the same coordinate as room 101
	// This simulates the bug where two rooms occupy the same space
	rooms := map[int]*RoomInfo{
		100: {RoomID: 100, Zone: "Test", Exits: map[string]int{"north": 101, "east": 102}},
		101: {RoomID: 101, Zone: "Test", Exits: map[string]int{"south": 100, "west": 102}},
		102: {RoomID: 102, Zone: "Test", Exits: map[string]int{"west": 100, "east": 101}},
	}

	v := NewValidator(rooms)
	result := v.Validate()

	// Should detect the geometry problem
	hasGeometryError := len(result.Errors) > 0
	if !hasGeometryError {
		t.Errorf("Expected geometry errors for invalid coordinate map")
	}
}

func TestKnown103To600Bug(t *testing.T) {
	// Recreate the documented bug:
	// - Room 100 at (0,0) has west->600
	// - Room 600 at (-1,0) has east->100 (correct)
	// - If room 103 at (1,0) had east->600, that would be wrong because
	//   east from (1,0) should be (2,0), not (-1,0)

	rooms := map[int]*RoomInfo{
		100: {RoomID: 100, Zone: "Starter Town", Exits: map[string]int{"west": 600, "east": 103}},
		103: {RoomID: 103, Zone: "Starter Town", Exits: map[string]int{"west": 100, "east": 600}}, // BUG: east->600 is wrong
		600: {RoomID: 600, Zone: "Bladeworks", Exits: map[string]int{"east": 100, "west": 103}},   // Tries to connect to 103 but coords don't match
	}

	v := NewValidator(rooms)
	result := v.Validate()

	// Should have multiple errors about the impossible geometry
	if len(result.Errors) == 0 {
		t.Errorf("Expected errors for the 103->600 bug scenario")
	}

	// Specifically should find that east from 103 doesn't reach 600 correctly
	foundDeltaError := false
	for _, err := range result.Errors {
		if err.RoomID == 103 && containsString(err.Message, "expected target at") {
			foundDeltaError = true
			break
		}
	}

	if !foundDeltaError {
		t.Errorf("Expected direction delta error for room 103's east exit to 600")
	}
}

func TestProposalValidation(t *testing.T) {
	// Existing world with one room at (0,0)
	rooms := map[int]*RoomInfo{
		100: {RoomID: 100, Zone: "Starter Town", Exits: map[string]int{"south": 500}},
	}

	v := NewValidator(rooms)
	v.Validate() // Compute coordinates first

	// Proposal that correctly connects to room 100
	proposal := &Proposal{
		Zone: "New Zone",
		Rooms: []ProposalRoom{
			{ID: 500, Coord: [3]int{0, -1, 0}, Exits: map[string]int{"north": 100}},
		},
	}

	propResult := v.ValidateProposal(proposal)

	if len(propResult.Errors) > 0 {
		t.Errorf("Expected no errors for valid proposal, got: %v", propResult.Errors)
	}
}

func TestProposalCoordinateCollision(t *testing.T) {
	// Existing world
	rooms := map[int]*RoomInfo{
		100: {RoomID: 100, Zone: "Starter Town", Exits: map[string]int{"south": 500}},
		500: {RoomID: 500, Zone: "Caves", Exits: map[string]int{"north": 100}},
	}

	v := NewValidator(rooms)
	v.Validate() // Compute coordinates

	// Proposal that would collide with existing room 500 at (0,-1)
	proposal := &Proposal{
		Zone: "Bad Zone",
		Rooms: []ProposalRoom{
			{ID: 700, Coord: [3]int{0, -1, 0}, Exits: map[string]int{}}, // Same coord as 500
		},
	}

	propResult := v.ValidateProposal(proposal)

	foundCollision := false
	for _, err := range propResult.Errors {
		if containsString(err.Message, "collide") {
			foundCollision = true
			break
		}
	}

	if !foundCollision {
		t.Errorf("Expected collision error for proposal at existing coordinate")
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstr(s, substr))
}

func containsSubstr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// TestActualWorld runs validation against the actual world files
// This is an integration test that should be run from the repo root
func TestActualWorld(t *testing.T) {
	roomsDir := "../../_datafiles/world/default/rooms"
	rooms, err := ParseRoomsFromDirectory(roomsDir)
	if err != nil {
		t.Skipf("Could not load rooms from %s: %v (run from repo root)", roomsDir, err)
	}

	t.Logf("Loaded %d rooms", len(rooms))

	v := NewValidator(rooms)
	result := v.Validate()

	// Log all errors and warnings for inspection
	for _, e := range result.Errors {
		t.Logf("[ERROR] Room %d: %s", e.RoomID, e.Message)
	}
	for _, w := range result.Warnings {
		t.Logf("[WARN]  Room %d: %s", w.RoomID, w.Message)
	}

	t.Logf("Summary: %d errors, %d warnings", len(result.Errors), len(result.Warnings))
}
