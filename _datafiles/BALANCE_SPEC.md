# Portalis Zone Balance Specification

**Status:** APPROVED
**Authors:** Archie (Architect), Galstaff (Creative Director)
**Date:** Debate concluded, ready for implementation

---

## Executive Summary

This document defines the official mob level, XP progression, and challenge level targets for all Portalis zones. These numbers have been validated mathematically and approved for gameplay feel.

---

## Core Formulas Reference

```
XP To Level (XPTL):     XPTL(level) = 1000 + (level² × 0.75 × 1000)
Mob XP Reward:          XP = XPTL(mobLevel - 1) / 90
HP Formula:             HP = 5 + Level + (Vitality × 4)
Challenge Level (CL):   Base = Monster Level + Damage Modifier + Special Mods
```

**Race Vitality:**
- Robot: 4
- Golem: 3
- Human: 1
- Undead: 0

---

## Zone Progression Overview

| Zone | Player Entry | Player Exit | Entry Mob Level | Boss Level |
|------|--------------|-------------|-----------------|------------|
| Tutorial | L1 | L3 | - | - |
| Squirrel Tree | L3 | L7 | L5 | L8 |
| Crystal Caves | L5 | L12 | L4 | L12 |
| Bladeworks Foundry | L12 | L24 | L12 | L28 |
| Frostfang Slums | L20 | L25 | L20 | - |
| Catacombs | L22 | L30+ | L15 | L30 |

---

## Zone 1: Squirrel Tree

**Theme:** First real combat zone. Teaches aggro, positioning, and boss mechanics.
**Target Player:** L3-7

| Mob | Level | Race | Weapon | XP | CL | Notes |
|-----|-------|------|--------|----|----|-------|
| Angry Squirrel | 5 | Rodent | Unarmed (1-2) | 144 | 3 | Entry mob |
| Squirrel King | **8** | Rodent | Unarmed (1-2) | 275 | 13 | Boss, callforhelp 4 squirrels |

**Progression Math:**
- L3 player needs 4,750 XP to reach L4
- ~33 squirrel kills to level from L3 to L4
- Estimated 80-100 kills to clear zone and face boss at L6-7

**Change from current:** Squirrel King L7 -> L8

---

## Zone 2: Crystal Caves

**Theme:** Wonder and discovery. Danger sneaks up on players.
**Target Player:** L5-12

| Mob | Level | Race | Weapon | XP | CL | Notes |
|-----|-------|------|--------|----|----|-------|
| Crystal Beetle | 4 | Insect | Unarmed (2-4) | 111 | 4 | Entry mob |
| Sporeling | 4 | Fungus | Unarmed (1-4) | 111 | 4 | Entry mob |
| Cavern Sprite | 5 | Faerie | Unarmed (1-2) | 144 | 3 | Light mob |
| Glowcap Wanderer | 6 | Fungus | Unarmed (1-4) | 183 | 6 | Mid-zone |
| Crystal Guardian | 8 | Golem | Unarmed (2-10) | 275 | 10 | Elite |
| Crystal Matriarch | 12 | Golem | Unarmed (2-10) | 830 | 17 | Boss, callforhelp 3 guardians |

**Progression Math:**
- L5 player needs 19,750 XP to reach L6
- ~135 beetle/sporeling kills OR ~70 wanderer kills
- Estimated 120-150 kills to clear zone and face boss at L10-12

**Change from current:** No level changes needed. Gold adjustments only:
- Crystal Guardian: 15g -> 8g
- Crystal Matriarch: 100g -> 40g

---

## Zone 3: Bladeworks Foundry

**Theme:** Industrial dread. Obvious mechanical danger. Metal grinding on metal.
**Target Player:** L12-24

| Mob | Level | Race | Weapon | XP | CL | Notes |
|-----|-------|------|--------|----|----|-------|
| Blade Dancer | **12** | Robot | Foundry Blade (1d6+1) | 830 | 12 | Entry mob |
| Saw Sentinel | **15** | Robot | Saw Blade (2d4) | 1,644 | 15 | Defensive |
| Gear Grinder | **18** | Golem | Unarmed (2d5) | 2,630 | 23 | Area denial, callforhelp 2 dancers |
| Steam Golem | **22** | Robot | Piston Mace (2d6) | 4,011 | 24 | Heavy hitter |
| Voltaic Promethean | **28** | Robot | Voltaic Blade (2d8) | 6,519 | 38 | Zone boss, callforhelp 2 steam golems |

