# GitHub Issue: Design: Squirrel Tree Mini-Zone (Level 5-8)

## Title
Design: Squirrel Tree Mini-Zone (Level 5-8 BIGRAT Difficulty)

## Labels
- design
- mini-zone
- exploration
- feature

---

## Description

The **Squirrel Tree** is a hidden mini-zone (Level 5-8 BIGRAT-tier difficulty) accessible from Room 1004 (Snarky Squirrel Commons) via a secret trap door in the ceiling. This dungeon features a vertical tree structure with branching rooms, aggressive squirrel swarms using swarm mechanics, and an optional Squirrel King boss encounter. Difficulty scales with player progression.

### Design Philosophy
- Hidden discovery mechanic (search skill DC 12 for trap door)
- Vertical gameplay (climbing up the tree rather than traditional horizontal exploration)
- Swarm combat tactics (mobs call for help and gang up on players)
- Resource scarcity (ACORNs for healing, scattered in branch rooms)
- Low-difficulty encounters with overwhelming numbers if careless

---

## Zone Overview

| Attribute | Value |
|-----------|-------|
| **Room Count** | 7 rooms |
| **Level Range** | 5-8 (BIGRAT tier, matches player progression) |
| **Difficulty** | Challenging solo, manageable for groups with caution |
| **Theme** | Natural (tree interior, woodland) |
| **Mechanics** | Vertical exploration + swarm combat + resource gathering |
| **Parent Zone** | Frostfang (Room 1004) |
| **Access** | Secret ceiling trap door (search DC 12) |

---

## Room Layout

```
                    [4007]
                   Canopy
                  Pinnacle
                      |
           [4005] ---- [4006]
         Branch Left  Branch Right
              \        /
              [4004]
            Mid-Trunk
              |
            [4003]
          Tree Ascent
              |
            [4002]
          Trunk Entry
              |
          [4001]
         Tree Entry
              |
         [1004] Ceiling
    Snarky Squirrel Commons
```

### Room Descriptions

| Room | Name | Purpose | Mobs | Items |
|------|------|---------|------|-------|
| 4001 | Tree Entry Hall | Safe transition chamber | None | None |
| 4002 | Trunk Entry | First encounters, flavor | 2-3 | None |
| 4003 | Tree Ascent | Mob density ramps up | 2-3 | 1x ACORN (50%) |
| 4004 | Mid-Trunk Hub | Central crossroads, secret passage | 1-2 | None |
| 4005 | Branch Left | Acorn cache room | 2 | 2-3x ACORN (70% each) |
| 4006 | Branch Right | Acorn cache room | 2 | 2-3x ACORN (70% each) |
| 4007 | Canopy Pinnacle | Boss chamber, high-risk | 3-4 + Boss | 3-4x ACORN (100%) |

---

## Mobs

### Mob 82: Angry Squirrel (Common Enemy, Level 5)

**Description:** A plump squirrel with russet coat, fur standing on end in perpetual agitation. Aggressive and territorial. BIGRAT-tier threat.

**Mechanics:**
- Spawns: 2-3 per room (rooms 4002-4003, 4005-4006), 3-4 in room 4007
- Combat: Uses `callforhelp 2:angry squirrel:shriek` to summon adjacent squirrels
- Swarm Tactic: 1 squirrel → 4+ quickly (deadly if not managed)
- Idle Behavior: Chittering, scratching, defending territory
- HP: Moderate (BIGRAT-tier equivalent)
- Loot: 1 gold (itemdropchance: 3%)

**YAML Template:**
```yaml
mobid: 82
zone: Squirrel Tree
itemdropchance: 3
hostile: false
maxwander: 1
groups:
  - squirrel-tree
  - squirrels
combatcommands:
  - 'callforhelp 2:angry squirrel:lets out a piercing shriek!'
idlecommands:
  - 'emote chittering aggressively'
  - 'emote scratches at the bark'
  - ''
activitylevel: 25
character:
  name: angry squirrel
  level: 5  # BIGRAT-tier difficulty
  raceid: 10
  gold: 1
  alignment: -50
  stats:
    vitality:
      training: -1
  equipment:
    body:
      itemid: 20001
```

### Mob 83: Squirrel King (Boss, Level 7, 30% Spawn Rate)

**Description:** Massive scarred squirrel, one eye missing, wears an acorn crown. Undisputed monarch of the tree. Mid-tier boss challenge.

**Mechanics:**
- Spawn: 30% chance in Room 4007 only
- Combat: Uses `callforhelp 4:angry squirrel` (summons 4 helpers!)
- Boss Drop: Crown of Acorns (100% drop rate when spawned)
- Behavior: Menacing, graceful, intelligent
- HP: Significant (2 levels above regular squirrels, demanding encounter)
- Loot: Crown of Acorns + gold (itemdropchance: 50%)

