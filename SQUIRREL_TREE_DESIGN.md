# Design: Squirrel Tree Mini-Zone (Level 1-2)

## Overview

**The Squirrel Tree** is a hidden mini-zone accessible from Room 1004 (Snarky Squirrel Commons) via a secret trap door in the ceiling. This low-level dungeon (comparable to Big Rat encounters) features a vertical tree structure with branching rooms and an aggressive swarm of Angry Squirrels.

**Design Philosophy:**
- Hidden discovery mechanic (search skill DC 12 for trap door)
- Vertical gameplay (climbing up the tree rather than traditional horizontal exploration)
- Swarm combat tactics (mobs call for help and gang up on players)
- Resource scarcity (ACORNs for healing, scattered in branch rooms)
- Low-difficulty encounters with overwhelming numbers if careless

---

## Zone Statistics

| Metric | Value |
|--------|-------|
| **Room Count** | 7 rooms total |
| **Level Range** | 5-8 (BIGRAT difficulty tier, scaling with player progression) |
| **Difficulty** | Dangerous for solo, challenging for groups if swarmed |
| **Theme** | Natural/Organic (tree interior, woodland vibes) |
| **Mechanics** | Vertical progression + swarm combat + resource gathering |
| **Connected Zones** | Frostfang (Room 1004) |

---

## Room Layout & Design

### Room IDs: 4001-4007

The zone uses a vertical tree structure moving UP three levels, then branching outward:

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

### Detailed Room Specifications

#### Room 4001: Tree Entry Hall (Transition Room)
- **Connection:** Secret trap door in ceiling of Room 1004
- **Layout:** Small chamber directly below the trap door
- **Description:** The soft interior of the tree hollow, roots forming natural archways and platforms. A spiral of bark ridges serves as primitive stairs leading upward. Fresh sunlight filters down from the opening above, and the air smells of acorns and tree sap.
- **Special Features:**
  - Trap door entrance (ceiling exit north to Room 1004)
  - Natural climbing surfaces (no climbing checks needed)
  - No mobs spawn here (safe resting point)
  - Exit up leads to 4002

#### Room 4002: Trunk Entry
- **Description:** The interior of the massive tree trunk, with bark walls forming a natural vertical tunnel. Darker than below, but pale mushrooms glow faintly on the walls providing light. Root systems form natural stepping stones spiraling upward. The air grows warmer and stuffier.
- **Mobs Spawned:** 2-3x Angry Squirrel (common)
- **Items:** None
- **Idle Messages:**
  - "A squirrel's chittering echoes from above."
  - "You hear the pitter-patter of tiny claws on wood."
  - "Acorn dust drifts down from higher branches."
- **Exits:** Down: 4001, Up: 4003

#### Room 4003: Tree Ascent
- **Description:** A chamber where the trunk begins to widen, preparing to branch. Knots and hollows in the bark create a labyrinthine wall texture. Several main branches split off in different directions, visible as dark passages. The smell of acorns intensifies here.
- **Mobs Spawned:** 2-3x Angry Squirrel (common)
- **Items:** 1x ACORN (50% spawn chance, heals 2 HP)
- **Idle Messages:**
  - "A quarrelsome chittering match breaks out above."
  - "You see shadows of squirrels darting between crevices."
  - "The creaking of branches echoes throughout the hollow."
- **Exits:** Down: 4002, Up: 4004

#### Room 4004: Mid-Trunk (Central Hub)
- **Description:** A spacious chamber at the tree's center where the trunk has widened considerably. Multiple major branches split off in different directions, creating a natural crossroads. Acorn husks litter the ground, evidence of seasons of squirrel activity. A warm air current moves through, suggesting the branches lead to more open areas.
- **Mobs Spawned:** 1-2x Angry Squirrel (common)
- **Items:** None visible (acorns in branches)
- **Idle Messages:**
  - "Multiple branches intersect here, each filled with activity."
  - "The chittering grows louder from multiple directions."
- **Special Features:** Secret passage to 4007 (search DC 14, high-risk endpoint)
- **Exits:** Down: 4003, West: 4005, East: 4006, Up (secret): 4007

