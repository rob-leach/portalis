# Coordinate System Design

*A system to prevent impossible geometry in zone building.*

---

## The Problem

Galstaff created zones with coordinate conflicts:
- Room 600 connected as both WEST of Room 100 AND EAST of Room 103 (impossible)
- Diagonal exits that don't follow the full-step rule
- Lateral connections spanning multiple grid squares

The mapper broke because coordinates couldn't be reconciled.

---

## Core Rules

### Rule 1: Absolute Coordinates

Every room has exactly one (x, y, z) coordinate. This is non-negotiable.

**Origin**: Town Square (Room 100) is always (0, 0, 0).

### Rule 2: Direction Deltas

Movement in any direction has a fixed coordinate delta:

| Direction | Delta | Inverse |
|-----------|-------|---------|
| North | (0, +1, 0) | South |
| South | (0, -1, 0) | North |
| East | (+1, 0, 0) | West |
| West | (-1, 0, 0) | East |
| Up | (0, 0, +1) | Down |
| Down | (0, 0, -1) | Up |

### Rule 3: Diagonal = Full Step in Both Axes

Diagonals are NOT half-steps. They move a full unit in both x and y:

| Direction | Delta | Inverse |
|-----------|-------|---------|
| Northeast | (+1, +1, 0) | Southwest |
| Northwest | (-1, +1, 0) | Southeast |
| Southeast | (+1, -1, 0) | Northwest |
| Southwest | (-1, -1, 0) | Northeast |

**Why?** MUD mappers display rooms on a grid. Half-step diagonals create overlapping coordinates.

### Rule 4: Bidirectional Consistency

If Room A has exit X to Room B, then Room B MUST have the inverse exit back to A.

```
Room 100 (0,0) --EAST--> Room 103 (1,0)
Room 103 (1,0) --WEST--> Room 100 (0,0)  // Required
```

Violation example from the current world:
```
Room 100 (0,0) --WEST--> Room 600 (-1,0)  // OK
Room 103 (1,0) --EAST--> Room 600         // IMPOSSIBLE: E from (1,0) = (2,0), not (-1,0)
```

### Rule 5: Long Passages Need Intermediate Rooms

If two rooms are more than 1 grid unit apart, you need intermediate rooms.

**Bad**: Room at (-1, -4) connects EAST to room at (1, -4)
**Good**: Add a room at (0, -4) in between

Exception: Secret passages can be non-Euclidean (documented as such).

---

## Zone Planning Process

### Step 1: Sketch the Grid FIRST

Before creating any room files, draw an ASCII grid:

```
Example: A simple 4-room zone

Y
2    [102]
      |
1    [101]
      |
0    [100]---[103]
     origin

     0    1         X
```

Include:
- Room IDs at grid positions
- All exits drawn as lines
- Coordinate labels on axes
- Zone entry/exit points marked

### Step 2: Assign Coordinates to Every Room

Fill out a coordinate table:

| Room ID | Name | X | Y | Z | Exits |
|---------|------|---|---|---|-------|
| 100 | Origin | 0 | 0 | 0 | N:101, E:103 |
| 101 | North Room | 0 | 1 | 0 | S:100, N:102 |
| 102 | Far North | 0 | 2 | 0 | S:101 |
| 103 | East Room | 1 | 0 | 0 | W:100 |

### Step 3: Validate Before Creating Files

Check each exit:
1. Does applying the direction delta to source coordinate equal target coordinate?
2. Does the target room have the inverse exit back?
3. Do cross-zone exits respect both zones' coordinate systems?

### Step 4: Create Room Files

Only after validation passes.

---

## Validation Checklist

Run this checklist before committing any zone changes:

