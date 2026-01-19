# Squirrel Tree Mini-Zone: Design Complete

## Executive Summary

**The Squirrel Tree** - A comprehensive design document for a hidden low-level (1-2) mini-zone featuring vertical tree exploration, aggressive squirrel swarms, resource scarcity mechanics, and an optional Squirrel King boss encounter.

**Status:** RESEARCH/DESIGN COMPLETE - Ready for implementation

**Location:** `/Users/i2pi/h/fun/bday2026/portalis/SQUIRREL_TREE_DESIGN.md` (602 lines, 23KB)

---

## What Was Designed

### Zone Structure
- **7 rooms** in a vertical tree configuration (Rooms 4001-4007)
- **Connected to:** Room 1004 (Snarky Squirrel Commons) via secret ceiling trap door
- **Progression:** Entry → Ascent → Central Hub → Branch Exploration → Optional Boss

### Room Layout
```
[4007] Canopy Pinnacle (BOSS) - High-risk arena with Squirrel King
  ↓
[4005]---[4006] Branch Left/Right - Acorn caches, medium mob density
  ↖   ↗
[4004] Mid-Trunk Hub - Central crossroads, secret passage to boss
  ↓
[4003] Tree Ascent - First acorn discovery, increasing mob density
  ↓
[4002] Trunk Entry - First squirrel encounters
  ↓
[4001] Tree Entry - Safe transition chamber (no mobs)
  ↓
[1004] Secret Trap Door (Snarky Squirrel Commons)
```

### Enemies

#### Mob 82: Angry Squirrel (Common)
- **Level:** 2
- **Quantity:** 2-3 per room (4002-4003, 4005-4006), 3-4 in 4007
- **Mechanic:** Uses `callforhelp 2:angry squirrel` to summon adjacent squirrels
- **Combat Tactic:** Swarm mechanic - teaches players aggro management
- **Flavor:** Territorial, squeaky, cooperative fighters
- **Idle Text:** Chittering, scratching, acorn disputes

#### Mob 83: Squirrel King (Boss - Optional)
- **Level:** 4
- **Spawn Rate:** 30% (only in Room 4007)
- **Appearance:** Scarred, one-eyed veteran with acorn crown
- **Mechanic:** Uses `callforhelp 4:angry squirrel` (summons 4 helpers!)
- **Loot:** Crown of Acorns (trophy item)
- **Optional:** Players can complete zone without encountering boss

### Items

#### Item 101: ACORN (Consumable Healing)
- **Healing:** 2 HP per use
- **Placement:** 9-14 total per zone cycle
  - Room 4003: 1x (50% spawn)
  - Room 4005: 2-3x (70% spawn each)
  - Room 4006: 2-3x (70% spawn each)
  - Room 4007: 3-4x (100% guaranteed)
- **Purpose:** Resource scarcity mechanic - players gather acorns to manage health

#### Item 100: Crown of Acorns (Trophy)
- **Drop:** 100% from Squirrel King (when spawned)
- **Benefit:** Flavor/bragging rights only
- **Purpose:** Boss trophy, player achievement marker

### Gameplay Mechanics

**Hidden Discovery:**
- Trap door in Room 1004 requires SEARCH command
- DC 12 difficulty (accessible to new players)
- Flavor text hints in idle messages

**Vertical Exploration:**
- Unique tree-climbing experience vs horizontal dungeons
- 3 levels up (trunk), then branching outward
- Rewards backtracking for acorn gathering

**Swarm Combat:**
- Individual mobs use callforhelp during combat
- 1 squirrel → 4+ quickly if player pulls aggressively
- Teaches tactical aggro management early

**Resource Loop:**
- Damage taken → limited acorns → recovery decisions
- "Do I have enough acorns to continue deeper?"
- Encourages exploration and planning

**Difficulty Scaling:**
- Solo L1-2: Challenging but doable with strategy
- Groups of 2-3: Comfortable with encounters
- Careless pulls: Deadly even for veterans

---

## Design Justifications

### Why This Zone Works

1. **Thematic Integration** - Perfectly extends Room 1004 (Snarky Squirrel Commons) lore
2. **Educational** - Teaches swarm mechanics early in progression
3. **Hidden Reward** - Trap door discovery encourages exploration
4. **Tactical Depth** - Callforhelp mechanic creates meaningful combat decisions
5. **Replayability** - Acorn respawns + 30% boss spawn = varied experiences
6. **Optional Content** - Boss doesn't block progression, rewards bravery
7. **Appropriate Challenge** - Level 1-2 content with engaging mechanics

