# Portalis Balance Roadmap

A comprehensive guide for designing balanced content. Built from analysis of GoMUD's existing content.

---

## Part 1: How The System Works

### XP Formula
```
XPTL(level) = 1000 + (level² × 0.75 × 1000)
Mob XP Reward = XPTL(mobLevel - 1) / 90
```

| Mob Level | XP Reward | XP to reach that level |
|-----------|-----------|------------------------|
| 1 | 11 | 1,750 |
| 5 | 144 | 19,750 |
| 10 | 686 | 76,000 |
| 15 | 1,644 | 169,750 |
| 20 | 3,019 | 301,000 |
| 25 | 4,811 | 469,750 |
| 30 | 7,019 | 676,000 |

**Key Insight:** Mob level primarily determines XP reward (exponential scaling).

### HP Formula
```
HP = 5 + Level + (Vitality × 4)
```

| Race | Base VIT | L10 HP | L20 HP | L30 HP |
|------|----------|--------|--------|--------|
| Human | 1 | 19 | 29 | 39 |
| Undead | 0 | 15 | 25 | 35 |
| Robot | 4 | 31 | 41 | 51 |
| Golem | 3 | 27 | 37 | 47 |
| Fungus | 4 | 31 | 41 | 51 |

**Key Insight:** Mob level adds 1 HP per level. VIT adds 4 HP per point.

### Damage System
```
If weapon equipped: Use weapon dice (completely replaces race damage)
If unarmed: Use race damage dice
```

**CRITICAL:** The `equipment:` block determines what a mob WIELDS.
The `items:` block is just INVENTORY (drops on death).

---

## Part 2: Existing Game Baseline

### Weapon Damage Tiers

| Tier | Dice | Range | Avg | Examples |
|------|------|-------|-----|----------|
| 0 | 1d2 | 1-2 | 1.5 | sharp stick, crowbar |
| 1 | 1d3-1d4 | 1-4 | 2.5 | dagger, wooden claw |
| 2 | 1d6 | 1-6 | 3.5 | broadsword |
| 3 | 1d6+1, 1d8 | 2-8 | 4.5 | shortsword, foundry blade, cudgel |
| 4 | 2d4 | 2-8 | 5 | obsidian dagger, serrated saw blade |
| 5 | 2d6 | 2-12 | 7 | piston mace |
| 6 | 2d8 | 2-16 | 9 | voltaic blade |
| 7 | 2d8+2, 2d10+1 | 4-18+ | 11 | ancient scepter, glowing battleaxe |

**Dual Wielding:** Two 1-handed weapons = 2 attacks. Each uses its own dice.

### Existing Mob Analysis

| Zone | Mob | Level | Weapon | Damage | Gold | Gold/Lvl |
|------|-----|-------|--------|--------|------|----------|
| **Frostfang** |
| | rat | 1 | unarmed (rodent) | 1-2 | 0 | 0 |
| | big rat | 5 | unarmed (rodent) | 1-2 | 0 | 0 |
| **Slums** |
| | ruffian | 20 | dual claws 1d4×2 | 2-8 | 10 | 0.5 |
| | dangerous ruffian | 25 | obsidian 2d4 | 2-8 | 20 | 0.8 |
| **Catacombs** |
| | skeleton | 15 | broadsword 1d6 | 1-6 | 3 | 0.2 |
| | dark acolyte | 20 | cudgel 1d8 | 1-8 | 7 | 0.35 |
| | dark acolyte | 28 | dual obsidian 2d4×2 | 4-16 | 7 | 0.25 |
| | lich (boss) | 30 | scepter 2d8+2 | 4-18 | 50 | 1.67 |

### Design Pattern Discovered

The existing game uses **HIGH LEVEL + LOW DAMAGE**:
- Mobs are L15-30 (lots of XP, tanky HP)
- Weapons are Tier 1-3 (1d4 to 1d8)
- Result: Long survivable fights, fast leveling

**NOT** low level + high damage (what we accidentally did).

---

## Part 3: Challenge Level Framework

### The Challenge Level Concept

**Challenge Level (CL)** = the effective difficulty of content, combining:
- Monster Level (XP, HP, stats)
- Damage Output (weapon tier)
- Special Abilities (call for help, combat commands)
- Encounter Density (how many mobs)

A **Player Level (PL)** can comfortably fight CL = PL ± 3.

### Challenge Level Formula

```
Base CL = Monster Level

Damage Modifier:
  Tier 0-1 weapons: CL - 2
  Tier 2-3 weapons: CL + 0
  Tier 4-5 weapons: CL + 3
  Tier 6-7 weapons: CL + 5
  Dual wield: +2 to modifier

Special Modifiers:
  Calls for help: CL + 2
  Multiple attacks (race): CL + 3
  Boss (zone-boss group): CL + 5
```

