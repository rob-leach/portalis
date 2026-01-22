// Command mapvalidator validates room coordinates and exits
// according to the coordinate system rules defined in docs/COORDINATE_SYSTEM.md
//
// Usage:
//
//	# Validate existing world
//	go run ./cmd/mapvalidator
//
//	# Validate with a proposal
//	go run ./cmd/mapvalidator -proposal path/to/proposal.yaml
//
//	# Specify custom rooms directory
//	go run ./cmd/mapvalidator -rooms ./_datafiles/world/default/rooms
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/GoMudEngine/GoMud/internal/mapvalidator"
)

func main() {
	// Parse command line flags
	roomsDir := flag.String("rooms", "_datafiles/world/default/rooms", "Path to rooms directory")
	proposalFile := flag.String("proposal", "", "Path to zone proposal YAML file (optional)")
	originRoom := flag.Int("origin", 100, "Room ID to use as origin (default: 100 = Town Square)")
	flag.Parse()

	fmt.Println("=== Map Validation Report ===")
	fmt.Println()

	// Parse all room files
	rooms, err := mapvalidator.ParseRoomsFromDirectory(*roomsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing rooms: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d rooms from %s\n\n", len(rooms), *roomsDir)

	// Create validator
	validator := mapvalidator.NewValidator(rooms)
	validator.SetOrigin(*originRoom)

	// Validate existing world
	fmt.Println("EXISTING WORLD:")
	result := validator.Validate()
	printResults(result)

	// If proposal file specified, validate it too
	if *proposalFile != "" {
		fmt.Println()
		fmt.Println("PROPOSAL:")

		proposal, err := mapvalidator.ParseProposalFile(*proposalFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Failed to parse proposal: %v\n", err)
		} else {
			fmt.Printf("Zone: %s (%d rooms)\n", proposal.Zone, len(proposal.Rooms))
			propResult := validator.ValidateProposal(proposal)
			printResults(propResult)
		}
	}

	// Print summary
	fmt.Println()
	totalErrors := len(result.Errors)
	totalWarnings := len(result.Warnings)
	if *proposalFile != "" {
		// Note: proposal results would need to be tracked separately
		// This is simplified for now
	}

	fmt.Printf("Summary: %d errors, %d warnings\n", totalErrors, totalWarnings)

	if totalErrors > 0 {
		os.Exit(1)
	}
}

func printResults(result *mapvalidator.ValidationResult) {
	if len(result.Errors) == 0 && len(result.Warnings) == 0 {
		fmt.Println("[OK] All validation checks passed")
		return
	}

	// Sort by room ID for consistent output
	sortedErrors := make([]mapvalidator.ValidationError, len(result.Errors))
	copy(sortedErrors, result.Errors)
	sort.Slice(sortedErrors, func(i, j int) bool {
		return sortedErrors[i].RoomID < sortedErrors[j].RoomID
	})

	sortedWarnings := make([]mapvalidator.ValidationError, len(result.Warnings))
	copy(sortedWarnings, result.Warnings)
	sort.Slice(sortedWarnings, func(i, j int) bool {
		return sortedWarnings[i].RoomID < sortedWarnings[j].RoomID
	})

	for _, err := range sortedErrors {
		fmt.Printf("[ERROR] Room %d: %s\n", err.RoomID, err.Message)
	}

	for _, warn := range sortedWarnings {
		fmt.Printf("[WARN]  Room %d: %s\n", warn.RoomID, warn.Message)
	}
}
