package mapvalidator

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// RoomYAML represents the structure of a room YAML file
type RoomYAML struct {
	RoomID int                    `yaml:"roomid"`
	Zone   string                 `yaml:"zone"`
	Title  string                 `yaml:"title"`
	Exits  map[string]ExitYAML    `yaml:"exits"`
}

// ExitYAML represents an exit in the YAML file
type ExitYAML struct {
	RoomID int    `yaml:"roomid"`
	Zone   string `yaml:"zone,omitempty"`
	Secret bool   `yaml:"secret,omitempty"`
}

// ParseRoomsFromDirectory walks the rooms directory and parses all YAML files
func ParseRoomsFromDirectory(roomsDir string) (map[int]*RoomInfo, error) {
	rooms := make(map[int]*RoomInfo)

	err := filepath.Walk(roomsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-YAML files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".yaml") {
			return nil
		}

		// Skip the ROOMS.md file if present
		if strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		// Parse the room ID from filename
		baseName := strings.TrimSuffix(info.Name(), ".yaml")
		roomID, err := strconv.Atoi(baseName)
		if err != nil {
			// Not a room file (could be a config file)
			return nil
		}

		room, err := parseRoomFile(path, roomID)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}

		rooms[roomID] = room
		return nil
	})

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

// parseRoomFile parses a single room YAML file
func parseRoomFile(path string, expectedID int) (*RoomInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var roomYAML RoomYAML
	if err := yaml.Unmarshal(data, &roomYAML); err != nil {
		return nil, err
	}

	// Verify the room ID matches the filename
	if roomYAML.RoomID != expectedID {
		return nil, fmt.Errorf("room ID mismatch: file says %d, expected %d", roomYAML.RoomID, expectedID)
	}

	// Convert exits to simple direction -> room ID map
	exits := make(map[string]int)
	for dir, exit := range roomYAML.Exits {
		exits[strings.ToLower(dir)] = exit.RoomID
	}

	return &RoomInfo{
		RoomID: roomYAML.RoomID,
		Zone:   roomYAML.Zone,
		Title:  roomYAML.Title,
		Exits:  exits,
	}, nil
}

// ParseProposalFile parses a zone proposal YAML file
func ParseProposalFile(path string) (*Proposal, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var proposal Proposal
	if err := yaml.Unmarshal(data, &proposal); err != nil {
		return nil, err
	}

	return &proposal, nil
}