### Example Calculations

**Ruffian:**
- Monster Level 20
- Dual claws (Tier 1): +0 damage mod, +2 dual wield
- CL = 20 + 0 + 2 = **CL 22**
- Appropriate for PL 19-25

**Skeleton:**
- Monster Level 15
- Broadsword (Tier 2): +0
- CL = 15 + 0 = **CL 15**
- Appropriate for PL 12-18

**Dark Acolyte (L28):**
- Monster Level 28
- Dual obsidian (Tier 4): +3, +2 dual wield
- CL = 28 + 3 + 2 = **CL 33**
- Appropriate for PL 30-36

**Lich:**
- Monster Level 30
- Scepter (Tier 7): +5
- Boss: +5
- CL = 30 + 5 + 5 = **CL 40**
- Appropriate for PL 37+ (endgame boss)

---

## Part 4: Zone Design Brackets

### Bracket 1: Tutorial (PL 1-3, CL 1-5)
**Purpose:** Learn the game, can't really die

| Role | Monster Level | Weapon Tier | Damage | Gold |
|------|---------------|-------------|--------|------|
| Fodder | 1-2 | 0 (unarmed) | 1-2 | 0 |
| Trainer | 3-5 | 0-1 | 1-4 | 0-1 |

**XP per kill:** 11-44
**Kills to level:** 40-80

### Bracket 2: Starter Zone (PL 3-8, CL 5-12)
**Purpose:** First real combat, introduce mechanics

| Role | Monster Level | Weapon Tier | Damage | Gold |
|------|---------------|-------------|--------|------|
| Fodder | 5-8 | 0-1 | 1-4 | 1-3 |
| Standard | 8-12 | 1-2 | 1-6 | 3-5 |
| Elite | 10-15 | 2-3 | 2-8 | 5-10 |
| Mini-boss | 12-18 | 3 | 2-8 | 15-25 |

**XP per kill:** 144-1,200
**Kills to level:** 15-40

### Bracket 3: Adventure Zone (PL 8-15, CL 12-22)
**Purpose:** Core gameplay loop, meaningful choices

| Role | Monster Level | Weapon Tier | Damage | Gold |
|------|---------------|-------------|--------|------|
| Fodder | 12-15 | 1-2 | 1-6 | 3-5 |
| Standard | 15-20 | 2-3 | 2-8 | 5-10 |
| Elite | 18-25 | 3-4 | 2-8 | 8-15 |
| Zone Boss | 22-28 | 4-5 | 4-12 | 30-50 |

**XP per kill:** 1,200-4,000
**Kills to level:** 8-20

### Bracket 4: Expert Zone (PL 15-25, CL 22-35)
**Purpose:** Challenge content, group encouraged

| Role | Monster Level | Weapon Tier | Damage | Gold |
|------|---------------|-------------|--------|------|
| Fodder | 18-22 | 2-3 | 2-8 | 5-10 |
| Standard | 22-28 | 3-4 | 2-8 | 8-15 |
| Elite | 25-32 | 4-5 | 4-12 | 12-25 |
| Zone Boss | 30-38 | 5-6 | 4-16 | 50-100 |

**XP per kill:** 2,500-6,500
**Kills to level:** 6-12

### Bracket 5: Endgame (PL 25+, CL 35+)
**Purpose:** Ultimate challenges, rare rewards

| Role | Monster Level | Weapon Tier | Damage | Gold |
|------|---------------|-------------|--------|------|
| Standard | 28-35 | 4-5 | 4-12 | 15-30 |
| Elite | 32-40 | 5-6 | 4-16 | 25-50 |
| World Boss | 40-50 | 6-7 | 6-18+ | 100-500 |

---

## Part 5: Gold Economy

### Gold Per Level Ratio

Based on existing content analysis:

| Content Type | Gold/Level Ratio |
|--------------|------------------|
| Trash mobs | 0.1-0.3 |
| Standard mobs | 0.3-0.5 |
| Elite mobs | 0.5-0.8 |
| Mini-boss | 1.0-2.0 |
| Zone boss | 2.0-4.0 |

### Gear Affordability

**Target:** 30-60 minutes of farming appropriate content = one gear upgrade

| Bracket | Gear Value Range | Mobs to Afford (standard) |
|---------|------------------|---------------------------|
| 1 | 0-20g | N/A (drops/quests) |
| 2 | 20-100g | 30-50 |
| 3 | 100-500g | 30-50 |
| 4 | 500-2000g | 40-60 |
| 5 | 2000-10000g | 50-100 |

---

## Part 6: Our Content - Current State & Fixes