#### Room 4005: Branch Left (Acorn Cache Left)
- **Description:** A relatively narrow branch passage curving gently. The wood here is smoother, polished by centuries of squirrel traffic. A natural widening creates a nesting chamber lined with shredded bark. Dozens of acorn shells are stacked in small piles, evidence of the area's primary resource.
- **Mobs Spawned:** 2x Angry Squirrel (medium priority)
- **Items:** 2-3x ACORN (70% spawn chance each)
- **Idle Messages:**
  - "A squirrel skitters frantically, defending its stash."
  - "The sound of cracking acorns echoes through the chamber."
  - "You see fresh bite marks in a nearly-complete acorn husk."
- **Exits:** East: 4004

#### Room 4006: Branch Right (Acorn Cache Right)
- **Description:** Mirror of Branch Left - another branch passage with a nesting chamber. This one shows signs of more recent activity, with fresher acorn shells and scattered nuts. The bark shows scratches and claw marks from disputes over resources.
- **Mobs Spawned:** 2x Angry Squirrel (medium priority)
- **Items:** 2-3x ACORN (70% spawn chance each)
- **Idle Messages:**
  - "A particularly aggressive squirrel hisses at your approach."
  - "Multiple squirrels are bickering over food nearby."
  - "The acorn cache here is actively being defended."
- **Exits:** West: 4004

#### Room 4007: Canopy Pinnacle (Top of Tree - Boss Room)
- **Description:** The topmost chamber of the tree, where the trunk splits into the finest branches. Sunlight streams directly down through an opening in the canopy high above, creating a dappled light show. This is clearly the command center of the squirrel territory - the most prized acorns are stored here in alcoves carved over years of habitation.
- **Mobs Spawned:**
  - 3-4x Angry Squirrel (aggressive)
  - 1x Squirrel King (optional boss, Level 4, 30% spawn rate)
- **Items:**
  - 3-4x ACORN (premium cache, 100% spawn rate)
  - Boss drop (if Squirrel King spawns): Crown of Acorns (misc item, +1 CHA, flavor text)
- **Idle Messages:**
  - "A large, scarred squirrel surveys its domain from above."
  - "The air crackles with territorial energy."
  - "You hear the distinct chatter of what sounds like... royalty?"
- **Special Mechanic:**
  - Secret passage from 4004 (search DC 14)
  - High-risk area (4+ squirrels vs 1-2 players is deadly)
  - Serves as "true ending" for zone exploration
- **Exits:** Down (secret): 4004

---

## Mobs: Angry Squirrels

### Mob 82: Angry Squirrel (Common Enemy, Level 5)

**YAML Structure:**
```yaml
mobid: 82
zone: Squirrel Tree
itemdropchance: 3
hostile: false  # Non-aggressive but territorial
maxwander: 1
groups:
  - squirrel-tree
  - squirrels
combatcommands:
  - 'callforhelp 2:angry squirrel:lets out a piercing shriek that summons help!'
  - 'emote chatters angrily and lunges with tiny claws'
  - 'emote circles, looking for an opening'
idlecommands:
  - 'emote chittering aggressively'
  - 'emote squeaks indignantly'
  - 'emote scratches at the bark'
  - ''
activitylevel: 25  # Very active
character:
  name: angry squirrel
  description: 'A plump squirrel with a distinctive russet coat, its fur standing on end in a perpetual state of agitation. Its bushy tail twitches erratically, and its beady black eyes gleam with territorial fury. Small tufts of fur stick out at odd angles, suggesting it has been in many scuffles over acorn supremacy. Tiny claws, wickedly sharp, click against the wood as it moves. Despite its adorable rodent features, something in its demeanor suggests this creature is FAR more dangerous than it appears - this is a squirrel that does NOT back down.'
  raceid: 10  # Rat race (rodent classification)
  level: 5  # BIGRAT-tier difficulty (matches player progression)
  alignment: -50  # Chaotic/aggressive
  gold: 1
  stats:
    vitality:
      training: -1
  equipment:
    body:
      itemid: 20001
```