**YAML Template:**
```yaml
mobid: 83
zone: Squirrel Tree
itemdropchance: 50
hostile: true
maxwander: 0
groups:
  - squirrel-tree
  - squirrel-royalty
combatcommands:
  - 'callforhelp 4:angry squirrel:ROYAL SHRIEK!'
idlecommands:
  - 'emote surveys its domain'
  - ''
activitylevel: 40
character:
  name: squirrel king
  level: 7  # Mid-tier boss challenge (2 levels above regular squirrels)
  raceid: 10
  gold: 5
  alignment: -100
  equipment:
    body:
      itemid: 20001
```

---

## Items

### Item 101: ACORN (Consumable Healing)

**Purpose:** Resource-scarce healing consumable scattered throughout zone

**Specifications:**
- Type: Food/Consumable
- Healing: 2 HP per use
- Uses: 1
- Placement: 9-14 per zone run
  - Room 4003: 1x (50% spawn)
  - Room 4005: 2-3x (70% spawn each)
  - Room 4006: 2-3x (70% spawn each)
  - Room 4007: 3-4x (100% guaranteed)
- Respawn: Zone reset (recommend 30 real minutes)

**YAML Template:**
```yaml
itemid: 101
name: acorn
namesimple: acorn
description: 'A perfectly formed acorn still in its cap. Fresh and nutty. You could eat this.'
type: food
subtype: consumable
uses: 1
buffids:
  - 18  # Healing buff (2 HP)
value: 0
```

### Item 100: Crown of Acorns (Boss Drop Trophy)

**Purpose:** Unique trophy item dropped by Squirrel King boss

**Specifications:**
- Type: Object/Unique
- Drop: 100% from Squirrel King (when spawned)
- Benefit: Flavor only (no mechanical stat bonuses)
- Purpose: Boss trophy, achievement marker, bragging rights
- Rarity: Limited (only from optional boss)

**YAML Template:**
```yaml
itemid: 100
name: crown of acorns
namesimple: crown
description: 'A crude crown woven from acorns and plant fibers. Surprisingly sturdy. A squirrel wore this. Now you do.'
type: object
subtype: unique
value: 10
```

---

## Gameplay Flow

### Discovery Phase
1. Player explores Room 1004 (Snarky Squirrel Commons)
2. Uses SEARCH command on ceiling (DC 12 check)
3. Discovers hidden trap door
4. Climbs up to Room 4001 (safe entry)

### Ascent Phase (Rooms 4002-4003)
- First squirrel encounters (2-3 each)
- Learn swarm mechanics via callforhelp
- Discover first ACORN (resource gathering begins)
- Mobs non-aggressive until provoked

### Branching Phase (Rooms 4005-4006)
- More squirrels (2 each) defending acorn caches
- Maximum acorn gathering opportunity
- Risk: Pulling too many mobs = deadly
- Reward: Sufficient acorns to continue safely

### Boss Phase (Room 4007, Optional)
- High-risk chamber (3-4 squirrels minimum)
- 30% chance Squirrel King appears
- Maximum acorn cache (guaranteed 3-4)
- Highest rewards but deadly difficulty

---

## Combat Mechanics

### Swarm Tactic System
- Individual Angry Squirrels use `callforhelp` during combat
- Summons adjacent squirrels from nearby rooms
- 1 squirrel becomes 4+ quickly if player pulls recklessly
- **Teaching Moment:** Teaches aggro management early in progression

### Resource Loop
1. Player takes damage from squirrel swarms
2. ACORNs heal 2 HP (sustenance-level, not full recovery)
3. Creates tension: "Do I have enough acorns to continue?"
4. Encourages exploration and recovery planning

### Difficulty Scaling
- **Solo L5-6:** Challenging but doable with strategy and acorn management
- **Solo L7+:** Difficult but manageable with caution, swarms still deadly
- **Groups of 2-3 (L5+):** Comfortable with most encounters
- **Careless Aggro:** Deadly even for experienced groups
- **Squirrel King:** Mid-tier boss challenge, requires preparation or group support

---

## Connection to Room 1004

### Room 1004 Modifications

**Add to exits:**
```yaml
ceiling:
  roomid: 4001
  lock:
    lockid: "squirrel-trapdoor"
    difficulty: 12  # Search DC 12
    sequence: "SEARCH"
```

**Add to nouns:**
```yaml
ceiling: "The ceiling shows signs of wear. Is that... a hole? Strange chittering from above."
trapdoor: "A trap door disguised as part of the rafters. Easy to miss if you don't search carefully."
```

**Add to idlemessages:**
```yaml
- "Something skitters in the ceiling above, causing dust to rain down."
- "You hear faint chittering from somewhere in the rafters."
```

---

## Design Rationale

### Why This Works

**Hidden Discovery**
- Trap door requires search skill, rewards exploration
- Natural for "Snarky Squirrel" location thematic

**Vertical Progression**
- Unique gameplay vs horizontal dungeons
- Tree-climbing creates immersion
- Visual progress (going UP)

