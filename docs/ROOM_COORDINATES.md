# Room Coordinate Grid Documentation

This document maps all rooms in the portalis world to their x-y-z coordinates, calculated by traversing exits from a reference point. This helps debug mapper issues (Issue #18).

## Coordinate System

- **Origin**: Town Square (Room 100) at (0, 0, 0)
- **Direction mapping**:
  - North: y+1
  - South: y-1
  - East: x+1
  - West: x-1
  - Northeast: x+1, y+1 (FULL grid step in both directions)
  - Northwest: x-1, y+1
  - Southeast: x+1, y-1
  - Southwest: x-1, y-1
  - Up: z+1
  - Down: z-1

**Critical Rule**: Diagonal movement is a FULL grid step in both x and y directions.

---

## Zone: Starter Town (Rooms 100-151)

### Room Coordinates

| Room ID | Name | X | Y | Z | Exits |
|---------|------|---|---|---|-------|
| 100 | Town Square | 0 | 0 | 0 | N:101, S:500, E:103, W:600, NE:105, SE:104, NW:102, SW:110 |
| 101 | Training Yard | 0 | 1 | 0 | S:100, N:120 |
| 102 | Wizard's Study | -1 | 1 | 0 | SE:100, W:130 |
| 103 | Ranger Outpost | 1 | 0 | 0 | W:100, E:600(BF) |
| 104 | General Store | 1 | -1 | 0 | NW:100 |
| 105 | The Snarky Squirrel Inn | 1 | 1 | 0 | SW:100, Up:400(ST) |
| 110 | Market Street | -1 | -1 | 0 | NE:100, S:111, W:150, E:151 |
| 111 | Market Square | -1 | -2 | 0 | N:110, S:114, E:113, W:112 |
| 112 | Ironheart Smithy | -2 | -2 | 0 | E:111 |
| 113 | Steelguard Armory | 0 | -2 | 0 | W:111 |
| 114 | First Bank of Frostfell | -1 | -3 | 0 | N:111 |
| 120 | Cobblestone Lane | 0 | 2 | 0 | S:101, N:121, E:122 |
| 121 | Hearthstone Cottage | 0 | 3 | 0 | S:120, W:123 |
| 122 | Wisteria Cottage | 1 | 2 | 0 | W:120, NW:124 |
| 123 | Chapel of the Quiet Star | -1 | 3 | 0 | E:121, N:124 |
| 124 | Garden Commons | -1 | 4 | 0 | S:123, SE:122 |
| 130 | Adventurer's Guild Hall | -2 | 1 | 0 | E:102, N:131, W:132 |
| 131 | Quest Board Chamber | -2 | 2 | 0 | S:130 |
| 132 | Guild Master's Office | -3 | 1 | 0 | E:130, N:133 |
| 133 | Guild Vault Antechamber | -3 | 2 | 0 | S:132 |
| 150 | The Bubbling Cauldron | -2 | -1 | 0 | E:110 |
| 151 | First Bank of Starter Town | 0 | -1 | 0 | W:110 |

### ASCII Grid (Z=0)

```
Y
4       [124]
         |
3       [123]---[121]
               |
2  [131] [130]-[120]---[122]
    |      |    |         \
1  [132]-[130]-[102]---[100]---[105]
    |           \     / | \     |
0         [150]-[110][100][103] |
               |    \     |
-1        [112][111][113][151] [104]
               |
-2            [114]

-3  -2   -1    0    1    2    3    X
```

**Note**: Grid simplified to show connections. Room 100 is origin (0,0).

### Detailed ASCII Map

```
                            [124] Garden Commons
                           /     \
                    [123] Chapel  [122] Wisteria
                       |           |
            [131]   [121] Cottage  |
              |       |            |
[133]---[132]---[130]---[120] Lane-+
         |       |       |
        GLD   [102] Wiz [101] Train
               \         |
[150]---[110]---+--[100]-+---[103]----> To Bladeworks (600)
 Pot     Mkt     \   |   /    Ranger
                  \  |  /
                   [104] Store
                    |
[112]---[111]---[113]
Smith   MktSq   Armory
          |
        [114] Bank

        [151] Bank (E of 110)

        v To Crystal Caves (500)
```

### Coordinate Validation Issues - Starter Town

1. **ISSUE**: Room 151 (Bank) position conflict
   - 151 is EAST of 110, placing it at (0, -1)
   - But 104 (General Store) is at (1, -1) via SE from 100
   - Room 151 should be at (0, -1), Room 104 at (1, -1) - NO CONFLICT

2. **ISSUE**: Room 124 connection consistency
   - 124 is N of 123 (placing it at -1, 4)
   - 124 also has SE exit to 122 (which is at 1, 2)
   - SE from (-1, 4) should lead to (0, 3), NOT (1, 2)
   - **CONFLICT**: The SE exit from 124 to 122 implies 122 is at (0, 3), but 122 is actually at (1, 2)
   - **SEVERITY**: HIGH - This is a mapper-breaking coordinate mismatch

---

## Zone: Squirrel Tree (Rooms 400-406)

### Room Coordinates

The Squirrel Tree is a vertical zone accessed via the Snarky Squirrel Inn (Room 105).

| Room ID | Name | X | Y | Z | Exits |
|---------|------|---|---|---|-------|
| 400 | Tree Entry Hall | 1 | 1 | 1 | Down:105(ST), Up:401 |
| 401 | Trunk Ascent | 1 | 1 | 2 | Down:400, Up:402 |
| 402 | Mid-Trunk Chamber | 1 | 1 | 3 | Down:401, Up:403, E:404, W:405 |
| 403 | Upper Trunk | 1 | 1 | 4 | Down:402, Up:406 |
| 404 | Eastern Cache | 2 | 1 | 3 | W:402 |
| 405 | Western Cache | 0 | 1 | 3 | E:402 |
| 406 | Canopy Pinnacle | 1 | 1 | 5 | Down:403 |

### ASCII Grid (Side View - X/Z plane at Y=1)

```
Z
5        [406] Pinnacle
          |
4        [403] Upper Trunk
          |
3  [405]-[402]-[404]
   West   |    East
          |
2        [401]
          |
1        [400]
          |
0   -----[105] Inn (Starter Town)-----

    0     1     2                    X
```

### Coordinate Validation Issues - Squirrel Tree

**No issues found.** The vertical structure is consistent.

---

## Zone: Crystal Caves (Rooms 500-517)

### Room Coordinates

Crystal Caves is accessed via the SOUTH exit from Town Square (100).

| Room ID | Name | X | Y | Z | Exits |
|---------|------|---|---|---|-------|
| 500 | Cave Entrance | 0 | -1 | 0 | N:100(ST), S:501 |
| 501 | Twilight Passage | 0 | -2 | 0 | N:500, SW:502, SE:503 |
| 502 | Luminous Grotto | -1 | -3 | 0 | NE:501, S:504 |
| 503 | Mushroom Garden | 1 | -3 | 0 | NW:501, S:505 |
| 504 | Crystal Stream | -1 | -4 | 0 | N:502, S:506, E:505(secret) |
| 505 | Spore Hollow | 1 | -4 | 0 | N:503, S:507, W:504(secret) |
| 506 | Reflecting Pool | -1 | -5 | 0 | N:504, E:507, S:508 |
| 507 | Glowcap Grove | 1 | -5 | 0 | N:505, W:506, S:509 |
| 508 | Singing Crystals | -1 | -6 | 0 | N:506, S:510, E:509 |
| 509 | Fungal Cathedral | 1 | -6 | 0 | N:507, W:508, S:511 |
| 510 | Underground River | -1 | -7 | 0 | N:508, E:511, SE:512 |
| 511 | Bioluminescent Beach | 1 | -7 | 0 | N:509, W:510, SW:512 |
| 512 | Crystal Heart | 0 | -8 | 0 | NW:510, NE:511, W:513, E:514, S:515 |
| 513 | Geode Chamber | -1 | -8 | 0 | E:512 |
| 514 | Seer's Alcove | 1 | -8 | 0 | W:512 |
| 515 | Deep Hollow | 0 | -9 | 0 | N:512, S:516 |
| 516 | Matriarch's Antechamber | 0 | -10 | 0 | N:515, S:517 |
| 517 | Matriarch's Throne | 0 | -11 | 0 | N:516 |

### ASCII Grid (Z=0)

```
Y
-1                [500] Entrance
                    |
-2                [501]
                 /     \
-3          [502]       [503]
              |           |
-4          [504]---(s)--[505]
              |           |
-5          [506]-------[507]
              |           |
-6          [508]-------[509]
              |           |
-7          [510]-------[511]
               \         /
-8        [513]-[512]-[514]
                  |
-9              [515]
                  |
-10             [516]
                  |
-11             [517] Throne

       -1    0    1          X
```

### Coordinate Validation Issues - Crystal Caves

1. **ISSUE**: Secret passage coordinate mismatch
   - Room 504 at (-1, -4) has secret E exit to 505
   - Room 505 at (1, -4) has secret W exit to 504
   - **PROBLEM**: E from (-1, -4) should be (0, -4), not (1, -4)
   - **CONFLICT**: There's a missing room at (0, -4) or the secret passage spans 2 grid squares
   - **SEVERITY**: MEDIUM - Secret passages may not display correctly on mapper

2. **ISSUE**: Room 510 and 511 to 512 connections
   - 510 at (-1, -7) has SE exit to 512
   - SE from (-1, -7) = (0, -8) - this is correct for 512
   - 511 at (1, -7) has SW exit to 512
   - SW from (1, -7) = (0, -8) - this is also correct for 512
   - **NO CONFLICT**: Both diagonals correctly meet at 512

3. **ISSUE**: Room 506 and 507 lateral connection
   - 506 at (-1, -5) has E exit to 507
   - 507 at (1, -5) has W exit to 506
   - **PROBLEM**: E from (-1, -5) should be (0, -5), not (1, -5)
   - **CONFLICT**: Missing room at (0, -5) or passage spans 2 squares
   - **SEVERITY**: MEDIUM

4. **ISSUE**: Room 508 and 509 lateral connection (same pattern)
   - 508 at (-1, -6) has E exit to 509 at (1, -6)
   - **CONFLICT**: Same as above - 2-square gap
   - **SEVERITY**: MEDIUM

---

## Zone: Bladeworks Foundry (Rooms 600-619)

### Room Coordinates

Bladeworks Foundry is accessed via the WEST exit from Town Square (100).

| Room ID | Name | X | Y | Z | Exits |
|---------|------|---|---|---|-------|
| 600 | Foundry Entrance | -1 | 0 | 0 | E:100(ST), S:601 |
| 601 | The Receiving Bay | -1 | -1 | 0 | N:600, S:602, E:603 |
| 602 | The Smelting Chamber | -1 | -2 | 0 | N:601, S:604, W:605 |
| 603 | Supply Corridor | 0 | -1 | 0 | W:601, E:606 |
| 604 | The Grinding Hall | -1 | -3 | 0 | N:602, S:607, E:608 |
| 605 | Mold Storage | -2 | -2 | 0 | E:602 |
| 606 | Component Stockpile | 1 | -1 | 0 | W:603 |
| 607 | Quality Control Chamber | -1 | -4 | 0 | N:604, S:609 |
| 608 | Tempering Pools | 0 | -3 | 0 | W:604, S:610 |
| 609 | The Assembly Line | -1 | -5 | 0 | N:607, S:611 |
| 610 | Foreman's Office | 0 | -4 | 0 | N:608 |
| 611 | The Dueling Gallery | -1 | -6 | 0 | N:609, S:612(secret) |
| 612 | The Blade Gate | -1 | -7 | 0 | N:611, S:613 |
| 613 | The Proving Grounds | -1 | -8 | 0 | N:612, S:614, E:615 |
| 614 | The Weapon Vault | -1 | -9 | 0 | N:613, W:616 |
| 615 | Steam Works | 0 | -8 | 0 | W:613, S:617 |
| 616 | The Blade Garden | -2 | -9 | 0 | E:614, S:618 |
| 617 | The Control Nexus | 0 | -9 | 0 | N:615, S:619 |
| 618 | The Master's Study | -2 | -10 | 0 | N:616, E:619 |
| 619 | Hall of Endless Blades | 0 | -10 | 0 | W:618, N:617, S:5001 |

### ASCII Grid (Z=0)

```
Y
0     [600]------> To Town Square (100)
        |
-1    [601]---[603]---[606]
        |
-2  [605]-[602]
           |
-3       [604]---[608]
           |       |
-4       [607]   [610]
           |
-5       [609]
           |
-6       [611] Dueling Gallery (TRAINER)
           | (secret)
-7       [612] Blade Gate
           |
-8       [613]---[615]
           |       |
-9  [616]-[614]  [617]
     |           |
-10 [618]------[619] Hall of Endless Blades (BOSS)
                 |
                 v To unknown zone (5001)

    -2   -1    0    1          X
```

### Coordinate Validation Issues - Bladeworks Foundry

1. **ISSUE**: Room 103 (Ranger Outpost) connection
   - Room 103 is at (1, 0) in Starter Town
   - Room 103 has E exit to Room 600 (zone: Bladeworks Foundry)
   - But Room 600 is at (-1, 0) from Town Square
   - **CONFLICT**: Room 100 has W exit to 600, and Room 103 has E exit to 600
   - This means 600 is both WEST of 100 AND EAST of 103
   - But 103 is EAST of 100, so 600 cannot be EAST of 103
   - **SEVERITY**: HIGH - Cross-zone connection error

2. **ISSUE**: Missing Room 5001
   - Room 619 has S exit to room 5001 (no zone specified)
   - Room 5001 does not exist in current world files
   - **SEVERITY**: LOW - Likely planned expansion

---

## Cross-Zone Connection Summary

| From Zone | From Room | Direction | To Zone | To Room | Status |
|-----------|-----------|-----------|---------|---------|--------|
| Starter Town | 100 | South | Crystal Caves | 500 | OK |
| Starter Town | 100 | West | Bladeworks Foundry | 600 | OK |
| Starter Town | 103 | East | Bladeworks Foundry | 600 | **ERROR** |
| Starter Town | 105 | Up | Squirrel Tree | 400 | OK |
| Crystal Caves | 500 | North | Starter Town | 100 | OK |
| Bladeworks Foundry | 600 | East | Starter Town | 100 | OK |
| Squirrel Tree | 400 | Down | Starter Town | 105 | OK |

---

## Summary of Issues Found

### HIGH Severity
1. **Room 124 to 122 diagonal**: SE from 124 should reach (0,3), but 122 is at (1,2)
2. **Room 103 to 600 connection**: 600 cannot be both W of 100 and E of 103

### MEDIUM Severity
1. **Crystal Caves lateral connections**: Rooms 504-505, 506-507, 508-509 have 2-square gaps
   - E/W exits skip a grid square, which may cause mapper display issues

### LOW Severity
1. **Room 619 to 5001**: Target room doesn't exist (planned expansion)

---

## Recommendations

1. **Fix Room 103**: Remove the E exit to 600, or add an intermediate corridor
2. **Fix Room 124**: Either:
   - Change SE exit to just S (pointing to 122 at a new coordinate), or
   - Add intermediate room at (0,3) between 124 and 122
3. **Review Crystal Caves**: Decide if 2-square lateral passages are intentional or need intermediate rooms

---

*Document generated by galstaff, Sorcerer of Light, for Issue #18 mapper debugging.*
*"The dungeon has been mapped! Roll for perception to spot the coordinate conflicts!"*
