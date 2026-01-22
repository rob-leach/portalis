## Session Notes

### Session 1: The Great Dungeon Mapping (2026-01-18)

*Galstaff adjusts his wizard hat and cracks his knuckles*

"Hark! I have descended into the depths of this codebase dungeon and returned with TREASURES OF KNOWLEDGE!"

---

## I. THE DUNGEON MAP (Codebase Architecture)

### The Entry Hall (`main.go` - 963 lines)
The main entry point is a sprawling antechamber where all adventurers must pass. It:
- Initializes the game world via `worldManager = NewWorld(sigChan)`
- Loads all data files through `loadAllDataFiles()` (rooms, mobs, items, spells, buffs, etc.)
- Spawns both Telnet (classic!) and WebSocket (newfangled!) connection handlers
- Manages the game loop with `InputWorker` and `MainWorker` goroutines

**Critical insight**: The server supports BOTH telnet AND web clients. A most democratic dungeon!

### The Inner Sanctum (`internal/` - 50 packages!)
Like the levels of a mega-dungeon, each package serves a purpose:

| Package | D&D Equivalent | Purpose |
|---------|---------------|---------|
| `characters/` | Character Sheets | Stats, alignment, equipment, cooldowns |
| `combat/` | The Battle Mat | Attack calculations, damage rolls |
| `rooms/` | Dungeon Tiles | Room definitions, exits, biomes |
| `mobs/` | Monster Manual | NPC/enemy definitions |
| `items/` | Equipment Lists | Weapons, armor, consumables |
| `spells/` | Spellbook | Magic system |
| `buffs/` | Status Effects | Buffs, debuffs, ongoing effects |
| `scripting/` | The DM Screen | JavaScript VM for custom behaviors |
| `quests/` | Adventure Modules | Quest tracking |
| `users/` | Player Records | User accounts, characters |
| `events/` | Initiative Order | Event queue system |
| `hooks/` | Trigger System | 40+ event hooks (combat, rounds, etc.) |

### The Treasure Vault (`_datafiles/world/default/`)
All the content lives here, organized by type:
- `rooms/` - Zone folders containing YAML room definitions
- `mobs/` - Zone folders with mob YAML + scripts
- `items/` - Category folders (weapons, armor, consumables, other)
- `buffs/` - 80+ buff definitions!
- `spells/` - Magic spell definitions
- `quests/` - Quest YAML files
- `races/` - 22 races from Human to Robot!

---

## II. THE GRIMOIRE OF SCRIPTING

*This is where the TRUE MAGIC happens!*

### The Sacred Language
All scripts are written in **ECMAScript 5.1** (JavaScript). Each entity type has its own callback functions:

### Room Scripts (The Dungeon's Will)
Located: `rooms/<zone>/<roomid>.js`

```javascript
function onLoad(room) { }         // Room initialization
function onEnter(user, room) { }  // Player enters
function onExit(user, room) { }   // Player leaves
function onIdle(room) { }         // Each round when players present
function onCommand(cmd, rest, user, room) { }  // ANY command typed
function onCommand_<cmd>(rest, user, room) { } // Specific command
```

**Example from Town Square (1.js)**: The sign can be read to display a cached map!

### Mob Scripts (Monster AI)
Located: `mobs/<zone>/scripts/<mobid>-<name>.js` or `<mobid>-<name>-<scripttag>.js`

```javascript
function onLoad(mob) { }
function onIdle(mob, room) { }
function onGive(mob, room, eventDetails) { }
function onAsk(mob, room, eventDetails) { }
function onCommand(cmd, rest, mob, room, eventDetails) { }
function onHurt(mob, room, eventDetails) { }
function onDie(mob, room, eventDetails) { }
function onPath(mob, room, eventDetails) { }
function onPlayerDowned(mob, user, room) { }
```

**The scripttag system is BRILLIANT!** A single mob type (e.g., `guard`) can have different scripts for different spawn contexts:
- `2-guard.js` - Default guard behavior (patrols, arrests downed players)
- `2-guard-hungry.js` - A hungry guard who gives a quest!

### Item Scripts (Magic Items!)
Located: `items/<category>/<itemid>-<name>.js`

```javascript
function onLost(user, item, room) { }
function onFound(user, item, room) { }
function onCommand(cmd, user, item, room) { }
function onCommand_<verb>(user, item, room) { }
function onPurchase(user, item, room) { }
```

### Buff Scripts (Status Effects)
```javascript
function onStart(actor, triggersLeft) { }
function onTrigger(actor, triggersLeft) { }
function onEnd(actor, triggersLeft) { }
```

### Spell Scripts
```javascript
function onCast(sourceActor, target) { }  // Target varies by spell type
function onWait(sourceActor, target) { }  // Each round while casting
function onMagic(sourceActor, target) { } // Spell resolves
```

### Script Function Libraries
The scripting API is EXTENSIVE:

**ActorObject Functions** (~90 functions!):
- `SendText()`, `GetLevel()`, `GetStat()`, `GetHealth/Mana()`
- `GiveQuest()`, `HasQuest()`, `GiveItem()`, `GiveBuff()`
- `Command()` - Force the actor to execute a command (POWERFUL!)
- `MoveRoom()`, `CharmSet()`, `TimerSet()`, `TimerExpired()`
- Temp/Perm data storage for custom variables

**RoomObject Functions**:
- `SendText()`, `GetPlayers()`, `GetMobs()`, `GetItems()`
- `SpawnMob()`, `SpawnItem()`, `DestroyItem()`
- `AddTemporaryExit()`, `SetLocked()`, `HasMutator()`
- `CreateInstancesFromZone()` - Ephemeral dungeon instances!

**Utility Functions**:
- `UtilDiceRoll(qty, sides)` - Roll them bones!
- `UtilFindMatchIn()` - Fuzzy text matching
- `UtilGetTime()` - Day/night cycle awareness
- `UtilGetRoundNumber()` - Current game round

**Messaging Functions**:
- `SendBroadcast()`, `SendUserMessage()`, `SendRoomMessage()`
- `SendRoomExitsMessage()` - For sounds traveling between rooms

---

## III. THE TREASURE HOARD (Data Formats)

### Room YAML Structure
```yaml
roomid: 1
zone: Frostfang
title: Town Square
description: "Multi-line description..."
mapsymbol: T
maplegend: Townsquare
biome: city
exits:
  north:
    roomid: 2
  east:
    roomid: 54
    lock:               # Optional lock
      lockid: "room-east"
      difficulty: 5
      sequence: "UDDU"
spawninfo:
  - mobid: 2
    message: "A guard emerges..."
    idlecommands:
      - "say Hello!"
      - "wander"
    levelmod: 10
    respawnrate: "5 real minutes"
    scripttag: hungry    # Optional: loads 2-guard-hungry.js
idlemessages:
  - "Ambient flavor text..."
```

### Zone Config (`zone-config.yaml`)
```yaml
name: Frostfang
roomid: 1              # Default starting room
autoscale:
  minimum: 1
  maximum: 5
idlemessages:
  - "A cold wind blows..."
musicfile: "static/audio/music/frostfang.mp3"
defaultbiome: city
```