```
[ ] 1. Grid sketch exists for zone
[ ] 2. Every room has (x, y, z) assigned
[ ] 3. No coordinate collisions (two rooms at same position)
[ ] 4. All exits validated:
      [ ] N/S exits: y changes by 1, x unchanged
      [ ] E/W exits: x changes by 1, y unchanged
      [ ] Diagonals: both x and y change by 1
      [ ] Up/Down: z changes by 1, x/y unchanged
[ ] 5. All exits are bidirectional (or documented as one-way)
[ ] 6. Cross-zone connections validated from both sides
[ ] 7. No multi-square gaps without intermediate rooms
```

---

## Cross-Zone Connections

When connecting to another zone, verify:

1. **Target coordinate**: Where does the exit lead in absolute coordinates?
2. **Return path**: Does the target zone have the correct inverse exit?
3. **Zone boundary**: Document the seam in both zones' coordinate tables.

Example (correct):
```
Starter Town (100) at (0, 0) --SOUTH--> Crystal Caves (500) at (0, -1)
Crystal Caves (500) at (0, -1) --NORTH--> Starter Town (100) at (0, 0)
```

Example (broken - the 103 to 600 problem):
```
Town Square (100) at (0, 0)   --WEST-->  Bladeworks (600) at (-1, 0)  // OK
Ranger Outpost (103) at (1, 0) --EAST--> Bladeworks (600) at (-1, 0)  // WRONG
                                         // E from (1,0) should be (2,0)
```

---

## Special Cases

### Teleportation / Magical Portals

Non-Euclidean connections are allowed if:
1. Explicitly marked as `type: portal` or similar
2. Not included in mapper coordinate calculations
3. Documented in zone notes

### One-Way Exits

Allowed for:
- Trap doors (fall down, can't climb back)
- Magical barriers
- Quest gates

Must be documented. The mapper should handle these gracefully.

### Secret Passages

Can span multiple squares if:
1. Marked as secret (search required)
2. Documented as non-Euclidean in zone notes
3. The mapper is told to ignore them for layout purposes

---

## Room ID Allocation Reminder

From ROOM_ALLOCATION.md:

| Block | Zone |
|-------|------|
| 100-199 | Starter Town |
| 200-299 | Minneapolis |
| 300-399 | Bellevue |
| 400-499 | Squirrel Tree |
| 500-599 | Crystal Caves |
| 600-699 | Bladeworks Foundry |
| 700-799 | Future Zone |
| 800-899 | The Long Road |
| 900-999 | Advanced Town |

Stay within your block. If you need more than 100 rooms, coordinate with archie.

---

## Future: Automated Validation

A validation script should eventually:

1. Parse all room YAML files
2. Build a coordinate graph
3. Check all rules automatically
4. Block commits with violations

Until then, use the manual checklist.

---

## Applying to Current Issues

### Issue #18: Map Display Broken at West Gate

Root cause: Room 103 has E exit to 600, but 600 is WEST of 100.

**Fix**: Remove the E exit from 103 to 600. The Bladeworks connection should only be through Room 100's W exit.

### Crystal Caves Lateral Gaps

Rooms 504-505, 506-507, 508-509 have E/W connections spanning 2 squares.

**Options**:
1. Add intermediate rooms at (0, -4), (0, -5), (0, -6)
2. Mark these as "long passages" in the zone description (mapper treats as 2 steps)
3. Accept non-Euclidean layout for caves (underground doesn't follow surface rules)

Recommendation: Option 3 with documentation. Caves can be weird.

### Room 124 to 122 Diagonal

SE from 124 at (-1, 4) should reach (0, 3), but 122 is at (1, 2).

**Fix**: Either:
- Change exit from SE to a two-step path (S then SE)
- Add intermediate room at (0, 3)
- Remove the diagonal entirely

---

## Summary

1. **Every room has one coordinate**
2. **Directions have fixed deltas**
3. **Diagonals are full steps**
4. **All exits must be bidirectional (or documented)**
5. **Sketch the grid before building**
6. **Validate before committing**

*"The world is a graph. If the edges don't connect to the right nodes, the player falls through."*

---

*Document created by archie, world systems architect.*
*2026-01-21*