**Behavior:**
- Spawns in groups of 2-3 per room
- Uses `callforhelp 2:angry squirrel:flavor text` command during combat
- Summons adjacent squirrels from neighboring rooms
- Moderate HP (BIGRAT-tier equivalent)
- Fast attack speed (high activity level)
- Flavor idle text emphasizing territorial rage and cooperation

### Mob 83: Squirrel King (Boss - Optional, Level 7)

**Appearance & Behavior:**
- Larger, scarred squirrel with distinctive features
- Level 7 (significantly harder than regular mobs, mid-tier boss challenge)
- Dual-wields tiny daggers (cosmetic flavor)
- Rare spawn (30% chance, only in Room 4007)
- Triggers when any player encounters the pinnacle chamber

**YAML Structure:**
```yaml
mobid: 83
zone: Squirrel Tree
itemdropchance: 50
hostile: true
maxwander: 0  # Never wanders from pinnacle
groups:
  - squirrel-tree
  - squirrel-royalty
combatcommands:
  - 'callforhelp 4:angry squirrel:lets out a ROYAL SHRIEK that echoes through the entire tree!'
  - 'emote swipes with tiny dagger claws'
  - 'emote leaps with surprising agility'
activitylevel: 40
character:
  name: squirrel king
  description: 'The undisputed monarch of this tree, a massive squirrel compared to its brethren. Its russet fur is matted with age and bearing the scars of countless territorial disputes. One eye is missing, replaced by a jagged scar that runs from cheekbone to temple, giving it a permanently menacing expression. A crown of acorns, somehow ingeniously balanced on its head, serves as both regalia and portable food supply. Its movements are deliberate and graceful, those of a creature that has survived through cunning and ferocity. When it moves, smaller squirrels instinctively defer - this is the KING.'
  raceid: 10
  level: 7  # Mid-tier boss challenge (2 levels above regular squirrels)
  alignment: -100  # Chaotic evil
  gold: 5
  equipment:
    body:
      itemid: 20001
```

**Boss Mechanics:**
- First appearance causes dramatic entry (idle message: "A MASSIVE scarred squirrel rises from the acorn throne!")
- Calls entire room to help (up to 4 additional squirrels)
- Death drops: Crown of Acorns (quest item / trophy for bragging rights)
- Optional encounter - players can avoid by not exploring Room 4007
- 30% spawn chance makes encounters with the King feel special
- Solo difficulty: Formidable challenge even for experienced players
- Group difficulty: Dangerous if not careful, but manageable with teamwork

---

## Items

### Item 101: ACORN

**Purpose:** Consumable healing item, scattered throughout zone

**YAML Structure:**
```yaml
itemid: 101
name: acorn
namesimple: acorn
description: 'A perfectly formed acorn still in its cap. It smells fresh and nutty. You could probably eat this... if you were that hungry. Or that brave.'
type: food
subtype: consumable
uses: 1
buffids:
  - 18  # Healing buff (verify: should heal 2 HP)
value: 0  # Found only in zone, not purchased
```

**Placement Strategy:**
- Room 4003: 1x (50% spawn) - First acorn discovery
- Room 4005: 2-3x (70% spawn each) - Cache room
- Room 4006: 2-3x (70% spawn each) - Cache room
- Room 4007: 3-4x (100% spawn - guaranteed) - Premium cache
- **Total potential:** 9-14 ACORNs on a full zone run

**Respawn:** ACORNs respawn on zone reset (recommend 30 real minutes)

**Healing Value:** 2 HP per acorn (represents sustenance-level recovery, not full healing)

### Item 100: Crown of Acorns (Boss Drop)

**Purpose:** Flavor/trophy item from Squirrel King boss encounter

**YAML Structure:**
```yaml
itemid: 100
name: crown of acorns
namesimple: crown
description: 'A crude crown constructed entirely of acorns woven together with plant fibers. It''s surprisingly sturdy and well-balanced. It still carries the faint scent of forest. A squirrel wore this. You''re now wearing acorn royalty regalia. Make of that what you will.'
type: object
subtype: unique
value: 10
equipmentslots: []  # Cosmetic only, not equippable
```

