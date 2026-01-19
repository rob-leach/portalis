# Portalis Macro Zone Map

*Galstaff's Strategic Campaign Atlas*

## Current Zone Layout

```
                                    TUTORIAL (0-1)
                                         |
                                    FROSTFANG (1-5)
                                    [Town Square]
                                   /      |      \
                                  /       |       \
                         West Gate    East Gate    Various
                            |            |          Exits
                            |            |
                    CRYSTAL CAVES    Old Road (278)
                       (3-8)             |
                                    BLADEWORKS FOUNDRY (6-15)  <-- PROBLEM!
                                    [Connected at 278]
                                         |
                                    Hall of Endless Blades (3020)
                                         |
                                    GREAT CROSSROADS (10-20)
                                   /      |      \
                                  /       |       \
                         [WEST]      [SOUTH]     [EAST]
                       Clockwork    Sunken      Verdant
                        Citadel    Kingdom     Reaches
                       (Future)   (Future)    (Future)

        WHISPERING WASTES (5-25) -----> Crystal Caves (alternate entry via 168)
               |
        CATACOMBS (8-13) <-- Accessed from Frostfang Slums

        FROSTFANG SLUMS (5-10) <-- South from Frostfang
```

## Zone Connection Details

| Zone | Entry From | Entry Room | Level Range |
|------|------------|------------|-------------|
| Tutorial | Character Creation | 0 | 0-1 |
| Frostfang | Tutorial | 1 | 1-5 |
| Frostfang Slums | Frostfang | Various | 5-10 |
| Catacombs | Frostfang Slums | 32 | 8-13 |
| Whispering Wastes | Frostfang | Various | 5-25 |
| Crystal Caves | Frostfang (35) OR Whispering Wastes (168) | 2001 | 3-8 |
| Bladeworks Foundry | Frostfang (278) | 3001 | 6-15 |
| Great Crossroads | Bladeworks Foundry (3020) | 5001 | 10-20 |

---

## BLADEWORKS RELOCATION RECOMMENDATION

### Current Problem
Bladeworks Foundry is connected to Room 278 (The Old Road), which is INSIDE Frostfang
zone and only 2 rooms away from the East Gate. This creates several issues:

1. **Level Disparity**: Players can stumble into a level 6-15 zone while still level 1-2
2. **Geographic Confusion**: A massive industrial foundry shouldn't be visible from town
3. **Progression Break**: Players can bypass mid-game content by going straight to Bladeworks

### Recommendation: Move Bladeworks Entry to Great Crossroads

**Proposed New Connection:**
- Remove south exit from Room 278 (The Old Road)
- Add WEST exit from Great Crossroads (5006) leading to a new Bladeworks approach zone
- Create a 3-room approach path from Crossroads to Bladeworks entrance

**Benefits:**
1. Players must progress through ~20 rooms (Foundry) to reach Crossroads
2. Bladeworks becomes appropriately distant from town
3. Creates logical "industrial district" far from civilization
4. Forces players to earn access to mid-game content

**Migration Steps (DO NOT IMPLEMENT YET):**
1. Create new `bladeworks_approach/` zone with rooms 5101-5103
2. Connect Great Crossroads (5006) WEST exit to 5101
3. Connect approach zone to current Bladeworks entrance (3001)
4. Remove/modify Room 278's south exit
5. Update Room 278 description to remove foundry references
6. Update Room 3001 description to reference new approach direction

**Alternative: Keep Current Connection BUT Add Level Warning**
If moving Bladeworks is too disruptive, add a guard or sign at Room 278 warning
players about the danger level ahead. This is less ideal but preserves existing
player familiarity.

---

## SKILL TRAINER PLACEMENT RECOMMENDATIONS

### Current Trainer Locations

| Skill | Current Location | Level | Notes |
|-------|------------------|-------|-------|
| map | Frostwarden Rangers (74) | 1-5 | Good - accessible early |
| search | Frostwarden Rangers (74) | 1-5 | Good - accessible early |
| track | Frostwarden Rangers (74) | 1-5 | Good - accessible early |
| cast | Magic Academy (879) | 1-5 | Needs verification of room location |
| skulduggery | Thieves Den (491) | 1-5 | In Slums - appropriate |
| brawling | Soldiers Training Yard (829) | 1-5 | Needs verification |
| scribe | Dark Acolyte's Chamber (160) | 1-4 | In Catacombs |
| portal | Obelisk (871) | 1-4 | In Whispering Wastes |
| dual-wield | Dueling Gallery (3012) | 1-4 | In Bladeworks - too far! |
| peep | Seer's Alcove (2015) | 1-4 | In Crystal Caves - good |

