// Package mapvalidator provides tools for validating room coordinates and exits
// according to the coordinate system rules defined in docs/COORDINATE_SYSTEM.md
package mapvalidator

import "fmt"

// Coord represents a 3D position in the world
type Coord struct {
	X int
	Y int
	Z int
}

// RoomInfo holds parsed room data relevant for validation
type RoomInfo struct {
	RoomID int
	Zone   string
	Title  string
	Exits  map[string]int // direction -> target room ID
	Coord  *Coord         // nil if not yet computed
}

// ExitInfo holds information about an exit for error reporting
type ExitInfo struct {
	FromRoom  int
	ToRoom    int
	Direction string
}

// ValidationError represents a single validation error
type ValidationError struct {
	Severity string // "ERROR" or "WARN"
	RoomID   int
	Message  string
}

// ValidationResult holds all validation results
type ValidationResult struct {
	Errors   []ValidationError
	Warnings []ValidationError
}

// ProposalRoom represents a room in a zone proposal
type ProposalRoom struct {
	ID    int            `yaml:"id"`
	Coord [3]int         `yaml:"coord"`
	Exits map[string]int `yaml:"exits"`
}

// Proposal represents a zone proposal file
type Proposal struct {
	Zone       string `yaml:"zone"`
	EntryPoint struct {
		ConnectsTo int    `yaml:"connects_to"`
		Direction  string `yaml:"direction"`
	} `yaml:"entry_point"`
	Rooms []ProposalRoom `yaml:"rooms"`
}

// DirectionDelta maps direction names to coordinate deltas
var DirectionDelta = map[string]Coord{
	"north":     {X: 0, Y: 1, Z: 0},
	"south":     {X: 0, Y: -1, Z: 0},
	"east":      {X: 1, Y: 0, Z: 0},
	"west":      {X: -1, Y: 0, Z: 0},
	"up":        {X: 0, Y: 0, Z: 1},
	"down":      {X: 0, Y: 0, Z: -1},
	"northeast": {X: 1, Y: 1, Z: 0},
	"northwest": {X: -1, Y: 1, Z: 0},
	"southeast": {X: 1, Y: -1, Z: 0},
	"southwest": {X: -1, Y: -1, Z: 0},
}

// InverseDirection maps each direction to its opposite
var InverseDirection = map[string]string{
	"north":     "south",
	"south":     "north",
	"east":      "west",
	"west":      "east",
	"up":        "down",
	"down":      "up",
	"northeast": "southwest",
	"northwest": "southeast",
	"southeast": "northwest",
	"southwest": "northeast",
}

// Add returns a new Coord that is the sum of c and other
func (c Coord) Add(other Coord) Coord {
	return Coord{
		X: c.X + other.X,
		Y: c.Y + other.Y,
		Z: c.Z + other.Z,
	}
}

// Equals returns true if c and other have the same coordinates
func (c Coord) Equals(other Coord) bool {
	return c.X == other.X && c.Y == other.Y && c.Z == other.Z
}

// String returns a formatted coordinate string
func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d,%d)", c.X, c.Y, c.Z)
}