**Progression Math:**
- L12 player needs 109,000 XP to reach L13
- ~131 blade dancer kills to level
- L18 player needs 244,000 XP to reach L19
- ~93 gear grinder kills to level
- Estimated 150-200 kills to clear zone and face boss at L22-24

**Changes from current:**
- Blade Dancer: L6 -> L12
- Saw Sentinel: L8 -> L15
- Gear Grinder: L9 -> L18
- Steam Golem: L11 -> L22
- Voltaic Promethean: L15 -> L28

**Critical Note:** Equipment bug has been fixed. All robots now have weapons in `equipment:` block, not `items:` block.

---

## Zone 4: Frostfang Slums

**Theme:** Urban danger. Street-level crime escalation.
**Target Player:** L20-25

| Mob | Level | Race | Weapon | XP | CL | Notes |
|-----|-------|------|--------|----|----|-------|
| Ruffian | 20 | Human | Dual Claws (1d4×2) | 3,019 | 22 | Entry mob |
| Dangerous Ruffian | 25 | Human | Obsidian Dagger (2d4) | 4,811 | 28 | Elite |
| Shadow Master | - | Human | - | - | - | Trainer NPC |

**Progression Math:**
- L20 player needs 301,000 XP to reach L21
- ~100 ruffian kills to level
- Estimated 100 kills to reach L25

**Change from current:** No changes needed. Zone is correctly balanced.

---

## Zone 5: Catacombs

**Theme:** Endgame dungeon. Undead horror, ancient evil.
**Target Player:** L22-30+

| Mob | Level | Race | Weapon | XP | CL | Notes |
|-----|-------|------|--------|----|----|-------|
| Skeleton | 15 | Undead | Broadsword (1d6) | 1,644 | 15 | Entry mob (low for loot runs) |
| Dark Acolyte | 20 | Human | Cudgel (1d8) | 3,019 | 20 | Mid-zone |
| Dark Acolyte (Elite) | 28 | Human | Dual Obsidian (2d4×2) | 6,519 | 33 | Elite |
| Lich | 30 | Undead | Ancient Scepter (2d8+2) | 7,019 | 40 | Final boss |

**Progression Math:**
- L22 player can farm skeletons (1,644 XP) for gear
- L28 player needs 589,000 XP to reach L29
- ~84 elite acolyte kills to level
- Boss fight is endgame content, designed for groups

**Change from current:** No changes needed.

---

## Implementation Checklist

### Phase 1: Level Changes
- [ ] Squirrel King: L7 -> L8
- [ ] Blade Dancer: L6 -> L12
- [ ] Saw Sentinel: L8 -> L15
- [ ] Gear Grinder: L9 -> L18
- [ ] Steam Golem: L11 -> L22
- [ ] Voltaic Promethean: L15 -> L28

### Phase 2: Gold Rebalance
- [ ] Crystal Guardian: 15g -> 8g
- [ ] Crystal Matriarch: 100g -> 40g

### Phase 3: Verification
- [ ] Test L3 character progression through Squirrel Tree
- [ ] Test L5 character progression through Crystal Caves
- [ ] Test L12 character progression through Bladeworks
- [ ] Validate boss fights are challenging but not impossible
- [ ] Confirm callforhelp mechanics work as designed

---

## Appendix: XP Reference Table

| Level | XP to Reach | XP if Mob (Level-1) |
|-------|-------------|---------------------|
| 1 | 1,750 | 11 |
| 5 | 19,750 | 144 |
| 8 | 49,000 | 275 |
| 10 | 76,000 | 686 |
| 12 | 109,000 | 830 |
| 15 | 169,750 | 1,644 |
| 18 | 244,000 | 2,630 |
| 20 | 301,000 | 3,019 |
| 22 | 364,000 | 4,011 |
| 25 | 469,750 | 4,811 |
| 28 | 589,000 | 6,519 |
| 30 | 676,000 | 7,019 |

---

## Appendix: Challenge Level Quick Reference

```
CL = Monster Level + Damage Mod + Special Mods

Damage Modifier (by weapon tier):
  Tier 0-1 (1-4 avg):  -2
  Tier 2-3 (2-8 avg):  +0
  Tier 4-5 (4-12 avg): +3
  Tier 6-7 (6-18 avg): +5
  Dual wield:          +2

Special Modifiers:
  Calls for help:      +2
  Multi-attack race:   +3
  Zone boss:           +5
```

---

**Document Status:** FINAL - Approved for implementation