### Missing Trainers (URGENT)

| Skill | Recommended Zone | Recommended Level Range | Notes |
|-------|------------------|------------------------|-------|
| **inspect** | Crystal Caves | 1-4 | Near Crystalseer (peep trainer) - synergy |
| enchant | New "Enchanter's Tower" zone | 1-4 | Off Whispering Wastes or Crossroads |
| protection | New "Temple of Light" zone | 1-4 | Off Frostfang or Crossroads |
| tame | New "Beastmaster's Lodge" zone | 1-4 | Off Whispering Wastes |
| trading | New "Merchant's Guild" zone | 1-4 | In Frostfang or Slums |

### INSPECT Trainer - Detailed Recommendation

**Best Location: Crystal Caves, Room 2015 (Seer's Alcove)**

The Crystalseer NPC already teaches PEEP. Adding INSPECT creates a natural
"appraisal hub" where players learn to examine items and see through illusions.

**Implementation:**
1. Add `inspect` to the skilltraining block in room 2015.yaml
2. Update Crystalseer mob dialogue to mention both skills
3. Create lore connection between crystal-gazing and item identification

**Alternative Location: New room in Crystal Caves (2019)**
If we want separate NPCs, create a "Crystal Appraiser" room adjacent to
the Seer's Alcove with a specialized inspect trainer.

---

## FUTURE ZONE PLANNING

### Great Crossroads Connections

**EAST - The Verdant Reaches (Planned)**
- Theme: Overgrown wilderness, nature reclaiming ruins
- Level Range: 15-25
- Skill Trainer Potential: tame (Beastmaster's Grove)
- Boss: Ancient Treant or Feral Alpha

**SOUTH - The Sunken Kingdom (Planned)**
- Theme: Flooded ruins, aquatic horrors
- Level Range: 20-30
- Skill Trainer Potential: enchant (Drowned Artificer's Workshop)
- Boss: Leviathan or Merfolk King

**WEST - The Clockwork Citadel (Planned)**
- Theme: Advanced construct civilization (friendly or hostile)
- Level Range: 25-35
- Skill Trainer Potential: trading (Automaton Merchants)
- Boss: Grand Construct or Clockwork Dragon
- Note: Would create thematic link back to Bladeworks Foundry lore

### Progression Path

```
Level 1-5:   Frostfang, Tutorial
Level 3-8:   Crystal Caves
Level 5-10:  Frostfang Slums
Level 6-15:  Bladeworks Foundry
Level 8-13:  Catacombs
Level 5-25:  Whispering Wastes (scales)
Level 10-20: Great Crossroads
Level 15-25: Verdant Reaches (future)
Level 20-30: Sunken Kingdom (future)
Level 25-35: Clockwork Citadel (future)
```

---

## ASCII ZONE MAP

```
                              N
                              |
                    +---------+---------+
                    |     TUTORIAL      |
                    +---------+---------+
                              |
         +--------------------+--------------------+
         |                    |                    |
    [SLUMS]             [FROSTFANG]          [W. WASTES]
     (5-10)               (1-5)                (5-25)
         |                 /|\                    |
    [CATACOMBS]           / | \              [CRYSTAL]
     (8-13)              /  |  \              (3-8)
                        /   |   \
               [DUELING]    |    [OTHER EXITS]
                 HALL       |
                            |
                    [BLADEWORKS]
                      (6-15)
                            |
                            |
                    [GREAT CROSSROADS]
                        (10-20)
                       /   |   \
                      /    |    \
              [WEST]   [SOUTH]   [EAST]
            Clockwork  Sunken    Verdant
            (25-35)   (20-30)   (15-25)
             future    future    future
```

---

*"The campaign map is drawn. The roads are charted. Now we build the dungeons
that lie beyond each horizon!"*

*-- Galstaff, Sorcerer of Light*
