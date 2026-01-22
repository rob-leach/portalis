package mapvalidator

import (
	"fmt"
	"sort"
)

// Validator performs map validation
type Validator struct {
	rooms     map[int]*RoomInfo
	errors    []ValidationError
	warnings  []ValidationError
	originID  int // Room ID of the origin (Town Square = 100)
}

// NewValidator creates a new validator with the given rooms
func NewValidator(rooms map[int]*RoomInfo) *Validator {
	return &Validator{
		rooms:    rooms,
		errors:   []ValidationError{},
		warnings: []ValidationError{},
		originID: 100, // Town Square is always the origin
	}
}

// SetOrigin sets the origin room ID (default is 100)
func (v *Validator) SetOrigin(roomID int) {
	v.originID = roomID
}

// addError adds an error to the validation results
func (v *Validator) addError(roomID int, format string, args ...interface{}) {
	v.errors = append(v.errors, ValidationError{
		Severity: "ERROR",
		RoomID:   roomID,
		Message:  fmt.Sprintf(format, args...),
	})
}

// addWarning adds a warning to the validation results
func (v *Validator) addWarning(roomID int, format string, args ...interface{}) {
	v.warnings = append(v.warnings, ValidationError{
		Severity: "WARN",
		RoomID:   roomID,
		Message:  fmt.Sprintf(format, args...),
	})
}

// Validate runs all validation checks and returns the results
func (v *Validator) Validate() *ValidationResult {
	// Step 1: Compute coordinates starting from origin
	v.computeCoordinates()

	// Step 2: Check for coordinate collisions
	v.checkCoordinateCollisions()

	// Step 3: Check bidirectional consistency
	v.checkBidirectionalConsistency()

	// Step 4: Check direction deltas
	v.checkDirectionDeltas()

	return &ValidationResult{
		Errors:   v.errors,
		Warnings: v.warnings,
	}
}

// computeCoordinates uses BFS from origin to assign coordinates to all reachable rooms
func (v *Validator) computeCoordinates() {
	origin, ok := v.rooms[v.originID]
	if !ok {
		v.addError(v.originID, "Origin room %d not found", v.originID)
		return
	}

	// Set origin at (0, 0, 0)
	origin.Coord = &Coord{X: 0, Y: 0, Z: 0}

	// BFS queue
	queue := []int{v.originID}
	visited := make(map[int]bool)
	visited[v.originID] = true

	for len(queue) > 0 {
		currentID := queue[0]
		queue = queue[1:]

		current, ok := v.rooms[currentID]
		if !ok || current.Coord == nil {
			continue
		}

		// Process each exit
		for dir, targetID := range current.Exits {
			delta, validDir := DirectionDelta[dir]
			if !validDir {
				v.addWarning(currentID, "Unknown direction '%s' to room %d", dir, targetID)
				continue
			}

			target, ok := v.rooms[targetID]
			if !ok {
				v.addError(currentID, "Exit %s leads to room %d which does not exist", dir, targetID)
				continue
			}

			expectedCoord := current.Coord.Add(delta)

			if target.Coord == nil {
				// First time visiting this room - assign coordinates
				target.Coord = &expectedCoord
				if !visited[targetID] {
					visited[targetID] = true
					queue = append(queue, targetID)
				}
			}
			// If already has coordinates, we'll check consistency in checkDirectionDeltas
		}
	}

	// Report rooms that couldn't be reached
	for roomID, room := range v.rooms {
		if room.Coord == nil {
			v.addWarning(roomID, "Room '%s' is not reachable from origin (room %d)", room.Title, v.originID)
		}
	}
}

// checkCoordinateCollisions checks for multiple rooms at the same coordinate
func (v *Validator) checkCoordinateCollisions() {
	coordToRooms := make(map[string][]int)

	for roomID, room := range v.rooms {
		if room.Coord == nil {
			continue
		}
		key := room.Coord.String()
		coordToRooms[key] = append(coordToRooms[key], roomID)
	}

	for coord, roomIDs := range coordToRooms {
		if len(roomIDs) > 1 {
			sort.Ints(roomIDs)
			v.addError(roomIDs[0], "Coordinate collision at %s: rooms %v occupy the same position", coord, roomIDs)
		}
	}
}