**Mechanics:**
- Dropped by Squirrel King (only if boss spawns, 30%)
- No mechanical stat benefits (pure flavor/trophy)
- Allows players to brag: "I defeated the Squirrel King!"
- Can be stacked in inventory for trophy collection
- Potential future use: NPC quest reward ("Bring me proof you've conquered the Squirrel King!")

---

## Connection to Room 1004

### Room 1004 (Snarky Squirrel Commons) Modifications

**Add to YAML exits section:**
```yaml
exits:
  northwest:
    roomid: 1
  ceiling:  # Secret trap door
    roomid: 4001
    lock:
      lockid: "squirrel-trapdoor"
      difficulty: 12  # Search skill DC 12
      sequence: "SEARCH"  # Flavor: player must search ceiling
```

**Add to nouns section:**
```yaml
nouns:
  (existing nouns...)
  ceiling: "The ceiling shows signs of wear and damage. Is that... a hole? And strange chittering sounds are coming from above."
  trapdoor: "A trap door disguised to look like part of the rafters. If you didn't know to look for it, you'd never notice it."
  hole: "A suspicious opening in the ceiling above the stage. Now that you look closer, you can definitely hear movement up there."
```

**Add to idlemessages:**
```yaml
idlemessages:
  - (existing messages...)
  - "Something skitters in the ceiling above, causing dust to rain down onto the stage."
  - "You hear faint chittering from somewhere in the rafters above."
```

---

## Gameplay Flow

### Player Entry Sequence

1. **Discovery Phase (Room 1004)**
   - Player explores Snarky Squirrel Commons normally
   - Uses SEARCH command on ceiling (DC 12)
   - Discovers hidden trap door with flavor text
   - Can open/climb without combat required

2. **Safe Arrival (Room 4001)**
   - Safe chamber for regrouping
   - Opportunity to prepare before going further
   - No mobs spawn here
   - Can return here to rest if needed

3. **Ascent Phase (Rooms 4002-4003)**
   - First squirrels encountered (2-3 each room)
   - Mobs non-aggressive initially but will attack if player initiates
   - First acorns discovered as resources
   - Squirrels use `callforhelp` to summon allies during combat
   - Players learn swarm mechanics

4. **Branching Phase (Rooms 4005-4006)**
   - More dense squirrel populations (2 each)
   - Maximum acorn caches available
   - Risk: Getting surrounded by 4+ squirrels if not careful
   - Recovery opportunity if players pick up acorns
   - Choice: go back, go forward, or explore both branches

5. **Optional Boss Phase (Room 4007)**
   - Highest density (3-4 regular squirrels minimum)
   - 30% chance Squirrel King appears
   - Heavy rewards (acorn cache + possible Crown)
   - Extremely dangerous encounter
   - Completely optional - doesn't block other content

### Combat Mechanics

**Swarm Tactic System:**
- Individual Angry Squirrels use `callforhelp 2:angry squirrel:shriek flavor` command
- This summons up to 2 adjacent squirrels from neighboring rooms
- One squirrel becomes 4+ very quickly
- Teaches players: "Don't pull recklessly, manage aggro carefully"

**Resource Loop:**
- Players take damage from squirrel swarms
- ACORNs found in rooms heal 2 HP each
- Creates tension: "Do I have enough acorns to survive further?"
- Encourages backtracking/exploration to find healing items
- Healing via acorns feels earned rather than automatic

**Difficulty Scaling:**
- Solo Level 5+ character: Challenging, requires careful strategy and tactics
- Solo Level 6+ character: Difficult but manageable, swarms still dangerous
- Group of 2-3 (Level 5+): Can handle full zone with reasonable caution
- Careless aggro pulls: Deadly even for experienced groups
- Squirrel King (optional): Boss-tier challenge, requires preparation or group support

---

## Design Rationale

### Why This Zone Works

