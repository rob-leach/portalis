//go:build e2e
// +build e2e

package main

import (
	"testing"

	"github.com/GoMudEngine/GoMud/internal/buffs"
	"github.com/GoMudEngine/GoMud/internal/colorpatterns"
	"github.com/GoMudEngine/GoMud/internal/configs"
	"github.com/GoMudEngine/GoMud/internal/items"
	"github.com/GoMudEngine/GoMud/internal/keywords"
	"github.com/GoMudEngine/GoMud/internal/mobs"
	"github.com/GoMudEngine/GoMud/internal/mudlog"
	"github.com/GoMudEngine/GoMud/internal/mutators"
	"github.com/GoMudEngine/GoMud/internal/pets"
	"github.com/GoMudEngine/GoMud/internal/plugins"
	"github.com/GoMudEngine/GoMud/internal/quests"
	"github.com/GoMudEngine/GoMud/internal/races"
	"github.com/GoMudEngine/GoMud/internal/rooms"
	"github.com/GoMudEngine/GoMud/internal/spells"
	"github.com/GoMudEngine/GoMud/internal/templates"
	"github.com/GoMudEngine/GoMud/internal/util"
)

// TestWorldDataIntegrity loads all world data files and validates:
// - No duplicate IDs (will panic if duplicates found)
// - All files parse correctly
// - World file structure is valid
//
// This test does NOT start the server or bind to any ports.
func TestWorldDataIntegrity(t *testing.T) {
	// Initialize logger (writes to stderr, no file)
	mudlog.SetupLogger(nil, "LOW", "", false)

	// Load configuration
	configs.ReloadConfig()
	c := configs.GetConfig()

	// Validate world file structure
	t.Run("ValidateWorldFiles", func(t *testing.T) {
		if err := util.ValidateWorldFiles(`_datafiles/world/default`, c.FilePaths.DataFiles.String()); err != nil {
			t.Fatalf("World file validation failed: %v", err)
		}
	})

	// Load all data files in order (same order as main.go)
	// Any duplicate IDs will cause a panic, which Go test will catch

	t.Run("LoadBiomes", func(t *testing.T) {
		rooms.LoadBiomeDataFiles()
	})

	t.Run("LoadSpells", func(t *testing.T) {
		spells.LoadSpellFiles()
	})

	t.Run("LoadRooms", func(t *testing.T) {
		rooms.LoadDataFiles()
	})

	t.Run("LoadBuffs", func(t *testing.T) {
		buffs.LoadDataFiles()
	})

	t.Run("LoadItems", func(t *testing.T) {
		items.LoadDataFiles()
	})

	t.Run("LoadRaces", func(t *testing.T) {
		races.LoadDataFiles()
	})

	t.Run("LoadMobs", func(t *testing.T) {
		mobs.LoadDataFiles()
	})

	t.Run("LoadPets", func(t *testing.T) {
		pets.LoadDataFiles()
	})

	t.Run("LoadQuests", func(t *testing.T) {
		quests.LoadDataFiles()
	})

	t.Run("LoadTemplates", func(t *testing.T) {
		templates.LoadAliases(plugins.GetPluginRegistry())
	})

	t.Run("LoadKeywords", func(t *testing.T) {
		keywords.LoadAliases(plugins.GetPluginRegistry())
	})

	t.Run("LoadMutators", func(t *testing.T) {
		mutators.LoadDataFiles()
	})

	t.Run("LoadColorPatterns", func(t *testing.T) {
		colorpatterns.LoadColorPatterns()
	})
}