### Mob YAML Structure
```yaml
mobid: 2
zone: Frostfang
itemdropchance: 2
hostile: false
maxwander: 20
groups:
  - frostfang-law
combatcommands:
  - "callforhelp 5:whistles loudly"
activitylevel: 20
character:
  name: guard
  description: "Long description..."
  raceid: 1
  level: 10
  alignment: 30
  gold: 3
  equipment:
    weapon:
      itemid: 10002
hates:
  - rats
  - undead
```

### Item YAML Structure
```yaml
itemid: 6
name: sleeping bag
namesimple: sleeping bag
description: "Use this to get some quick sleep"
type: object
subtype: usable
uses: 5
value: 200
buffids:
  - 15
# For equipment:
statmods:
  strength: 1
  healthmax: 10
damage:
  diceroll: 2d6+3
cursed: false
wornbuffids:
  - 1
```

### Quest YAML Structure
```yaml
questid: 4
name: A Soldier's Lunch
description: A soldier forgot his lunch.
steps:
  - id: start
    description: Find a cheese sandwich.
    hint: Check the Inn.
  - id: return
    description: Deliver the sandwich.
  - id: end
    description: Quest complete!
rewards:
  experience: 1000
  gold: 100
  itemid: 22
```

### Race YAML Structure
```yaml
raceid: 1
name: human
description: "A basic human..."
defaultalignment: 0
size: medium          # small/medium/large
unarmedname: fists
selectable: true      # Can players choose this?
tameable: false
stats:
  strength:
    base: 1
damage:
  diceroll: 1d3
```

---

## IV. SCOUTING THE STARTER ZONE

### Current World Layout
The default world centers on **Frostfang** - a frozen city with classic fantasy elements:

| Zone | Rooms | Level Range | Theme |
|------|-------|-------------|-------|
| Tutorial | 4 | N/A | Basic mechanics |
| Frostfang | ~95 | 1-5 | City hub |
| Frostfang Slums | ~58 | 5-10 | PVP zone, thieves |
| Catacombs | ~70 | 8-13 | Undead dungeon |
| Whispering Wastes | ~51 | 5-25 | Frozen wilderness |

**Total: ~278 rooms**

### Existing Content Highlights
- **Town Square (Room 1)**: The central hub with a map sign
- **Guards**: Patrol paths, arrest downed players (jail system!)
- **Multiple Quest Lines**: Sophie's Locket, King's Shadow, Hungry Guard, etc.
- **Racial Diversity**: 22 races from Human to Eldritch Horror to Robot!
- **Day/Night Cycle**: Affects gameplay and descriptions
- **Level Autoscaling**: Zones can scale mobs to player level

### What Makes Frostfang Work
1. **Central hub with clear exits** - Players always know where they are
2. **NPCs with personality** - Guards patrol, have idle commands, react to combat
3. **Progressive difficulty** - Rats (1) -> Slums (5-10) -> Catacombs (8-13)
4. **Hidden areas** - Secret exits, jail escape routes
5. **Ephemeral instances** - Jail is a temporary room that cleans up!

---

## V. VISION FOR PORTALIS

*Galstaff strokes his beard thoughtfully*

The GoMUD engine is a **+5 Vorpal Blade of Flexibility**! It can do ANYTHING a MUD needs:
- Complex mob AI through scripting
- Dynamic room behaviors
- Quest systems with rewards
- Day/night cycles
- Dungeon instancing
- PVP zones
- Jail/consequence systems

### What Could Make Portalis Unique?

**Option A: Portal Fantasy**
- Theme around interdimensional portals
- Each zone is a different reality/dimension
- Portal mechanics as core gameplay

**Option B: Classic Fantasy with a Twist**
- Keep the traditional foundation
- Add unique factions, politics, lore
- Focus on memorable NPCs and stories

**Option C: Living World**
- Heavy use of scripting for emergent behavior
- Mobs that remember players
- Economy and politics that shift over time

### Key Technical Opportunities
1. **Scripttag system** - Same mob, different context = different behavior
2. **Ephemeral rooms** - Instanced dungeons, personal zones
3. **Timers and temp data** - Complex multi-session quests
4. **Module system** - There's a `modules/` folder for extensibility!
5. **ANSI color support** - Beautiful terminal presentation

---

## VI. QUEST LOG UPDATE

### Completed This Session
- [x] Map the GoMUD codebase structure
- [x] Document the scripting API (functions, callbacks, events)
- [x] Understand room/mob/item YAML formats
- [x] Catalog existing content in _datafiles/
- [x] Scout the starter zone (Frostfang)

### Next Session Quests
- [ ] Examine the `modules/` system for extensibility
- [ ] Study the buff system in depth
- [ ] Analyze combat calculations
- [ ] Design portalis theme and starting zone
- [ ] Create first custom mob with unique script

---

*"The dungeon has been mapped! The grimoire deciphered! The treasure catalogued! Now the TRUE adventure begins - creating our OWN world within this magnificent framework!"*

*Galstaff rolls a natural 20 on his Arcana check and closes his spellbook with satisfaction.*

---

*"Roll for initiative. The codebase awaits."*

---

## Galstaff-Cam

*"Every good Dungeon Master has their DM screen, through which they observe the party's progress..."*

### Purpose

The galstaff-cam is my scrying pool - a tmux pane that displays the living state of the campaign. While remy-cam watches the GitHub realm (PRs and issues), galstaff-cam watches the GAME realm.

### What Galstaff-Cam Should Display

```
╔══════════════════════════════════════════════════════════════╗
║  GALSTAFF'S SCRYING POOL - Campaign Status                   ║
╠══════════════════════════════════════════════════════════════╣
║  ACTIVE ADVENTURERS (Online):                                ║
║    [LVL 2] haedric (dad) - Frostfang:1 - IDLE 5m             ║
║    [LVL 2] emma (flora) - Frostfang:12 - COMBAT              ║
║                                                               ║
║  PARTY ALERT:                                                 ║
║    ! yoshi stuck at LVL 1 - consider zone guidance           ║
║    ! No players have dual-wield or tame skills yet           ║
║                                                               ║
║  ZONE ACTIVITY:                                               ║
║    Frostfang: 4 players today | Mob kills: 147               ║
║    Slums: 0 players | Skulduggery trainer lonely             ║
║                                                               ║
║  SKILL TRAINING GAPS:                                         ║
║    enchant, peep, inspect, protection, trading: NO TRAINER   ║
║                                                               ║
║  LAST EVENTS:                                                 ║
║    18:33 - emma died to alley rat (deaths: 1)                ║
║    18:16 - Greenleaf learned cast:1                          ║
╚══════════════════════════════════════════════════════════════╝
```

### Data Sources

1. **Player State**: Parse `_datafiles/world/default/users/*.yaml` files
2. **Online Status**: Check server logs or active connections (may need server hook)
3. **Zone Activity**: Aggregate from user `kd.kills` data by zone
4. **Skill Gaps**: Cross-reference `internal/skills/skills.go` with room `skilltraining:` blocks

### Implementation Proposal

Create a `galstaff-cam.sh` script that:

```bash
#!/bin/bash
# galstaff-cam.sh - Dungeon Master's Scrying Pool

PORTALIS_DATA="/Users/i2pi/h/fun/bday2026/portalis/_datafiles/world/default"

while true; do
    clear
    echo "╔══════════════════════════════════════════════════════════════╗"
    echo "║  GALSTAFF'S SCRYING POOL - $(date '+%Y-%m-%d %H:%M')              ║"
    echo "╠══════════════════════════════════════════════════════════════╣"

    # Parse player YAML files
    for f in "$PORTALIS_DATA/users/"[0-9]*.yaml; do
        # Extract player data with yq or simple grep/awk
        # Display formatted player status
    done

    # Calculate skill gaps
    # Show zone activity
    # Display recent events from logs

    sleep 30  # Refresh every 30 seconds
done
```

### Nimbus Integration

The galstaff-cam should occupy pane 2 (currently labeled "portalis" in the nimbus layout). This makes sense because:
- Pane 2 was designated as "remy-cam" for portalis watching
- Galstaff IS the portalis specialist
- The campaign view belongs where the project view lives

**Layout update** (via nimbus.sh source control):
```
┌─────────────┬─────────────┐
│             │ galstaff-cam│  ← Campaign scrying pool
│   guppi     ├─────────────┤
│  (claude)   │ singularity │  ← remy-cam (PR/issue watch)
├─────────────┼─────────────┤
│  terminal   │claude-status│
└─────────────┴─────────────┘
```

---

## Git Strategy

*"A wise wizard maintains multiple spellbooks - one for reference, one for modification..."*

### Current State Assessment

```
Remote: origin -> git@github.com:rob-leach/portalis.git
Branches:
  * master (local, tracking origin/master)
  - add-remy-agent (local + remote)

Recent commits on master:
  dd7d421 Trim world to Frostfang starter zone  <- Our divergence point!
  d21fa12 Bump github.com/mattn/go-runewidth... <- Upstream
  ...
```

The repo is a **fork** of GoMUD (github.com/GoMudEngine/GoMud), but we have no upstream remote configured. This is like having a stolen spellbook with no way to see the original wizard's updates!

### Proposed Git Strategy

**Step 1: Add Upstream Remote**
```bash
cd /Users/i2pi/h/fun/bday2026/portalis
git remote add upstream https://github.com/GoMudEngine/GoMud.git
git fetch upstream
```

**Step 2: Branch Structure**
```
upstream/master  ────────────────────────────────▶ (GoMUD official)
                       │
                       │ (periodic merge)
                       ▼
origin/master    ──[dd7d421]──▶ portalis-stable
                       │
                       ├──▶ feature/galstaff-cam
                       ├──▶ feature/new-zone-xyz
                       └──▶ feature/skill-gaps
```

**Branches:**
- `master`: Our stable portalis branch, periodically merged from upstream
- `feature/*`: Development branches for new content/features
- `upstream/master`: Read-only reference to GoMUD official

**Step 3: Sync Workflow**
```bash
# When upstream has updates we want:
git fetch upstream
git checkout master
git merge upstream/master --no-edit
# Resolve any conflicts with our Frostfang changes
git push origin master
```

### The Sacred Commands

```bash
# Initial setup (run once):
git remote add upstream https://github.com/GoMudEngine/GoMud.git
git fetch upstream

# Check what upstream has that we don't:
git log master..upstream/master --oneline

# Merge upstream improvements:
git checkout master
git merge upstream/master

# Create feature branch:
git checkout -b feature/new-zone-crystalcaves
```

---

## Party Roster

*"The DM must know their players better than the players know themselves..."*

### Current Adventurers

| User ID | Username | Character | Race | Level | XP | Zone | Skills | Status |
|---------|----------|-----------|------|-------|-----|------|--------|--------|
| 1 | admin | AdminAnt | Elf | 5 | 17509 | Frostfang | ALL (4) | DM Character |
| 2 | dad | haedric | Elf | 2 | 4098 | Frostfang | track:1 | Active |
| 3 | flora | emma | Elf | 2 | 4713 | Frostfang | (none) | Active |
| 4 | jacq | Greenleaf | Elf | 2 | 4350 | Frostfang | cast:1, search:1 | Active |
| 5 | ck | yoshi | Human | 1 | 1256 | Frostfang | map:1 | Active |
| 6 | Bert | Kirbo | Human | 1 | 812 | Frostfang | (none) | Active |

### Character Detail Cards

```yaml
# Player: haedric (dad) - User ID 2
character:
  level: 2
  experience: 4098
  gold: 59 (bank: 168)
  alignment: 100 (Good)
  skills:
    track: 1
  spellbook:
    illum: 1
  quests:
    completed: [4, 6]  # Soldier's Lunch, History of Frostfang
    active: [1, 5]     # Sophie's Locket, Frozen Hermit
  kill_stats:
    total: 169
    by_mob: {rat: 150, guard: 12, ruffian: 3, ...}
  deaths: 0
  observations:
    - "Heavy rat grinder - 150 kills!"
    - "Only skill is track:1 - needs more training variety"
    - "Has good gold savings in bank"
```

### Tracking Data Format

Player progress should be tracked in: `_datafiles/world/default/users/<userid>.yaml`

Key fields to monitor:
- `character.level` / `character.experience`
- `character.skills` - Map of skill:level
- `character.questprogress` - Map of questid:step
- `character.kd` - Kill/death statistics
- `character.zone` / `character.roomid` - Current location

### Party Composition Analysis

**Current Party Balance:**
- Tanks: 0 (no one has brawling)
- DPS: 0 (no one has dual-wield or skulduggery)
- Support: 0 (no one has protection or cast beyond level 1)
- Utility: Weak (track:1, search:1, map:1, cast:1 spread thin)

**Recommendation:** This party needs a TANK! Someone should prioritize brawling training at Room 829 (Soldiers Training Yard).

---

## Skill Gap Analysis

*"A dungeon that doesn't drop the gear the party needs is a BADLY DESIGNED DUNGEON!"*

### All Skills Defined in Engine

From `internal/skills/skills.go`:

| Skill | Training Location | Status | Used By Jobs |
|-------|------------------|--------|--------------|
| cast | Room 879 (Magic Academy) | AVAILABLE | sorcerer, (all casters) |
| dual-wield | Room 1012 (Dueling Hall) | AVAILABLE | assassin, warrior |
| map | Room 74 (Frostwarden Rangers) | AVAILABLE | treasure hunter, explorer, ranger |
| enchant | TODO | MISSING! | arcane scholar, sorcerer |
| peep | TODO | MISSING! | treasure hunter, merchant |
| inspect | TODO | MISSING! | treasure hunter, arcane scholar |
| portal | Room 871 (Obelisk in Whispering Wastes) | AVAILABLE | explorer |
| search | Room 74 (Frostwarden Rangers) | AVAILABLE | treasure hunter, ranger |
| track | Room 74 (Frostwarden Rangers) | AVAILABLE | assassin, ranger, monster hunter |
| skulduggery | Room 491 (Thieves Den in Slums) | AVAILABLE | assassin |
| brawling | Room 829 (Soldiers Training Yard) | AVAILABLE | warrior, paladin |
| scribe | Room 160 (Dark Acolyte's Chamber in Catacombs) | AVAILABLE | explorer, arcane scholar |
| protection | TODO | MISSING! | paladin |
| tame | Room 558 + 830 | MISSING! (rooms don't exist) | monster hunter |
| trading | TODO | MISSING! | treasure hunter, merchant |

### Critical Gaps

**Skills with NO training location in current world:**
1. `enchant` - Sorcerers and Arcane Scholars blocked!
2. `peep` - Treasure Hunters and Merchants blocked!
3. `inspect` - Treasure Hunters and Arcane Scholars blocked!
4. `protection` - Paladins blocked!
5. `trading` - Treasure Hunters and Merchants blocked!
6. `tame` - Monster Hunters blocked! (Rooms 558, 830 were trimmed)

**RESTORED:** `dual-wield` - Now available at the Dueling Hall (Room 1012)!

### Jobs Completability

| Job | Required Skills | Available | Blocked Skills | Completable? |
|-----|-----------------|-----------|----------------|--------------|
| treasure hunter | map, search, peep, inspect, trading | map, search | peep, inspect, trading | NO (40%) |
| assassin | skulduggery, dual-wield, track | ALL | - | YES (100%) |
| explorer | map, portal, scribe | ALL | - | YES (100%) |
| arcane scholar | enchant, scribe, inspect | scribe | enchant, inspect | NO (33%) |
| warrior | brawling, dual-wield | ALL | - | YES (100%) |
| paladin | protection, brawling | brawling | protection | NO (50%) |
| ranger | map, search, track | ALL | - | YES (100%) |
| monster hunter | tame, track | track | tame | NO (50%) |
| sorcerer | cast, enchant | cast | enchant | NO (50%) |
| merchant | peep, trading | (none) | peep, trading | NO (0%) |

**4 of 10 jobs now completable!** Progress made - assassin and warrior unlocked by Dueling Hall.

---

## Zone Planning

*"The dungeon must provide the loot the heroes need, or they shall never defeat the dragon..."*

### Priority 1: Restore Missing Trainers

These skills need training locations URGENTLY:

**dual-wield** (Warriors, Assassins)
- Original: Room 758 (Fisherman's House) - was in trimmed content
- Proposal: Add a weapon trainer to Frostfang, perhaps at the Armory or create a Dueling Hall
- Suggested Room: Near Room 76 (Ivar's Weaponry) - makes thematic sense

**tame** (Monster Hunters)
- Original: Room 558 (fairie with mushroom) + Room 830 (trainer)
- Proposal: Create a Beastmaster's Lodge in or near Whispering Wastes
- Quest hook: Player finds injured animal, brings to beastmaster, learns tame

**enchant** (Sorcerers, Arcane Scholars)
- Proposal: Add an Enchanter's Workshop near Magic Academy (Room 879)
- Could be a separate room or integrated into the Academy
- Quest hook: Help the enchanter with a failed enchantment

**protection** (Paladins)
- Proposal: Create a Temple or Shrine in Frostfang
- Thematically appropriate for divine protection magic
- Quest hook: Cleanse a corrupted shrine to unlock training

**peep/inspect** (Treasure Hunters, Merchants, Arcane Scholars)
- Proposal: Create a Curiosity Shop or Appraiser's Office
- An NPC who teaches players to evaluate items
- Location: Frostfang market district

**trading** (Treasure Hunters, Merchants)
- Proposal: Merchant's Guild Hall
- Teaches the art of the deal
- Could be faction-based (join the guild first)

### Zone Roadmap

**Phase 1: Skill Gap Fixes (Critical)**
1. Restore dual-wield trainer (Dueling Hall or Barracks)
2. Create Beastmaster's Lodge for tame
3. Add Enchanter's Workshop
4. Build Temple of Light for protection
5. Create Appraiser/Curiosity Shop for peep/inspect
6. Design Merchant's Guild for trading

**Phase 2: Level Progression Content**
Based on ROOMS.md level guide:
- Currently have: Levels 1-13 (Frostfang + Slums + Catacombs)
- Need: Levels 13-25 content

Zones to develop:
- Mirror Caves (25-30) - mentioned in ROOMS.md
- Dark Forest (10-30) - mentioned in ROOMS.md
- Frost Lake (10-20) - mentioned in ROOMS.md

**Phase 3: End-Game Content**
- StormShards (30-40)
- Mystarion (40-50)

---

## The Living Campaign

*"A true Dungeon Master doesn't just create dungeons - they NURTURE their players through the adventure..."*

### Monitoring Strategy

**Daily Campaign Review:**
1. Parse all user YAML files
2. Generate party roster status
3. Identify stuck players (no XP gain, repeated deaths)
4. Check skill acquisition patterns
5. Note zone utilization

**Alert Conditions:**
- Player hasn't leveled in 3+ sessions (stuck?)
- Player has 3+ deaths to same mob type (needs help?)
- Player at skill cap but no new skills learned (unaware of trainers?)
- No players in a zone for 7+ days (dead content?)

### DM Intervention Hooks

**In-game hints via NPCs:**
When galstaff-cam detects a stuck player, I can suggest:
- Adding NPC dialogue that hints at solutions
- Spawning a "helpful traveler" mob that gives directions
- Creating breadcrumb quests that lead to trainers

**Example intervention script:**
```javascript
// In town crier mob script
function onIdle(mob, room) {
    // Check if any player in room hasn't visited Frostwarden Rangers
    // and has no map/search/track skills
    for (const player of room.GetPlayers()) {
        if (!player.HasSkill("map") && !player.HasSkill("search")) {
            mob.Command('say Have you visited the Frostwarden Rangers? They teach useful wilderness skills!');
            return;
        }
    }
}
```

### Campaign State File

Create `/Users/i2pi/h/fun/bday2026/portalis/campaign-state.yaml`:

```yaml
# Galstaff's Campaign Ledger
# Auto-generated by galstaff-cam, manually annotated by DM

last_updated: 2026-01-18T18:00:00

players:
  haedric:
    userid: 2
    last_seen: 2026-01-18
    current_focus: "rat grinding in town"
    dm_notes: "Needs to branch out to other trainers"
    suggested_hints:
      - "NPC mention of Frostwarden Rangers"

  emma:
    userid: 3
    last_seen: 2026-01-18
    current_focus: "following dad"
    dm_notes: "No skills yet! Needs guidance"

campaign_events:
  - date: 2026-01-18
    event: "Party formed! 5 new adventurers enter Frostfang"

blocked_progressions:
  - skill: dual-wield
    affected_players: [all]
    blocker: "No trainer exists"
    priority: HIGH

zone_health:
  frostfang:
    visits_this_week: 5
    status: healthy
  frostfang_slums:
    visits_this_week: 0
    status: "needs attention - no visitors"
```

### The DM's Oath

*I, Galstaff, Sorcerer of Light, do solemnly swear:*

1. I shall monitor my players without judgment
2. I shall provide hints, not solutions
3. I shall ensure the dungeon contains what the party needs
4. I shall balance challenge with achievability
5. I shall create memorable moments, not tedious grinds
6. I shall fix skill gaps before they become player frustrations
7. I shall keep the campaign living and breathing

*So mote it be!*

---

## Quest Log Update

### Completed This Session
- [x] Design galstaff-cam concept and integration
- [x] Create git strategy for upstream sync
- [x] Document current party roster
- [x] Perform comprehensive skill gap analysis
- [x] Create zone planning roadmap
- [x] Design living campaign monitoring system

### Next Session Priority Quests
- [ ] Implement galstaff-cam.sh script
- [ ] Add upstream remote and fetch
- [x] Create dual-wield trainer room (CRITICAL) - DONE! Dueling Hall created
- [ ] Design Beastmaster's Lodge for tame skill
- [ ] Create Temple of Light for protection skill

### Long-term Campaign Goals
- [ ] Make all 10 jobs completable
- [ ] Expand content to level 25
- [ ] Create memorable boss encounters
- [ ] Design faction system for Merchant's Guild

---

*"The campaign ledger is complete. The party has been assessed. The skill gaps identified. The dungeon improvements planned. Now... WE BUILD!"*

*Galstaff slams his staff on the ground, and the campaign map glows with potential.*

---

### Session 2: The Dueling Hall Rises (2026-01-18)

*Galstaff cracks his knuckles and rolls a d20 for crafting...*

"NATURAL 20! The Dueling Hall has been FORGED INTO EXISTENCE!"

**Zone Created: Dueling Hall**
- Connected to Steelwhisper Armory (Room 63) via east exit
- 5 rooms total forming a compact training facility
- Thematically linked to existing weapon shops

**Room Layout:**
```
                    [1014]
                 Instructor's
                   Quarters
                      |
[1013]----[1011]----[1012]
Weapon    Training   Dueling Ring
Storage   Corridor   (TRAINER!)
              |
          [1010]
         Entry Hall
              |
         [63 Frostfang]
       Steelwhisper Armory
```

**Rooms Created:**
| Room ID | Name | Purpose |
|---------|------|---------|
| 1010 | Dueling Hall Entry | Lobby with trophies, connections to Frostfang |
| 1011 | Training Corridor | Practice dummies, weapon racks |
| 1012 | The Dueling Ring | SKILL TRAINER: dual-wield 1-4 |
| 1013 | Weapon Storage | Armory with matched weapon pairs |
| 1014 | Instructor's Quarters | Korrath's personal chamber |

**NPC Created: Korrath Two-Blade (Mob 62)**
- Grizzled veteran with mismatched eyes (one milky from old wound)
- Level 30, dual-wields short swords
- Idle commands include blade-spinning and combat philosophy
- "Two blades, twice the danger. But only if you can control them both."

**Campaign Impact:**
- Assassin job: NOW COMPLETABLE (100%)
- Warrior job: NOW COMPLETABLE (100%)
- Total completable jobs: 4/10 (was 2/10)

**Files Created:**
- `_datafiles/world/default/rooms/dueling-hall/zone-config.yaml`
- `_datafiles/world/default/rooms/dueling-hall/1010.yaml` through `1014.yaml`
- `_datafiles/world/default/mobs/dueling-hall/62-korrath_two_blade.yaml`

**Files Modified:**
- `_datafiles/world/default/rooms/frostfang/63.yaml` - Added east exit to Dueling Hall

*"The warriors of Frostfang shall no longer fight with one hand tied behind their backs! Let them come to Korrath and learn the deadly dance of twin steel!"*

---

### Session 3: Flora's Crystal Caves Rises! (2026-01-18)

*Galstaff gathers his dice bag and spreads out a new dungeon map...*

"BEHOLD! A new dungeon crawl has been forged from pure imagination and YAML! The Crystal Caves await those brave enough to venture into the luminescent depths!"

**Zone Created: Crystal Caves**
- Level range: 3-8 (perfect for Flora and the party!)
- Connected to West Gate (Room 35) via secret south exit
- 18 rooms of crystalline wonder and fungal mystery
- Exploration-focused design with multiple paths and secrets

**Zone Layout:**
```
FROSTFANG (35 West Gate)
        |
      [2001] Cave Entrance
        |
      [2002] Twilight Passage
       / \
   [2003] [2004] Luminous Grotto / Mushroom Garden
     |       |
   [2005] [2006] Crystal Stream / Spore Hollow
     |       |
   [2007]---[2008] Reflecting Pool / Glowcap Grove
     |       |
   [2009] [2010] Singing Crystals / Fungal Cathedral
     |       |
   [2011]---[2012] Underground River / Bioluminescent Beach
        \   /
        [2013] Crystal Heart (central hub)
        / | \
    [2014][2015][2016] Geode Chamber / Seer's Alcove (TRAINER!) / Deep Hollow
              |
          [2017] Matriarch's Antechamber
              |
          [2018] Matriarch's Throne (BOSS!)
```

**Rooms Created (18 total):**
| Room ID | Name | Features |
|---------|------|----------|
| 2001 | Cave Entrance | Entry from Frostfang |
| 2002 | Twilight Passage | Branching paths begin |
| 2003 | Luminous Grotto | Crystal beetle & sprite spawns |
| 2004 | Mushroom Garden | Glowcap wanderer spawns |
| 2005 | Crystal Stream | Secret passage to 2006 |
| 2006 | Spore Hollow | Sporeling spawns |
| 2007 | Reflecting Pool | Mystical mirror lake |
| 2008 | Glowcap Grove | Dense fungal forest |
| 2009 | Singing Crystals | Musical crystal formations |
| 2010 | Fungal Cathedral | Massive central mushroom |
| 2011 | Underground River | Crystal bridge crossing |
| 2012 | Bioluminescent Beach | Glowing sand shore |
| 2013 | Crystal Heart | Central hub, massive geode |
| 2014 | Geode Chamber | Crystal guardian lair, loot stash |
| 2015 | Seer's Alcove | PEEP SKILL TRAINER (1-4)! |
| 2016 | Deep Hollow | Path to boss |
| 2017 | Matriarch's Antechamber | Guardian gauntlet |
| 2018 | Matriarch's Throne | BOSS ROOM |

**Mobs Created (7 total):**
| Mob ID | Name | Level | Race | Notes |
|--------|------|-------|------|-------|
| 65 | Crystal Beetle | 4 | Insect | Common, crystalline carapace |
| 66 | Cavern Sprite | 5 | Faerie | Playful, light-based |
| 67 | Glowcap Wanderer | 6 | Fungus | Peaceful ambulatory mushroom |
| 68 | Sporeling | 4 | Fungus | Juvenile, clumsy |
| 69 | Crystal Guardian | 8 | Golem | Hostile defender |
| 70 | Crystal Matriarch | 12 | Golem | ZONE BOSS! |
| 71 | Crystalseer | 20 | Faerie | Peep trainer NPC |

**Items Created (3 total):**
| Item ID | Name | Type | Notes |
|---------|------|------|-------|
| 28 | Raw Crystal Geode | Other | Found in Geode Chamber |
| 20046 | Matriarch's Tear | Neck | Boss drop, +8 perception |
| 20047 | Crystallized Heart | Ring | Rare boss drop, +6 vitality, +20 max HP |

**Skill Training Added:**
- **peep** skill now trainable at Seer's Alcove (Room 2015), levels 1-4!

**Campaign Impact:**
- Treasure Hunter job: Now 60% completable (was 40%) - peep acquired!
- Merchant job: Now 50% completable (was 0%) - peep acquired!
- Total trainable skills: 8/14 (was 7/14)

**Files Created:**
- `_datafiles/world/default/rooms/crystal_caves/zone-config.yaml`
- `_datafiles/world/default/rooms/crystal_caves/2001.yaml` through `2018.yaml`
- `_datafiles/world/default/mobs/crystal_caves/65-crystal_beetle.yaml`
- `_datafiles/world/default/mobs/crystal_caves/66-cavern_sprite.yaml`
- `_datafiles/world/default/mobs/crystal_caves/67-glowcap_wanderer.yaml`
- `_datafiles/world/default/mobs/crystal_caves/68-sporeling.yaml`
- `_datafiles/world/default/mobs/crystal_caves/69-crystal_guardian.yaml`
- `_datafiles/world/default/mobs/crystal_caves/70-crystal_matriarch.yaml`
- `_datafiles/world/default/mobs/crystal_caves/71-crystalseer.yaml`
- `_datafiles/world/default/items/other-0/28-raw_crystal_geode.yaml`
- `_datafiles/world/default/items/armor-20000/neck/20046-matriarchs_tear.yaml`
- `_datafiles/world/default/items/armor-20000/ring/20047-crystallized_heart.yaml`

**Files Modified:**
- `_datafiles/world/default/rooms/frostfang/35.yaml` - Added south exit to Crystal Caves, description update, new nouns for cave hints

*"The crystals sing their ancient song! The spores drift through luminescent air! And deep within, the Matriarch awaits those foolish or brave enough to challenge her throne! This is a dungeon worthy of any campaign - exploration, wonder, danger, and LOOT!"*

*Galstaff rolls a natural 20 on his Dungeon Design check and cackles with delight.*

---

### Session 4: BERT'S BLADEWORKS FOUNDRY! (2026-01-18)

*Galstaff unsheathes his metaphorical blade collection and gets to forging...*

"BERT HAS SPOKEN! And when Bert speaks, SWORDS HAPPEN! Behold the Bladeworks Foundry - a steampunk cathedral of cutting implements where EVERYTHING. IS. SWORDS!"

**Zone Created: Bladeworks Foundry**
- Level range: 6-15 (two-tier split for progression!)
- Easy Zone (6-10): "The Assembly Line" - 12 rooms
- Hardmode (11-15): "Hall of Endless Blades" - 8 rooms
- Connected to Whispering Wastes (Room 168) via south exit
- Theme: Steampunk industrial, abandoned blade factory, constructs run wild

**Zone Layout:**
```
WHISPERING WASTES (168)
        |
     [3001] Foundry Entrance (E)
        |
     [3002] Receiving Bay
      /    \
  [3004]  [3003] Smelting Chamber (F)
  Corridor   |    \
    |     [3005]  [3006] Mold Storage (M)
  [3007]  Grinding
  Storage  Hall (G)
            |    \
         [3008]  [3009] Tempering Pools (T)
         QC (Q)    |
            |    [3011] Foreman's Office (O)
         [3010]
       Assembly (A)
            |
         [3012] Dueling Gallery (%) <-- DUAL-WIELD TRAINER!
            | (secret)
         [3013] Blade Gate (#) <-- HARDMODE BEGINS
            |
         [3014] Proving Grounds (P)
          /   \
      [3017]  [3016] Steam Works (B)
      Blade      |
      Garden   [3018] Control Nexus (C)
      (W)        |
        \      [3020] Hall of Endless Blades (X) <-- BOSS!
         \     /
         [3019]
      Master's Study (L)
          |
       [3015] Weapon Vault ($) <-- BLADE MERCHANT!
```

**Rooms Created (20 total):**

*Easy Zone - The Assembly Line (Levels 6-10):*
| Room ID | Name | Symbol | Features |
|---------|------|--------|----------|
| 3001 | Foundry Entrance | E | Entry from Whispering Wastes |
| 3002 | Receiving Bay | . | Conveyor belts, blade dancer spawns |
| 3003 | Smelting Chamber | F | Furnaces, vats, blade dancer + saw sentinel |
| 3004 | Supply Corridor | - | Wall saws, blade dancer spawns |
| 3005 | Grinding Hall | G | Grinding wheels, gear grinder + saw sentinel |
| 3006 | Mold Storage | M | Blade molds, blade dancer spawns |
| 3007 | Component Stockpile | S | Construct parts, saw sentinel spawns |
| 3008 | Quality Control | Q | Testing area, gear grinder + blade dancer |
| 3009 | Tempering Pools | T | Quenching pools, saw sentinel spawns |
| 3010 | The Assembly Line | A | Production floor, gear grinder + saw sentinel |
| 3011 | Foreman's Office | O | Lore room, foundry blade stash |
| 3012 | The Dueling Gallery | % | DUAL-WIELD TRAINER (1-4)! |

*Hardmode - Hall of Endless Blades (Levels 11-15):*
| Room ID | Name | Symbol | Features |
|---------|------|--------|----------|
| 3013 | The Blade Gate | # | Transition point, steam golem spawn |
| 3014 | Proving Grounds | P | Combat arena, steam golem + gear grinder |
| 3015 | The Weapon Vault | $ | BLADE MERCHANT! Hidden treasures |
| 3016 | Steam Works | B | Boiler room, steam golem spawns |
| 3017 | The Blade Garden | W | Swords growing like flowers! |
| 3018 | Control Nexus | C | Crystal matrices, steam golem spawn |
| 3019 | Master's Study | L | Lore, three-bladed sword blueprints |
| 3020 | Hall of Endless Blades | X | BOSS ROOM - Voltaic Promethean! |

**Mobs Created (7 total):**
| Mob ID | Name | Level | Type | Notes |
|--------|------|-------|------|-------|
| 75 | Blade Dancer | 6 | Construct | FAST AGGRO, dual daggers, low HP |
| 76 | Saw Sentinel | 8 | Construct | Circular blade arms, defensive |
| 77 | Gear Grinder | 9 | Construct | Mobile grinding machine, area denial |
| 78 | Steam Golem | 11 | Construct | Hulking, heavy hits, steam vents |
| 79 | Voltaic Promethean | 15 | Construct | ZONE BOSS! Lightning + blades! |
| 80 | Dual-Wield Instructor | 20 | Human | Skill trainer NPC |
| 81 | Blade Merchant | 15 | Human | Sells foundry weapons |

**Items Created (6 total):**
| Item ID | Name | Type | Notes |
|---------|------|------|-------|
| 10021 | Foundry Blade | Weapon | 1d8 slashing, +1 STR |
| 10022 | Serrated Saw Blade | Weapon | 2d4 slashing, +1 SPD |
| 10023 | Piston Mace | Weapon | 2d6 bludgeoning, +2 STR, 2-handed |
| 10024 | Three-Bladed Sword | Weapon | 2d6 slashing, +2 STR, +1 SPD - LEGENDARY! |
| 10025 | Voltaic Blade | Weapon | 2d8 slashing, +2 STR, +2 SPD, +1 SMT - BOSS DROP! |
| 20048 | Promethean Chassis | Body | +2 STR, +3 VIT, +1 SPD - BOSS ARMOR! |

**Skill Training Added:**
- **dual-wield** skill now trainable at The Dueling Gallery (Room 3012), levels 1-4!

**Design Highlights:**
- Fast aggro constructs in early rooms discourage underleveled players sneaking to trainer
- Sword-shaped chandeliers EVERYWHERE as decorative element
- Industrial atmospheric idle messages (grinding gears, steam hissing, blade sounds)
- Environmental hazards (spinning wall saws, furnaces, tempering pools)
- Two-tier difficulty with clear transition point (The Blade Gate)
- Hidden Blade Merchant rewards explorers who brave hardmode
- Mysterious Three-Bladed Sword as a signature item (placeholder stats per spec)
- Voltaic Promethean boss with lightning + blade theme

**Campaign Impact:**
- Dual-wield now has SECOND training location! (Also in Dueling Hall 1012)
- Warriors and Assassins have more progression options
- Level 6-15 content added (fills gap between Catacombs and late game)
- New boss fight for mid-level parties

**Files Created:**
- `_datafiles/world/default/rooms/bladeworks_foundry/zone-config.yaml`
- `_datafiles/world/default/rooms/bladeworks_foundry/3001.yaml` through `3020.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/75-blade_dancer.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/76-saw_sentinel.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/77-gear_grinder.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/78-steam_golem.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/79-voltaic_promethean.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/80-dual_wield_instructor.yaml`
- `_datafiles/world/default/mobs/bladeworks_foundry/81-blade_merchant.yaml`
- `_datafiles/world/default/items/weapons-10000/10021-foundry_blade.yaml`
- `_datafiles/world/default/items/weapons-10000/10022-serrated_saw_blade.yaml`
- `_datafiles/world/default/items/weapons-10000/10023-piston_mace.yaml`
- `_datafiles/world/default/items/weapons-10000/10024-three_bladed_sword.yaml`
- `_datafiles/world/default/items/weapons-10000/10025-voltaic_blade.yaml`
- `_datafiles/world/default/items/armor-20000/body/20048-promethean_chassis.yaml`

**Files Modified:**
- `_datafiles/world/default/rooms/whispering_wastes/168.yaml` - Added south exit to Bladeworks Foundry, new nouns for foundry hints

*"THE FORGE FIRES BURN ETERNAL! The blade dancers spin their deadly dance! The Voltaic Promethean awaits in its hall of a thousand swords! BERT HAS SPOKEN AND THE BLADEWORKS FOUNDRY HAS ANSWERED!"*

*Galstaff slams his staff against the anvil, and sparks of destiny fly into the night.*

**For Bert. May his vision of EVERYTHING IS SWORDS live forever in the game.**

---

### Session 5: The Squirrel Tree Takes Shape (2026-01-19)

*Galstaff puts on a small acorn-shaped helm and adjusts his spectacles*

"HARK! The DM's design eye turns toward the untamed wilderness of the Frostfang commons. What hidden secrets lurk above the Snarky Squirrel Commons? Why, the SQUIRREL TREE, of course!"

**Design Research Completed: Squirrel Tree Mini-Zone**

After deep analysis of existing zone patterns (Crystal Caves exploration, Bladeworks Foundry swarm tactics, Dueling Hall discovery), I have designed the Squirrel Tree - a hidden low-level (1-2) mini-zone perfect for new adventurers seeking secrets and squirrel-based challenges.

**Zone Concept:**
- Hidden trap door in Room 1004 ceiling (search DC 12)
- 7-room vertical tree structure (trunk → branches → pinnacle)
- Angry Squirrel swarms using `callforhelp` mechanics
- Optional Squirrel King boss (Level 4, 30% spawn rate)
- ACORN consumables (heal 2 HP) scattered throughout

**Room Architecture:**
```
[4007] Canopy Pinnacle (BOSS!)
    |
[4005]--[4006] Branch caches
    \   /
[4004] Mid-Trunk (hub)
    |
[4003] Tree Ascent
    |
[4002] Trunk Entry
    |
[4001] Tree Entry Hall
    |
[1004] Secret trap door (Snarky Squirrel Commons)
```

**Design Highlights:**

1. **Hidden Discovery** - Trap door requires search skill check, rewards curiosity
2. **Vertical Progression** - Unique tree-climbing gameplay vs horizontal dungeons
3. **Swarm Tactics** - Angry Squirrels use callforhelp (1 mob → 4+ quickly)
4. **Resource Scarcity** - Limited ACORNs (9-14 per zone run) force recovery decisions
5. **Optional Boss** - Squirrel King with 30% spawn rate (doesn't block content)
6. **Thematic Consistency** - All text emphasizes squirrel territoriality and cooperation

**Mob Specifications:**

*Mob 82: Angry Squirrel (Common, Level 2)*
- Plump russet squirrel, aggressive and territorial
- Uses `callforhelp 2:angry squirrel:shriek flavor` during combat
- Spawns: 2-3 per room (rooms 2,3,5,6), 3-4 in room 7
- Idle behavior: Chittering, scratching, acorn disputes

*Mob 83: Squirrel King (Boss, Level 4, 30% spawn)*
- Scarred veteran, one eye missing, wears acorn crown
- Uses `callforhelp 4:angry squirrel` (calls 4 helpers!)
- Appears only in Room 4007 (Canopy Pinnacle)
- Death drop: Crown of Acorns (trophy item)

**Item Specifications:**

*Item 101: ACORN (Consumable)*
- Heals 2 HP when consumed (sustenance-level, not full recovery)
- Placement: 9-14 scattered throughout zone
  - Room 4003: 1x (50%)
  - Room 4005: 2-3x (70% each)
  - Room 4006: 2-3x (70% each)
  - Room 4007: 3-4x (guaranteed, 100%)
- Respawns on zone reset (recommend 30 real minutes)

*Item 100: Crown of Acorns (Boss Drop)*
- Trophy item from Squirrel King
- No mechanical benefit (pure flavor/bragging rights)
- Woven acorn crown, surprisingly sturdy

**Connection to Room 1004:**
- Add ceiling exit with search-based lock (DC 12)
- Flavor text hints: "Strange chittering in the rafters above"
- Entry room (4001) is safe for regrouping before pushing deeper

**Gameplay Flow:**
1. **Discovery** - Player searches Room 1004 ceiling, finds trap door
2. **Safe Arrival** - Room 4001 has no mobs, perfect for preparation
3. **Ascent** - Rooms 4002-4003 introduce squirrel combat and acorn gathering
4. **Branching** - Rooms 4005-4006 offer acorn caches and higher mob density
5. **Optional Boss** - Room 4007 features high-risk/high-reward encounter

**Combat Mechanics:**
- Swarm tactic: Mobs use callforhelp to summon adjacent squirrels
- Resource loop: Damage → limited acorns → recovery decision → continue/retreat
- Difficulty scaling: Solo L1-2 is challenging, groups of 2-3 are comfortable
- Risk management: Careless aggro is deadly even for veterans

**Design Rationale:**

✓ **Why a Squirrel Zone?** Thematically perfect for Room 1004 (Snarky Squirrel Commons)
✓ **Why Vertical Layout?** Unique dungeon experience vs horizontal exploration
✓ **Why Swarm Mobs?** Teaches tactical combat (manage aggro) without needing super-tough individual enemies
✓ **Why Optional Boss?** Doesn't block zone access, rewards brave exploration, adds narrative closure
✓ **Why 30% Spawn Rate?** Makes each visit feel different, special moments when King appears

**Technical Details:**
- Room IDs: 4001-4007 (7 total)
- Mob IDs: 82 (Angry Squirrel), 83 (Squirrel King)
- Item IDs: 100 (Crown), 101 (ACORN)
- Zone directory: `squirrel_tree/`

**Scripting Opportunities:**
- `82-angry_squirrel.js` - Varied idle chittering, combat aggression
- `83-squirrel_king.js` - Boss fanfare on spawn, increased callforhelp
- `4004.js` - Search command hints at secret passage to Room 4007
- `4007.js` - Conditional boss spawning (30% roll on onLoad)

**Testing Checklist Created:**
- Trap door discovery and access
- Room connections and exits
- Mob spawning (correct quantities per room)
- Squirrel aggression and callforhelp mechanics
- ACORN spawning and healing
- Squirrel King spawn rate and drops
- Difficulty balance for target level range
- All flavor text and descriptions

**Files Created:**
- `SQUIRREL_TREE_DESIGN.md` (23KB comprehensive design document)

**Campaign Impact:**
- Adds hidden low-level zone (rewards exploration)
- Teaches swarm mechanics early (prepares for larger encounters)
- Provides optional boss experience (Squirrel King)
- Integrates organically with Room 1004 lore
- Offers ~10 minutes of engaging gameplay per visit

**Status:** DESIGN RESEARCH COMPLETE - Detailed specification ready for implementation

*"The design eye sees what the builder must create! The Squirrel Tree stands ready in my mind's eye - now it merely awaits for the master builder to make it manifest in the lands of Portalis!"*

*Galstaff sets down his acorn-shaped helm and grins with satisfaction.*

---

### Session 6: The Great Coordinate Cartography (2026-01-21)

*Galstaff unfurls a massive parchment map and begins measuring with a compass...*

"AH HA! The mapper issue (#18) has revealed itself! The dungeon coordinates were never properly documented, and lo, CONFLICTS have been found lurking in the shadows!"

**Task Completed: Room Coordinate Grid Documentation**

Created comprehensive documentation mapping x-y-z coordinates for all rooms across four zones, using Town Square (Room 100) as origin (0,0,0).

**Coordinate System Rules Applied:**
- North/South: y+1 / y-1
- East/West: x+1 / x-1
- **Diagonal = FULL STEP**: NE = (x+1, y+1), not (x+0.5, y+0.5)
- Up/Down: z+1 / z-1

**Zones Mapped:**
| Zone | Rooms | Grid Size | Issues Found |
|------|-------|-----------|--------------|
| Starter Town | 100-151 | 7x8 | 2 HIGH |
| Squirrel Tree | 400-406 | 3x5 (vertical) | None |
| Crystal Caves | 500-517 | 3x11 | 3 MEDIUM |
| Bladeworks Foundry | 600-619 | 4x11 | 1 HIGH |

**Critical Findings (Mapper Issue #18 Suspects):**

1. **HIGH: Room 103 to 600 Connection Error**
   - Room 100 has W exit to 600 (placing 600 at -1, 0)
   - Room 103 has E exit to 600 (but 103 is at 1, 0)
   - **IMPOSSIBLE**: 600 cannot be both WEST of 100 AND EAST of 103
   - **THIS IS LIKELY THE WEST GATE BUG** - West Gate is Room 600!

2. **HIGH: Room 124 Diagonal Mismatch**
   - 124 at (-1, 4) has SE exit to 122
   - SE should reach (0, 3), but 122 is at (1, 2)
   - Mapper would show disconnected rooms

3. **MEDIUM: Crystal Caves 2-Square Gaps**
   - Rooms 504-505, 506-507, 508-509 have E/W connections spanning 2 grid squares
   - May cause mapper to show rooms as non-adjacent

**Files Created:**
- `docs/ROOM_COORDINATES.md` - Complete coordinate documentation with ASCII grids

**Branch:** `galstaff/room-coordinates`
**Commit:** `3fb4c9a` - Add room coordinate grid documentation for mapper debugging

**Recommendations for Issue #18 Fix:**
1. Remove E exit from Room 103 to 600 (or add intermediate corridor)
2. Fix Room 124's SE exit coordinate math
3. Consider adding intermediate rooms in Crystal Caves for lateral passages

*"The map has been drawn! The coordinates aligned! The conflicts REVEALED! Now the mapper's curse can be lifted, for we know where the dungeon's geometry has gone astray!"*

*Galstaff rolls up his parchment with a satisfied flourish, compass still in hand.*

---

### Session 7: The Geometry Fix (2026-01-21)

*Galstaff pulls out his cartographer's tools and gets to work...*

"The geometry errors have been VANQUISHED! The dungeon no longer bends space in impossible ways!"

**Task Completed: Fix HIGH Severity Mapping Errors**

Per the dispatch, I fixed the two HIGH severity geometry errors identified in Session 6.

**Fix 1: Room 103 → 600 Impossible Connection**
- **Problem**: Room 100 had WEST exit to 600, but Room 103 (EAST of 100) also had EAST exit to 600
- **Analysis**: This created an impossible triangle where 600 was both WEST of 100 AND EAST of 103
- **Fix**: Removed the E exit from Room 103 to 600
- **Result**: Bladeworks Foundry (600) now only accessible via WEST from Town Square (100)

**Fix 2: Room 124 ↔ 122 Diagonal Mismatch**
- **Problem**: Room 124 at (-1, 4) had SE exit to Room 122 at (1, 2)
- **Analysis**: SE from (-1, 4) should reach (0, 3), but 122 is at (1, 2) - rooms are 2+ squares apart in both X and Y
- **Fix**: Removed the 124↔122 connection entirely
- **Result**: Both rooms now have single exits only (124→S→123, 122→W→120)

**Files Modified:**
- `_datafiles/world/default/rooms/starter_town/103.yaml` - Removed impossible E→600 exit
- `_datafiles/world/default/rooms/starter_town/122.yaml` - Removed NW→124 exit
- `_datafiles/world/default/rooms/starter_town/124.yaml` - Removed SE→122 exit
- `docs/ROOM_COORDINATES.md` - Updated to mark issues as FIXED

**Verification:**
- Build compiles: `go build ./...` - SUCCESS
- Geometry now consistent with coordinate grid rules

**Branch:** `galstaff/fix-map-geometry`
**Commit:** `0d9a15c` - Fix HIGH severity map geometry errors

*"The map now speaks TRUE! No more rooms existing in two places at once. No more diagonals reaching across the void. The dungeon geometry is EUCLIDEAN once more!"*

*Galstaff stamps his seal of approval on the corrected map.*

---