1. **Hidden Discovery** - Trap door in Room 1004 requires search skill, rewards curiosity and exploration
2. **Vertical Progression** - Unique gameplay compared to typical horizontal dungeon layouts
3. **Swarm Mechanic** - Emphasizes the "squirrel army" concept with BIGRAT-tier difficulty
4. **Resource Scarcity** - Acorns are rewards for brave exploration, not guaranteed survival items
5. **Optional Boss** - Squirrel King is completely optional, doesn't block zone access or progression
6. **Thematic Consistency** - All text emphasizes squirrel territorial behavior and cooperation
7. **Level Appropriateness** - Matches difficulty of Big Rat (Level 5) mobs with escalated boss challenge
8. **Replayability** - Acorn respawns and 30% boss spawn rate create different experiences
9. **Player Progression** - Zone scales with player leveling (difficulty matches current party capability)

### Boss Encounter Decision: YES to Squirrel King

**Arguments For Including a Boss:**
- Provides a "true ending" for zone exploration
- Gives veteran players something to aim for
- Rare 30% spawn adds mystery and replayability
- Loot drop (Crown of Acorns) serves as trophy
- Can be completely avoided (not mandatory)
- Perfectly fits "Snarky Squirrel" lore - has a KING who rules

**Alternative (Not Recommended):** Could skip boss and make Room 4007 just a high-density chamber with max acorn cache. However, the King adds personality and narrative closure to the zone.

---

## Technical Implementation

### IDs to Reserve

| Type | ID Range | Purpose |
|------|----------|---------|
| Rooms | 4001-4007 | Squirrel Tree zone |
| Mobs | 82-83 | Angry Squirrel, Squirrel King |
| Items | 100-101 | Crown of Acorns, ACORN |

### Zone Configuration File

Create `_datafiles/world/default/rooms/squirrel_tree/zone-config.yaml`:

```yaml
name: Squirrel Tree
roomid: 4001  # Default entry room
autoscale:
  minimum: 1
  maximum: 1  # No autoscaling - intended as low-level zone with fixed difficulty
idlemessages:
  - "The tree trembles with the activity of hundreds of tiny squirrels."
  - "You hear chittering in the distance, punctuated by acorn crunches."
  - "The scent of fresh acorns fills the air."
musicfile: null  # Or "static/audio/music/forest-ambience.mp3" if available
defaultbiome: natural
```

### Scripting Opportunities

**File: mobs/squirrel_tree/scripts/82-angry_squirrel.js**
```javascript
// Varied chittering and behavioral messages
function onIdle(mob, room) {
  // Squirrels fighting each other for acorns
  // Territorial displays
}

function onHurt(mob, room, eventDetails) {
  // Increased aggression when damaged
  // Uses callforhelp immediately
}

function onDie(mob, room, eventDetails) {
  // Death cry attracts others
  // "The squirrel lets out a final indignant squeak!"
}
```

**File: mobs/squirrel_tree/scripts/83-squirrel_king.js**
```javascript
function onLoad(mob) {
  // Boss fanfare on spawn
}

function onEnter(user, room) {
  // Dramatic flavor text when player encounters throne
}

function onCombat(mob, user, room) {
  // Squirrel King combat messages
  // Calls 4 squirrels instead of 2
}

function onDie(mob, room, eventDetails) {
  // Broadcast victory to zone
  // Crown drops
}
```

**File: rooms/squirrel_tree/4004.js (Central Hub)**
```javascript
function onEnter(user, room) {
  // Flavor: "You've reached the heart of the tree..."
}

function onCommand_search(rest, user, room) {
  // If player searches in 4004, hint at secret passage to 4007
  // "Among the twisted roots, you notice a narrow passage leading further UP..."
}
```

**File: rooms/squirrel_tree/4007.js (Pinnacle)**
```javascript
function onLoad(room) {
  // Conditionally spawn Squirrel King (30% chance)
  // var roll = UtilDiceRoll(1, 100);
  // if (roll <= 30) { SpawnMob(83, ...) }
}

function onEnter(user, room) {
  // Dramatic entry flavor
  // "You emerge into the topmost chamber of the tree..."
}
```