**Swarm Mechanics**
- Teaches tactical combat without super-tough individual mobs
- Emphasizes "squirrel army" concept authentically
- Players learn aggro management early

**Resource Scarcity**
- 9-14 ACORNs total (limited supply)
- 2 HP healing (sustenance, not full)
- Forces planning and recovery decisions

**Optional Boss**
- Doesn't block zone progression
- Rewards brave exploration
- 30% spawn makes encounters special
- Trophy drop adds achievement feeling

### Boss Decision: YES to Squirrel King

**Why Include:**
- Narrative closure (squirrels have a KING!)
- Optional encounter (players choose this risk)
- 30% spawn rate (makes each visit unpredictable)
- Suitable for family campaign (surprising, memorable moments)
- Trophy drop provides achievement

---

## Technical Specifications

### IDs to Reserve

| Type | IDs | Purpose |
|------|-----|---------|
| Rooms | 4001-4007 | Squirrel Tree zone (7 rooms) |
| Mobs | 82-83 | Angry Squirrel, Squirrel King |
| Items | 100-101 | Crown of Acorns, ACORN |

### Directory Structure
```
_datafiles/world/default/
├── rooms/squirrel_tree/
│   ├── zone-config.yaml
│   ├── 4001.yaml through 4007.yaml
├── mobs/squirrel_tree/
│   ├── 82-angry_squirrel.yaml
│   ├── 83-squirrel_king.yaml
│   └── scripts/
│       ├── 82-angry_squirrel.js
│       ├── 83-squirrel_king.js
│       ├── 4004.js
│       └── 4007.js
└── items/
    ├── other-0/100-crown_of_acorns.yaml
    └── consumables-30000/101-acorn.yaml
```

### Files to Modify
- `_datafiles/world/default/rooms/frostfang/1004.yaml` - Add ceiling exit

---

## Testing Checklist

- [ ] Trap door is discoverable in Room 1004 via search DC 12
- [ ] Trap door connects to Room 4001 properly
- [ ] All 7 rooms created with correct exits
- [ ] Room connections form proper vertical tree structure
- [ ] Squirrel mobs spawn in correct quantities per room
- [ ] Squirrels don't aggro until player initiates combat
- [ ] Squirrels use callforhelp during combat correctly
- [ ] ACORNs spawn in correct rooms with correct percentages
- [ ] ACORNs can be picked up and used
- [ ] ACORNs heal 2 HP when consumed
- [ ] ACORNs respawn on zone reset
- [ ] Squirrel King spawns 30% of the time in Room 4007
- [ ] Squirrel King drops Crown of Acorns
- [ ] Zone difficulty is appropriate for Level 1-2
- [ ] Secret passage to Room 4007 is discoverable (search DC 14)
- [ ] All flavor text displays correctly
- [ ] No typos or formatting errors

---

## Implementation Notes

### Estimated Effort
- Room creation: 1-2 hours
- Mob/item creation: 30 minutes
- Script writing: 1-2 hours
- Testing & balance: 1-2 hours
- **Total: 4-6 hours**

### Priority
- Medium (adds low-level content and exploration element)
- Complements existing Frostfang zone
- Can be deferred until after critical skill trainers are in place

### Dependencies
- Room 1004 must remain unmodified structurally
- IDs 4001-4007, 82-83, 100-101 must be available
- Buff ID 18 (healing effect) must exist or be created

---

## References

**Related Zones:**
- Crystal Caves (exploration + resource gathering pattern)
- Bladeworks Foundry (swarm mechanics, two-tier difficulty)
- Dueling Hall (discovery mechanic)

**Engine Features Used:**
- `callforhelp` command (mob group summons)
- Search skill checks (hidden discoveries)
- Mob groups/factions (squirrel cooperation)
- Consumable items (healing)
- Conditional spawning (optional boss)

**Lore Integration:**
- Extends Snarky Squirrel Commons narrative
- Explains mysterious squirrel references in Room 1004
- Fits naturally into Frostfang low-level progression

---

## Acceptance Criteria

A PR implementing this design will be accepted when:

1. ✓ All 7 rooms created with proper descriptions and exits
2. ✓ Both mob types created with correct stats and behaviors
3. ✓ Both item types created with correct effects
4. ✓ Room 1004 ceiling exit connects properly
5. ✓ All spawning rates match specifications
6. ✓ Callforhelp mechanics function as designed
7. ✓ ACORN healing works correctly
8. ✓ Squirrel King 30% spawn works
9. ✓ All testing checklist items pass
10. ✓ Difficulty is appropriate for level range

---

## Full Design Documentation

Complete specifications available in: `SQUIRREL_TREE_DESIGN.md` (602 lines, 23KB)

This document includes:
- Detailed room descriptions
- Full YAML templates
- Scripting opportunities
- Extended design rationale
- Future expansion ideas

---

**Design Status:** RESEARCH COMPLETE - Ready for Implementation
**Designer:** Galstaff, Sorcerer of Light
**Created:** 2026-01-19