// checkBidirectionalConsistency checks that all exits have proper inverse exits
func (v *Validator) checkBidirectionalConsistency() {
	for roomID, room := range v.rooms {
		for dir, targetID := range room.Exits {
			inverse, hasInverse := InverseDirection[dir]
			if !hasInverse {
				continue // Unknown direction, already warned
			}

			target, ok := v.rooms[targetID]
			if !ok {
				continue // Missing room, already reported
			}

			// Check if target has inverse exit back to this room
			if returnID, hasReturn := target.Exits[inverse]; hasReturn {
				if returnID != roomID {
					v.addError(roomID, "Exit %s to room %d: target's %s exit leads to room %d, not back to %d",
						dir, targetID, inverse, returnID, roomID)
				}
			} else {
				// Check if target has ANY exit back to this room
				hasAnyReturn := false
				for _, retID := range target.Exits {
					if retID == roomID {
						hasAnyReturn = true
						break
					}
				}
				if hasAnyReturn {
					v.addWarning(roomID, "Exit %s to room %d: target room has return path but not via %s (may be intentional)",
						dir, targetID, inverse)
				} else {
					v.addError(roomID, "Exit %s to room %d: target room has no exit back to %d (missing inverse %s)",
						dir, targetID, roomID, inverse)
				}
			}
		}
	}
}

// checkDirectionDeltas verifies that exits lead to rooms at expected coordinates
func (v *Validator) checkDirectionDeltas() {
	for roomID, room := range v.rooms {
		if room.Coord == nil {
			continue
		}

		for dir, targetID := range room.Exits {
			delta, validDir := DirectionDelta[dir]
			if !validDir {
				continue // Already warned
			}

			target, ok := v.rooms[targetID]
			if !ok || target.Coord == nil {
				continue // Already reported
			}

			expectedCoord := room.Coord.Add(delta)
			if !expectedCoord.Equals(*target.Coord) {
				// Calculate actual delta
				actualDeltaX := target.Coord.X - room.Coord.X
				actualDeltaY := target.Coord.Y - room.Coord.Y
				actualDeltaZ := target.Coord.Z - room.Coord.Z

				// Check if it's a multi-square jump
				distance := abs(actualDeltaX) + abs(actualDeltaY) + abs(actualDeltaZ)
				if distance > 2 {
					v.addWarning(roomID, "Exit %s to room %d spans %d squares (actual delta: %d,%d,%d) - may need intermediate rooms",
						dir, targetID, distance, actualDeltaX, actualDeltaY, actualDeltaZ)
				} else {
					v.addError(roomID, "Exit %s to room %d: expected target at %s but found at %s",
						dir, targetID, expectedCoord.String(), target.Coord.String())
				}
			}
		}
	}
}