---

## Testing Checklist

### Connection & Access
- [ ] Trap door is discoverable in Room 1004 via search skill (DC 12)
- [ ] Trap door opens when unlocked
- [ ] Walking through connects to Room 4001
- [ ] Ceiling exit in 4001 returns to Room 1004

### Room Structure
- [ ] All 7 rooms created with proper exits
- [ ] Room connections form proper vertical tree structure
- [ ] Exit descriptions make sense contextually
- [ ] All room descriptions load and display correctly
- [ ] Idle messages display in rooms

### Mob Spawning
- [ ] Angry Squirrels spawn in correct quantities per room
  - Room 4002: 2-3
  - Room 4003: 2-3
  - Room 4004: 1-2
  - Room 4005: 2
  - Room 4006: 2
  - Room 4007: 3-4 (minimum, plus Squirrel King 30%)
- [ ] Squirrel King spawns 30% of the time in Room 4007
- [ ] Mobs spawn with correct level and stats
- [ ] Mobs have proper descriptions and names

### Combat Mechanics
- [ ] Squirrel mobs don't aggro until player initiates combat
- [ ] Squirrel mobs use callforhelp during combat
- [ ] callforhelp summons adjacent squirrels correctly
- [ ] Squirrel King uses callforhelp 4 (higher count)
- [ ] Mobs drop gold and items per dropchance
- [ ] Squirrel King drops Crown of Acorns

### Item System
- [ ] ACORNs spawn in correct rooms
  - Room 4003: 1x (50%)
  - Room 4005: 2-3x (70% each)
  - Room 4006: 2-3x (70% each)
  - Room 4007: 3-4x (100%)
- [ ] ACORNs can be picked up and carried
- [ ] ACORNs heal 2 HP when used
- [ ] ACORNs respawn on zone reset (30 min recommended)
- [ ] Crown of Acorns drops from Squirrel King
- [ ] Crown of Acorns can be picked up and carried

### Difficulty & Balance
- [ ] Zone is challenging but doable for Level 1-2 solo
- [ ] Zone is comfortable for groups of 2-3
- [ ] Swarm tactics make solo difficult if caught unprepared
- [ ] Acorn scarcity creates meaningful resource decisions
- [ ] Squirrel King is significantly harder than normal squirrels
- [ ] High-risk/high-reward trade-offs work as intended

### Flavor & Polish
- [ ] All mob descriptions display correctly
- [ ] All room descriptions display correctly
- [ ] Idle messages don't repeat excessively
- [ ] Combat messages feel thematic
- [ ] Player feedback is clear when searching for trap door
- [ ] No typos or formatting errors in text

---

## Future Expansion Ideas

### Related Enhancements
- **Related Quest:** "Squirrel King's Crown" - NPC wants Crown of Acorns for reward
- **Faction System:** Join the "Woodland Protectors" faction (anti-squirrel) or "Acorn Hoarders" faction
- **Lore Expansion:** History of how squirrels invaded this tree (hint in idle messages)
- **Seasonal Variant:** Winter version where fewer acorns spawn (harder difficulty)
- **Easter Egg:** If players spend too long in the zone, they can trigger a "Squirrel Migration" event

### Balance Tweaks Based on Feedback
- If zone is too easy: Increase squirrel count or add a second boss type
- If zone is too hard: Add additional acorn spawns or reduce squirrel health
- If players ignore the zone: Add NPCs in Room 1004 that hint at the treasure above

---

## References & Related Content

- **Connected Location:** Room 1004 - Snarky Squirrel Commons
- **Similar Zones:** Crystal Caves (exploration + resource gathering), Bladeworks Foundry (swarm tactics)
- **Mob Design Patterns:** Based on existing rat mobs (1-rat, 12-big_rat) and group mechanics
- **Level Progression:** Fits between Frostfang (L1-5) and higher content

---

**Status:** DESIGN COMPLETE - Ready for Implementation

This document provides all specifications needed to implement the Squirrel Tree mini-zone. Create the GitHub issue with this content to track implementation progress.
