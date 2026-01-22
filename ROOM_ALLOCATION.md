# Room & Entity ID Allocation Map

*The master bounding box registry. Zones web out from the central town.*

---

## World Structure

```
                              ┌─────────────────────┐
                              │   ADVANCED TOWN     │
                              │     (Level 15+)     │
                              │      900-999        │
                              └──────────┬──────────┘
                                         │
                              ┌──────────┴──────────┐
                              │   THE LONG ROAD     │
                              │   (Level 10-15)     │
                              │      800-899        │
                              │                     │
                              │  [GUARDED GATE]     │
                              │  Requires: L10 quest│
                              └──────────┬──────────┘
                                         │
         ┌───────────────────────────────┼───────────────────────────────┐
         │                               │                               │
    ┌────┴────┐                   ┌──────┴──────┐                   ┌────┴────┐
    │SQUIRREL │                   │   STARTER   │                   │BLADEWORKS│
    │  TREE   │◄──── (hidden) ────│    TOWN     │──── (distant) ───►│ FOUNDRY │
    │ 400-499 │                   │   100-199   │                   │ 600-699 │
    │ Lvl 1-2 │                   │   Lvl 1-3   │                   │ Lvl 6-15│
    └─────────┘                   │             │                   └─────────┘
                                  └──────┬──────┘
                                   /     │     \
                          ┌───────┘      │      └───────┐
                          │              │              │
                   ┌──────┴──────┐ ┌─────┴─────┐ ┌──────┴──────┐
                   │ MINNEAPOLIS │ │  CRYSTAL  │ │  BELLEVUE   │
                   │   200-299   │ │   CAVES   │ │   300-399   │
                   │  (40 homes) │ │  500-599  │ │  (40 homes) │
                   └─────────────┘ │  Lvl 3-8  │ └─────────────┘
                                   └───────────┘
```

---

## ID Allocation (100-room blocks)

| Block | Zone | Purpose | Notes |
|-------|------|---------|-------|
| **0-99** | Tutorial / System | Tutorial, shadow realm, death | From empty template |
| **100-199** | Starter Town | Core hub, trainers, shops, plaza | ~50 rooms active |
| **200-299** | Minneapolis | Residential neighborhood | 40 houses (200-239) |
| **300-399** | Bellevue | Residential neighborhood | 40 houses (300-339) |
| **400-499** | Squirrel Tree | Hidden zone off town | 7 rooms (migrate) |
| **500-599** | Crystal Caves | Main adventure zone | 18 rooms (migrate) |
| **600-699** | Bladeworks Foundry | Industrial zone | 20 rooms (migrate) |
| **700-799** | *[Future Zone]* | Expansion slot | |
| **800-899** | The Long Road | Travel zone with gate | Quest-locked |
| **900-999** | Advanced Town | Second hub, L15+ | Endgame content |

---

## Residential Neighborhoods

### Minneapolis (Block 200-299)

| Room ID | House # | Resident Name |
|---------|---------|---------------|
| 200 | 1 | _________________ |
| 201 | 2 | _________________ |
| 202 | 3 | _________________ |
| 203 | 4 | _________________ |
| 204 | 5 | _________________ |
| 205 | 6 | _________________ |
| 206 | 7 | _________________ |
| 207 | 8 | _________________ |
| 208 | 9 | _________________ |
| 209 | 10 | _________________ |
| 210 | 11 | _________________ |
| 211 | 12 | _________________ |
| 212 | 13 | _________________ |
| 213 | 14 | _________________ |
| 214 | 15 | _________________ |
| 215 | 16 | _________________ |
| 216 | 17 | _________________ |
| 217 | 18 | _________________ |
| 218 | 19 | _________________ |
| 219 | 20 | _________________ |
| 220 | 21 | _________________ |
| 221 | 22 | _________________ |
| 222 | 23 | _________________ |
| 223 | 24 | _________________ |
| 224 | 25 | _________________ |
| 225 | 26 | _________________ |
| 226 | 27 | _________________ |
| 227 | 28 | _________________ |
| 228 | 29 | _________________ |
| 229 | 30 | _________________ |
| 230 | 31 | _________________ |
| 231 | 32 | _________________ |
| 232 | 33 | _________________ |
| 233 | 34 | _________________ |
| 234 | 35 | _________________ |
| 235 | 36 | _________________ |
| 236 | 37 | _________________ |
| 237 | 38 | _________________ |
| 238 | 39 | _________________ |
| 239 | 40 | _________________ |
| 240-249 | - | Streets & common areas |