### Boss Decision: Squirrel King YES

**Why Include?**
- Narrative closure (the squirrels have a KING!)
- Optional encounter (doesn't block zone)
- 30% spawn rate (makes each visit special)
- Trophy drop (Crown of Acorns for bragging rights)
- Perfect for Portalis family campaign (surprising, memorable moments)

**Alternative:** Could skip boss, just high-density chamber. But King adds personality.

---

## Technical Specifications

### IDs to Reserve
| Entity Type | ID Range | Notes |
|-------------|----------|-------|
| Rooms | 4001-4007 | 7 rooms in squirrel_tree zone |
| Mobs | 82-83 | Angry Squirrel, Squirrel King |
| Items | 100-101 | Crown of Acorns, ACORN |

### Files to Create
```
_datafiles/world/default/rooms/squirrel_tree/
  ├── zone-config.yaml
  ├── 4001.yaml through 4007.yaml

_datafiles/world/default/mobs/squirrel_tree/
  ├── 82-angry_squirrel.yaml
  └── 83-squirrel_king.yaml

_datafiles/world/default/items/
  ├── other-0/100-crown_of_acorns.yaml
  └── consumables-30000/101-acorn.yaml
```

### Files to Modify
```
_datafiles/world/default/rooms/frostfang/1004.yaml
  - Add ceiling exit to Room 4001
  - Add ceiling-related nouns
  - Add idle messages about rafters
```

### Scripting Opportunities
- **82-angry_squirrel.js** - Varied idle behaviors, onHurt() aggression boost
- **83-squirrel_king.js** - Boss entry fanfare, onDie() broadcast
- **4004.js** - onCommand_search() hints at secret passage
- **4007.js** - onLoad() conditional boss spawning (30% roll)

---

## Design Document Contents

The comprehensive design document (`SQUIRREL_TREE_DESIGN.md`) includes:

1. **Overview & Philosophy** - Core concept and design goals
2. **Zone Statistics** - Size, level range, difficulty, mechanics
3. **Room Layout & Design** - Detailed specifications for all 7 rooms
4. **Mob Specifications** - Complete YAML structures and behaviors
5. **Item Specifications** - Consumables and trophy drops
6. **Connection Details** - How Room 1004 connects to zone
7. **Gameplay Flow** - Player progression through zone
8. **Combat Mechanics** - Swarm tactics, resource loops, difficulty
9. **Design Rationale** - Why each element works
10. **Technical Implementation** - IDs, files, scripting opportunities
11. **Testing Checklist** - Verification steps for implementation
12. **Future Expansion Ideas** - Potential enhancements
13. **References** - Related zones and lore connections

**Document Stats:**
- 602 lines
- 42 major sections (## headers)
- 23KB total size
- Includes YAML templates, room diagrams, testing checklists

---

## Key Design Features

### Hidden Discovery
- Trap door requires SEARCH skill check (DC 12)
- Flavor hints in Room 1004 idle messages
- Rewards player curiosity

### Vertical Progression
- Unique tree-climbing gameplay
- 3 levels UP (trunk), then branches OUT
- Different from typical horizontal dungeon layouts

### Swarm Mechanics
- `callforhelp` system teaches aggro management
- 1 mob becomes 4+ quickly
- Prepared players thrive, reckless players die

### Resource Scarcity
- 9-14 ACORNs total (not guaranteed)
- 2 HP healing each (sustenance, not full recovery)
- Forces planning: "Can I afford to continue?"

### Optional Boss
- Squirrel King (30% spawn in Room 4007)
- Doesn't block zone completion
- High-risk/high-reward trophy encounter

### Thematic Consistency
- All text emphasizes squirrel territoriality
- Acorn economy (food, currency, resources)
- King-subject hierarchy

---

## Campaign Integration

### In Portalis Campaign
- **Hidden for new players** - Rewards thorough exploration of Room 1004
- **Teaches mechanics** - Swarm tactics appear in larger zones later
- **Boss introduction** - Players encounter their first real boss encounter (optional)
- **Resource management** - ACORNs teach planning and recovery
- **Lore building** - Squirrel kingdom narrative extends world

### Player Experience
- ~10 minutes of gameplay per visit
- Repeatable for acorn gathering and respawning boss
- Appropriate challenge for Levels 1-2
- Natural breakpoint: can stop in Room 4004 (hub) before continuing

---

## What's Next

### Implementation (Not Done - Design Only)
1. Create zone directory and files
2. Build room YAML files with proper exits
3. Create mob YAML definitions
4. Create item YAML definitions
5. Modify Room 1004 to add entrance
6. Write mob and room scripts
7. Test all mechanics and difficulty

### GitHub Issue
The design is ready to be converted to a GitHub issue to track implementation:
- Title: "Design: Squirrel Tree Mini-Zone (Level 1-2)"
- Label: "design", "mini-zone"
- Content: Summary of this design document

### Estimated Implementation Time
- Room creation: 1-2 hours
- Mob/item creation: 30 minutes
- Script writing: 1-2 hours
- Testing & balance: 1-2 hours
- **Total: 4-6 hours** of implementation work

---

## Deliverable Summary

### What Was Delivered

✓ **Comprehensive Design Document** (SQUIRREL_TREE_DESIGN.md)
- 602 lines of detailed specifications
- Complete room layouts and descriptions
- Full mob and item YAML templates
- Gameplay flow documentation
- Testing checklist
- Technical implementation guide

✓ **Design Validation**
- Analyzed existing zones (Crystal Caves, Bladeworks Foundry, Dueling Hall)
- Verified patterns and mechanics used
- Ensured consistency with GoMUD engine architecture
- Confirmed available ID ranges

✓ **Galstaff Campaign Notes**
- Updated GALSTAFF.md with design research session
- Documented decision-making rationale
- Noted campaign impact
- Prepared for next phase (implementation)

### What's NOT Included (Design Only)
- No zone directory created
- No YAML files written
- No scripts implemented
- No testing performed
- No git commits made

This is strictly RESEARCH/DESIGN phase.

---

## Files Created

1. **SQUIRREL_TREE_DESIGN.md** (23KB, 602 lines)
   - Comprehensive design specification
   - Ready for implementation team

2. **GALSTAFF.md** (updated)
   - Session 5 notes on design research
   - Campaign impact analysis
   - Decision documentation

3. **SQUIRREL_TREE_SUMMARY.md** (this file)
   - Quick reference guide
   - Deliverable overview
   - Implementation roadmap

---

## How to Proceed

### For Implementation
1. Read `SQUIRREL_TREE_DESIGN.md` thoroughly
2. Create GitHub issue from design document
3. Create branch: `feature/squirrel-tree`
4. Follow implementation checklist in design document
5. Test all mechanics and difficulty
6. Create PR with zone implementation

### For Design Review
1. Check Room IDs don't conflict (4001-4007)
2. Verify mob/item IDs are available (82-83, 100-101)
3. Confirm Room 1004 can be modified (safe to add ceiling exit)
4. Review mob difficulty vs Big Rat (Level 5) baseline

---

## References

**Related Zones (Design Patterns Used):**
- Crystal Caves: Exploration + resource gathering
- Bladeworks Foundry: Swarm mechanics, two-tier difficulty
- Dueling Hall: Discovery mechanic, NPC trainer

**Engine Mechanics Leveraged:**
- `callforhelp` - Summon adjacent mobs during combat
- Search skill checks - Hidden entrance discovery
- Mob groups - Swarm coordination
- Consumable items - Healing via food
- Conditional spawning - Optional boss encounter

---

## Conclusion

The Squirrel Tree is a thoughtfully designed mini-zone that:
- Adds depth to Portalis' early game
- Teaches important combat mechanics (swarm tactics, aggro management)
- Rewards exploration (hidden trap door)
- Offers optional challenges (Squirrel King boss)
- Provides memorable moments (acorn gathering, rare boss encounters)

The design is complete, specification is comprehensive, and the zone is ready for implementation whenever development time allows.

**Galstaff's parting words:**

*"The Squirrel Tree stands complete in the realm of design! Now awaits only the master builder to bring it forth into reality. But mark me - when the first player discovers that ceiling trap door and climbs into the heart of the tree, faces down the chittering horde, and perhaps encounters the legendary Squirrel King... THEN shall our work be justified!"*

*Galstaff nods with satisfaction, his acorn-shaped helm catching the light.*

---

**Document Created:** 2026-01-19
**Designer:** Galstaff, Sorcerer of Light
**Campaign:** Portalis (birthday2026)
**Status:** DESIGN RESEARCH COMPLETE