// ValidateProposal validates a zone proposal against existing rooms
func (v *Validator) ValidateProposal(proposal *Proposal) *ValidationResult {
	propErrors := []ValidationError{}
	propWarnings := []ValidationError{}

	// Convert proposal rooms to RoomInfo
	proposedRooms := make(map[int]*RoomInfo)
	for _, pr := range proposal.Rooms {
		proposedRooms[pr.ID] = &RoomInfo{
			RoomID: pr.ID,
			Zone:   proposal.Zone,
			Coord:  &Coord{X: pr.Coord[0], Y: pr.Coord[1], Z: pr.Coord[2]},
			Exits:  pr.Exits,
		}
	}

	// Check for ID collisions with existing rooms
	for propID := range proposedRooms {
		if _, exists := v.rooms[propID]; exists {
			propErrors = append(propErrors, ValidationError{
				Severity: "ERROR",
				RoomID:   propID,
				Message:  fmt.Sprintf("Room ID %d already exists in the world", propID),
			})
		}
	}

	// Check for coordinate collisions with existing rooms
	for propID, propRoom := range proposedRooms {
		for existingID, existingRoom := range v.rooms {
			if existingRoom.Coord != nil && propRoom.Coord.Equals(*existingRoom.Coord) {
				propErrors = append(propErrors, ValidationError{
					Severity: "ERROR",
					RoomID:   propID,
					Message:  fmt.Sprintf("Would collide with existing room %d at %s", existingID, propRoom.Coord.String()),
				})
			}
		}
	}

	// Check for coordinate collisions within proposal
	for id1, room1 := range proposedRooms {
		for id2, room2 := range proposedRooms {
			if id1 >= id2 {
				continue
			}
			if room1.Coord.Equals(*room2.Coord) {
				propErrors = append(propErrors, ValidationError{
					Severity: "ERROR",
					RoomID:   id1,
					Message:  fmt.Sprintf("Would collide with proposed room %d at %s", id2, room1.Coord.String()),
				})
			}
		}
	}

	// Validate entry point connection
	if proposal.EntryPoint.ConnectsTo > 0 {
		existingRoom, exists := v.rooms[proposal.EntryPoint.ConnectsTo]
		if !exists {
			propErrors = append(propErrors, ValidationError{
				Severity: "ERROR",
				RoomID:   0,
				Message:  fmt.Sprintf("Entry point references non-existent room %d", proposal.EntryPoint.ConnectsTo),
			})
		} else if existingRoom.Coord != nil {
			// Find the first proposed room that connects to the entry point
			for propID, propRoom := range proposedRooms {
				for dir, targetID := range propRoom.Exits {
					if targetID == proposal.EntryPoint.ConnectsTo {
						// Check if the direction delta is correct
						delta, ok := DirectionDelta[dir]
						if ok {
							expectedCoord := propRoom.Coord.Add(delta)
							if !expectedCoord.Equals(*existingRoom.Coord) {
								propErrors = append(propErrors, ValidationError{
									Severity: "ERROR",
									RoomID:   propID,
									Message:  fmt.Sprintf("Exit %s to existing room %d: expected room at %s but existing room is at %s",
										dir, targetID, expectedCoord.String(), existingRoom.Coord.String()),
								})
							}
						}
					}
				}
			}
		}
	}

	// Validate internal consistency of proposal
	for propID, propRoom := range proposedRooms {
		for dir, targetID := range propRoom.Exits {
			delta, validDir := DirectionDelta[dir]
			if !validDir {
				propWarnings = append(propWarnings, ValidationError{
					Severity: "WARN",
					RoomID:   propID,
					Message:  fmt.Sprintf("Unknown direction '%s'", dir),
				})
				continue
			}

			// Check if target is in proposal or existing world
			var targetRoom *RoomInfo
			if tr, ok := proposedRooms[targetID]; ok {
				targetRoom = tr
			} else if tr, ok := v.rooms[targetID]; ok {
				targetRoom = tr
			} else {
				propErrors = append(propErrors, ValidationError{
					Severity: "ERROR",
					RoomID:   propID,
					Message:  fmt.Sprintf("Exit %s leads to room %d which doesn't exist", dir, targetID),
				})
				continue
			}

			if targetRoom.Coord != nil {
				expectedCoord := propRoom.Coord.Add(delta)
				if !expectedCoord.Equals(*targetRoom.Coord) {
					propErrors = append(propErrors, ValidationError{
						Severity: "ERROR",
						RoomID:   propID,
						Message:  fmt.Sprintf("Exit %s to room %d: expected target at %s but found at %s",
							dir, targetID, expectedCoord.String(), targetRoom.Coord.String()),
					})
				}
			}
		}
	}

	return &ValidationResult{
		Errors:   propErrors,
		Warnings: propWarnings,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