### Bellevue (Block 300-399)

| Room ID | House # | Resident Name |
|---------|---------|---------------|
| 300 | 1 | _________________ |
| 301 | 2 | _________________ |
| 302 | 3 | _________________ |
| 303 | 4 | _________________ |
| 304 | 5 | _________________ |
| 305 | 6 | _________________ |
| 306 | 7 | _________________ |
| 307 | 8 | _________________ |
| 308 | 9 | _________________ |
| 309 | 10 | _________________ |
| 310 | 11 | _________________ |
| 311 | 12 | _________________ |
| 312 | 13 | _________________ |
| 313 | 14 | _________________ |
| 314 | 15 | _________________ |
| 315 | 16 | _________________ |
| 316 | 17 | _________________ |
| 317 | 18 | _________________ |
| 318 | 19 | _________________ |
| 319 | 20 | _________________ |
| 320 | 21 | _________________ |
| 321 | 22 | _________________ |
| 322 | 23 | _________________ |
| 323 | 24 | _________________ |
| 324 | 25 | _________________ |
| 325 | 26 | _________________ |
| 326 | 27 | _________________ |
| 327 | 28 | _________________ |
| 328 | 29 | _________________ |
| 329 | 30 | _________________ |
| 330 | 31 | _________________ |
| 331 | 32 | _________________ |
| 332 | 33 | _________________ |
| 333 | 34 | _________________ |
| 334 | 35 | _________________ |
| 335 | 36 | _________________ |
| 336 | 37 | _________________ |
| 337 | 38 | _________________ |
| 338 | 39 | _________________ |
| 339 | 40 | _________________ |
| 340-349 | - | Streets & common areas |

---

## Mob ID Allocation (10-mob blocks)

| Block | Zone | Notes |
|-------|------|-------|
| **1-9** | Starter Town | Rats, guards, shopkeepers |
| **10-19** | Tutorial / Common | Shared enemies |
| **20-29** | Squirrel Tree | Squirrels, Squirrel King |
| **30-39** | Crystal Caves | Beetles, fungi, Matriarch |
| **40-49** | Bladeworks Foundry | Constructs, Promethean |
| **50-59** | The Long Road | Road encounters |
| **60-69** | Advanced Town | Guards, trainers |
| **70+** | Future zones | Expansion |

---

## Item ID Allocation

Items use category prefixes (no change needed):

| Range | Category |
|-------|----------|
| 1-999 | Other / Misc / Quest |
| 10000+ | Weapons |
| 20000+ | Armor |
| 30000+ | Consumables |

---

## Progression Flow

```
LEVEL 1-2:   [Tutorial] → [Starter Town] → (optional) [Squirrel Tree]
                              │
                         [Minneapolis]  [Bellevue]
                          (housing)      (housing)
                              │
LEVEL 3-5:                    └──→ [Crystal Caves]
                                        │
LEVEL 6-10:                             └──→ [Bladeworks Foundry]
                                                    │
LEVEL 10:              ═══════════════════════════════════════════
                       ║  COMPLETE TOWN QUEST TO PASS GATE  ║
                       ═══════════════════════════════════════════
                                        │
LEVEL 10-15:                    [The Long Road]
                                        │
LEVEL 15+:                      [Advanced Town] → [Endgame zones]
```

---

## Zone Connections from Starter Town

| Direction | Leads To | Discovery | Gate |
|-----------|----------|-----------|------|
| North | The Long Road (800) | Obvious | Quest-locked |
| South | Crystal Caves (500) | Main exit | Open |
| East | Bladeworks (600) | Signposted | Open (distant) |
| West | Minneapolis (200) | Neighborhood | Open |
| Northwest | Bellevue (300) | Neighborhood | Open |
| Up/Hidden | Squirrel Tree (400) | Search check | Hidden |

---

## Migration Renumbering

| Entity | Old IDs | New IDs |
|--------|---------|---------|
| Squirrel Tree rooms | 4001-4007 | 400-406 |
| Squirrel Tree mobs | 82-83 | 20-21 |
| Crystal Caves rooms | 2001-2018 | 500-517 |
| Crystal Caves mobs | 65-71 | 30-36 |
| Bladeworks rooms | 3001-3020 | 600-619 |
| Bladeworks mobs | 75-81 | 40-46 |

---

*Last updated: 2026-01-20*
*Maintainer: archie*