### Squirrel Tree (Target: Bracket 2)

| Mob | Current | Problem | Fix |
|-----|---------|---------|-----|
| angry squirrel | L5, unarmed rodent (1-2), 1g | Fine | Keep |
| squirrel king | L7, unarmed rodent (1-2), 25g | Gold too high, damage too low for boss | L10, give weapon (1d4), 10g |

### Crystal Caves (Target: Bracket 2)

| Mob | Current | Problem | Fix |
|-----|---------|---------|-----|
| crystal beetle | L4, unarmed insect (2-4), 2g | Fine | Keep |
| sporeling | L4, unarmed fungus (1-4), 1g | Fine | Keep |
| cavern sprite | L5, unarmed faerie (1-2), 5g | Fine | Keep |
| glowcap wanderer | L6, unarmed fungus (1-4), 3g | Fine | Keep |
| crystal guardian | L8, unarmed golem (2-10), 15g | Gold high | L8, 8g |
| crystal matriarch | L12, unarmed golem (2-10), 100g | Gold very high | L15, 40g |

### Bladeworks Foundry (Target: Bracket 2-3)

**CRITICAL BUG:** All robots have weapons in `items:` not `equipment:`. They fight UNARMED with 2×2d6 (4-24 damage) - way too strong!

| Mob | Current State | Actual Damage | Fix |
|-----|---------------|---------------|-----|
| blade dancer | L6, UNARMED robot | 4-24 | L12, equip foundry blade (1-8), 5g |
| saw sentinel | L8, UNARMED robot | 4-24 | L15, equip saw blade (2-8), 6g |
| gear grinder | L9, UNARMED robot | 4-24 | L18, stays unarmed (thematic), 8g |
| steam golem | L11, UNARMED robot | 4-24 | L22, equip piston mace (2-12), 10g |
| voltaic promethean | L15, UNARMED robot | 4-24 | L28, equip voltaic blade (2-16), 50g |

**Reasoning:**
- Robots unarmed = 4-24 damage = Tier 6-7 equivalent
- CL = Monster Level + 5 (high damage) + 3 (multi-attack)
- Need Monster Level 12-28 to match CL 20-36
- This gives appropriate XP for the danger

### Construction Robot (Frostfang)

| Current | Problem | Fix |
|---------|---------|-----|
| L18, cudgel (1-8), 35g | Gold high for location | L18, 15g |

---

## Part 7: Implementation Checklist

### Phase 1: Fix Equipment Bugs
- [ ] Blade dancer: Move foundry blade from `items` to `equipment.weapon`
- [ ] Saw sentinel: Move saw blade from `items` to `equipment.weapon`
- [ ] Steam golem: Move piston mace from `items` to `equipment.weapon`
- [ ] Voltaic promethean: Move voltaic blade from `items` to `equipment.weapon`

### Phase 2: Rebalance Levels
- [ ] Blade dancer: L6 → L12
- [ ] Saw sentinel: L8 → L15
- [ ] Gear grinder: L9 → L18
- [ ] Steam golem: L11 → L22
- [ ] Voltaic promethean: L15 → L28
- [ ] Squirrel king: L7 → L10, add weapon
- [ ] Crystal matriarch: L12 → L15

### Phase 3: Rebalance Gold
- [ ] Squirrel king: 25g → 10g
- [ ] Crystal guardian: 15g → 8g
- [ ] Crystal matriarch: 100g → 40g
- [ ] Construction robot: 35g → 15g
- [ ] Blade dancer: 5g (keep)
- [ ] Saw sentinel: 5g → 6g
- [ ] Gear grinder: 6g → 8g
- [ ] Steam golem: 8g → 10g
- [ ] Voltaic promethean: 50g (keep)

---

## Part 8: New Zone Template

When creating a new zone, use this template:

```
Zone Name: [TBD with family]
Target Bracket: [1-5]
Target Player Level: [X-Y]

Entry Mob:
  Monster Level: [bracket fodder level]
  Weapon Tier: [0-2]
  Gold: [bracket range low]

Standard Mob:
  Monster Level: [bracket standard level]
  Weapon Tier: [bracket tier]
  Gold: [bracket range mid]

Elite Mob:
  Monster Level: [bracket elite level]
  Weapon Tier: [bracket tier + 1]
  Gold: [bracket range high]
  Special: [call for help / combat commands]

Zone Boss:
  Monster Level: [bracket boss level]
  Weapon Tier: [bracket tier + 1-2]
  Gold: [bracket × 3-5]
  Special: [call for help, unique drops]
  Groups: [zone-boss]

Trainer:
  Skills: [list]
  Levels: [min-max per skill]

Shop:
  Items: [list with values in bracket range]
```
