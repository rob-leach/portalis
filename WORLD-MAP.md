# GoMud World Map

## Area Summary

**As of January 2026, the world map has been trimmed to only include the CORE regions. All optional, special, and WIP areas have been removed.**

| Area                | Directory           | Room Count | Level Range | Biome | Purpose/Theme                        |
|---------------------|--------------------|------------|-------------|-------|--------------------------------------|
| **Tutorial**        | `tutorial/`        | 4          | N/A         | N/A   | New player introduction              |
| **Frostfang**       | `frostfang/`       | ~95        | 1-5         | city  | Starting city hub                    |
| **Frostfang Slums** | `frostfang_slums/` | ~58        | 5-10        | city  | Low-level zone (PVP enabled)         |
| **Catacombs**       | `catacombs/`       | ~70        | 8-13        | dungeon | Undead dungeon beneath Sanctuary     |
| **Whispering Wastes** | `whispering_wastes/` | ~51     | 5-25        | snow  | Frozen wasteland west of Frostfang   |

**TOTAL: ~278 room files across 5 areas**

---

## Core Zones (Starter/Tutorial Areas)

These zones form the essential early-game experience:
- **Tutorial**: 4 rooms teaching basic mechanics
- **Frostfang**: Central hub city (Town Square = Room 1)
- **Frostfang Slums**: Level 5-10 progression
- **Catacombs**: Level 8-13 dungeon
- **Whispering Wastes**: Level 5-25 wilderness

---

## World Connection Overview

```mermaid
flowchart TD
    subgraph SYSTEM["System Zones"]
        VOID["-1: Nowhere<br/>The Void"]
        SHADOW["75: Shadow Realm<br/>Death Waiting Room"]
    end

    subgraph TUTORIAL["Tutorial Zone"]
        T900["900: Learning to Look"]
        T901["901: Learning about You"]
        T902["902: Learning to Fight"]
        T903["903: Training Complete"]
        T900 -->|east| T901
        T901 -->|south| T902
        T902 -->|west| T903
    end

    subgraph FROSTFANG["Frostfang (City Hub)"]
        FF1["1: Town Square"]
        FF_CASTLE["36-53: Castle Area"]
        FF_TEMPLE["17-18: Sanctuary"]
        FF_SHOPS["61-76: Shops District"]
        FF_WEST["7-11: West Road"]
        FF_EAST["54-60: East Gate"]
        FF_SOUTH["278-305: Southern/Old Road"]
        FF_BATTLE["782-829: Battlements"]
        FF_ALLEY["21-34: Dark Alleys"]
        FF1 --> FF_CASTLE
        FF1 --> FF_TEMPLE
        FF1 --> FF_SHOPS
        FF1 --> FF_WEST
        FF1 --> FF_EAST
        FF1 --> FF_SOUTH
    end

    subgraph SLUMS["Frostfang Slums (Lvl 5-10)"]
        SL_ENTRY["434-439: Entry Streets"]
        SL_DEEP["440-491: Deep Slums"]
        SL_ENTRY --> SL_DEEP
    end

    subgraph CATACOMBS["Catacombs (Lvl 8-13)"]
        CAT_ENTRY["32: Entrance"]
        CAT_MAIN["77-138: Tunnels & Crypts"]
        CAT_LICH["138: Lich Chamber"]
        CAT_TRAIN["160-165: Dark Acolyte"]
        CAT_ENTRY --> CAT_MAIN
        CAT_MAIN --> CAT_LICH
        CAT_MAIN --> CAT_TRAIN
    end

    subgraph WASTES["Whispering Wastes (Lvl 5-25)"]
        WW_ENTRY["168: Frozen Wasteland"]
        WW_MAIN["169-216: Ice Fields"]
        WW_CAVE["217: Cave Entrance"]
        WW_KEEP["202: Stormwatchers Entrance"]
        WW_ENTRY --> WW_MAIN
        WW_MAIN --> WW_CAVE
        WW_MAIN --> WW_KEEP
    end

    %% CROSS-AREA CONNECTIONS (CORE ONLY)
    FF_ALLEY -->|"16 south"| SL_ENTRY
    FF_ALLEY -->|"34 south"| SL_ENTRY
    FF_TEMPLE -->|"31 down"| CAT_ENTRY
    FF_WEST -->|"167 west"| WW_ENTRY

    style TUTORIAL fill:#90EE90
    style FROSTFANG fill:#87CEEB
    style SLUMS fill:#FFB6C1
    style CATACOMBS fill:#DDA0DD
    style WASTES fill:#E0FFFF
```

---

## Detailed Area Maps

### Tutorial (Rooms 900-903)

```mermaid
flowchart LR
    subgraph Tutorial["Tutorial Zone"]
        900["900: Learning to Look<br/><i>Start here</i>"]
        901["901: Learning about You<br/><i>Character status</i>"]
        902["902: Learning to Fight<br/><i>Combat training</i>"]
        903["903: Training Complete<br/><i>Graduation</i>"]

        900 -->|east| 901
        901 -->|south| 902
        902 -->|west| 903
    end

    style 900 fill:#90EE90
    style 903 fill:#FFD700
```

---

### Frostfang (City Hub - ~95 rooms)

```mermaid
flowchart TD
    subgraph Frostfang["Frostfang - Starting City"]
        1["1: Town Square<br/><i>Central Hub</i>"]

        subgraph North["North - Castle"]
            2["2: Cobblestone Way"]
            3["3: Cobblestone Way"]
            4["4: Cobblestone Way"]
            5["5: Cobblestone Way"]
            36["36: Castle"]
            50["50: Throne Room"]
        end

        subgraph West["West - Road & Gate"]
            7["7: West Road"]
            8["8: West Road"]
            9["9: West Road"]
            11["11: West Road"]
            35["35: West Gate"]
            167["167: Outside West Gate"]
        end

        subgraph South["South - Temple & Alleys"]
            12["12: Beggars Lane"]
            14["14: Beggars Lane"]
            16["16: Beggars Lane"]
            17["17: Sanctuary Doors"]
            18["18: Sanctuary"]
            31["31: Shadow Vault"]
        end

        subgraph East["East - Shops & Gate"]
            54["54: Eastwind"]
            55["55: Eastwind"]
            57["57: Frostfire Inn Area"]
            60["60: East Gate"]
            61["61: Frostfire Inn"]
            62["62: Icy Emporium"]
            63["63: Steelwhisper Armory"]
        end

        subgraph SouthEast["Southeast - Old Road"]
            278["278: Old Road"]
            296["296: Southern Road"]
            299["299: Old Road"]
        end

        1 -->|north| 2
        2 -->|north| 3
        3 -->|north| 4
        4 -->|north| 5
        5 -->|north| 36
        36 -->|"..."| 50

        1 -->|west| 7
        7 -->|west| 8
        8 -->|west| 9
        9 -->|west| 11
        11 -->|west| 35
        35 -->|west| 167

        1 -->|south| 12
        12 -->|south| 14
        14 -->|south| 16
        12 -->|east| 17
        17 -->|east| 18
        18 -->|"down"| 31

        1 -->|east| 54
        54 -->|east| 55
        55 -->|"..."| 60

        60 -->|east| 278
        278 -->|"..."| 296
        278 -->|"..."| 299
    end

    %% External connections
    167 -.->|"to Whispering Wastes"| EXT1["168"]
    31 -.->|"down to Catacombs"| EXT2["32"]
    16 -.->|"south to Slums"| EXT3["439"]
    296 -.->|"southeast to Frost Lake"| EXT4["304"]
    299 -.->|"east to Dark Forest"| EXT5["300"]

    style 1 fill:#FFD700
    style 167 fill:#E0FFFF
    style 31 fill:#DDA0DD
    style 16 fill:#FFB6C1
```

---

### Catacombs (Dungeon - ~70 rooms)

```mermaid
flowchart TD
    subgraph Catacombs["Catacombs - Dungeon Level 8-13"]
        32["32: Entrance<br/><i>From Shadow Vault</i>"]

        subgraph MainTunnels["Main Tunnels"]
            77["77-99: Tunnels"]
            100["100-130: Deep Catacombs"]
        end

        subgraph SpecialRooms["Special Areas"]
            118["118: Junction"]
            136["136: Above Lich"]
            138["138: Lich Chamber<br/><i>Boss</i>"]
            160["160: Dark Acolyte Trainer<br/><i>Scribe Training</i>"]
        end

        32 -->|east| 77
        32 -->|north| 118
        32 -->|west| 81["81"]
        77 --> 100
        100 --> 136
        136 -->|"down (trap)"| 138
        110["110"] -->|west| 160
    end

    UP["31: Shadow Vault<br/>(Frostfang)"] -.->|up| 32

    style 32 fill:#90EE90
    style 138 fill:#FF6347
    style 160 fill:#FFD700
```

---

### Frostfang Slums (Lvl 5-10 - ~58 rooms)

```mermaid
flowchart TD
    subgraph Slums["Frostfang Slums - PVP Zone"]
        direction TB

        subgraph Entry["Entry Points"]
            434["434: Poorly Lit Street<br/><i>From Dark Alley 34</i>"]
            439["439: Poorly Lit Street<br/><i>From Beggars Lane 16</i>"]
        end

        subgraph Streets["Slum Streets (440-491)"]
            440["440+: Twisted Streets"]
            460["460+: Deep Slums"]
            480["480+: Dangerous Area"]
        end

        434 -->|west| 435["435"]
        435 --> 440
        439 -->|west| 440
        440 --> 460
        460 --> 480
    end

    FF16["16: Beggars Lane"] -.->|south| 439
    FF34["34: Dark Alley"] -.->|south| 434

    style 434 fill:#FFB6C1
    style 439 fill:#FFB6C1
```

---

### Whispering Wastes (Lvl 5-25 - ~51 rooms)

```mermaid
flowchart TD
    subgraph Wastes["Whispering Wastes - Frozen Wilderness"]
        168["168: Frozen Wasteland<br/><i>Entry from Frostfang</i>"]

        subgraph Eastern["Eastern Wastes (Lvl 5-15)"]
            169["169-180: Ice Fields"]
        end

        subgraph Western["Western Wastes (Lvl 15-25)"]
            190["190-216: Deep Wastes"]
            202["202: Stormwatchers Keep<br/><i>Gate to Keep</i>"]
            217["217: Cave Entrance<br/><i>To Mirror Caves</i>"]
        end

        168 -->|northwest| 169
        169 --> 190
        190 --> 202
        190 --> 217
    end

    FF167["167: Outside West Gate"] -.->|west| 168
    217 -.->|cave| MC218["218: Mirror Caves"]
    202 -.->|north| SK880["880: Stormwatchers Keep"]

    style 168 fill:#E0FFFF
    style 217 fill:#DDA0DD
    style 202 fill:#87CEEB
```

---

### Frost Lake (Lvl 10-20 - ~125 rooms)

```mermaid
flowchart TD
    subgraph Lake["Frost Lake Region"]
        304["304: Shore<br/><i>Entry Point</i>"]

        subgraph Shore["Lakeshore (306-350)"]
            306["306+: Western Shore"]
            330["330+: Eastern Shore"]
            361["361: Northern Shore"]
        end

        subgraph Water["Lake Waters"]
            370["370+: Frozen Surface"]
            734["734+: Lake Center"]
        end

        subgraph Island["Lake Island (Lvl 15-20)"]
            750["750+: Island"]
        end

        304 -->|north| 361
        304 -->|southwest| 306
        306 --> 330
        330 --> 370
        370 --> 734
        734 --> 750
    end

    FF296["296: Southern Road"] -.->|southeast| 304

    style 304 fill:#87CEEB
```

---

### Dark Forest (Lvl 10-30 - ~165 rooms)

```mermaid
flowchart TD
    subgraph Forest["Dark Forest"]
        300["300: Forest Road<br/><i>Entry from Old Road</i>"]

        subgraph Western["Western Forest (Lvl 10-20)"]
            385["385-430: Dark Woods"]
        end

        subgraph Eastern["Eastern Forest (Lvl 20-25)"]
            492["492-570: Deeper Woods"]
        end

        subgraph Spider["Spider Den (Lvl 25-30)"]
            560["560+: Spider Webs"]
        end

        subgraph Exit["Mountain Exit"]
            574["574: Forest Road<br/><i>To Stormshards</i>"]
        end

        300 -->|east| 385
        385 --> 492
        492 --> 560
        492 --> 574
    end

    FF299["299: Old Road"] -.->|east| 300
    574 -.->|east| SS575["575: Stormshards"]

    style 300 fill:#228B22
    style 574 fill:#A0522D
```

---

### Stormshards (Lvl 30-40 - ~40 rooms)

```mermaid
flowchart LR
    subgraph Mountains["Stormshards - Mountain Path"]
        575["575: Mountain Path<br/><i>Entry</i>"]
        580["580-590: Climbing"]
        600["600-609: Near Summit"]
        610["610: Mountain Top<br/><i>To Mystarion</i>"]

        575 --> 580
        580 --> 600
        600 --> 610
    end

    DF574["574: Dark Forest"] -.->|east| 575
    610 -.->|east| MY612["612: Mystarion"]

    style 575 fill:#A0522D
    style 610 fill:#FFD700
```

---

### Mystarion (Lvl 40-50 - ~130 rooms)

```mermaid
flowchart TD
    subgraph City["Mystarion - End-Game City"]
        612["612: Main Road<br/><i>Entry from Mountains</i>"]

        subgraph MainCity["City Districts (613-730)"]
            613["613+: City Streets"]
            650["650-680: Markets"]
            700["700-730: Noble District"]
        end

        subgraph Arena["Arena Complex (831-863)"]
            831["831+: Arena Exterior"]
            859["859: Grand Arcane Arena"]
            863["863: Arena Stands"]
        end

        612 --> 613
        613 --> 650
        650 --> 700
        700 --> 831
        831 --> 859
    end

    SS610["610: Mountain Top"] -.->|east| 612

    style 612 fill:#9370DB
    style 859 fill:#FF6347
```

---

### Special Zones

```mermaid
flowchart TB
    subgraph Special["Special & System Zones"]
        subgraph MirrorCaves["Mirror Caves (Lvl 25-30)"]
            218["218: Cave Entrance"]
            219["219-257: Cave System"]
            218 --> 219
        end

        subgraph Keep["Stormwatchers Keep (Lvl 15-25)"]
            880["880: Keep Entrance"]
            881["881-893: Keep Interior"]
            880 --> 881
        end

        subgraph Trash["Endless Trashheap"]
            139["139: The Wasteland<br/><i>Loot Goblin Domain</i>"]
            140["140-159: Trash Piles"]
            139 --> 140
        end

        subgraph System["System Zones"]
            75["75: Shadow Realm<br/><i>Death Recovery</i>"]
            -1["-1: Nowhere<br/><i>The Void</i>"]
            1000["1000: Sun Anvil<br/><i>Desert - WIP</i>"]
        end
    end

    WW217["217: Wastes"] -.-> 218
    WW202["202: Wastes"] -.-> 880

    style 139 fill:#FFD700
    style 75 fill:#808080
    style -1 fill:#2F4F4F,color:#FFF
    style 1000 fill:#F4A460
```

---

## Level Progression Path

```mermaid
flowchart LR
    T["Tutorial<br/>Lvl 0"] --> FF["Frostfang<br/>Lvl 1-5"]
    FF --> SL["Slums<br/>Lvl 5-10"]
    FF --> CAT["Catacombs<br/>Lvl 8-13"]
    FF --> WW["Wastes E<br/>Lvl 5-15"]

    SL --> CAT
    CAT --> WW

    FF --> FL["Frost Lake<br/>Lvl 10-15"]
    FF --> DF_W["Dark Forest W<br/>Lvl 10-20"]

    WW --> WW2["Wastes W<br/>Lvl 15-25"]
    WW2 --> SK["Stormwatchers<br/>Lvl 15-25"]
    WW2 --> MC["Mirror Caves<br/>Lvl 25-30"]

    FL --> FL2["Frost Lake Island<br/>Lvl 15-20"]

    DF_W --> DF_E["Dark Forest E<br/>Lvl 20-25"]
    DF_E --> DF_S["Spider Den<br/>Lvl 25-30"]
    DF_E --> SS["Stormshards<br/>Lvl 30-40"]

    SS --> MY["Mystarion<br/>Lvl 40-50"]

    style T fill:#90EE90
    style FF fill:#87CEEB
    style MY fill:#9370DB
```

---

## Core vs Optional Content

### CORE (Keep for Minimal Viable World)
- **Tutorial** (4 rooms): Essential onboarding
- **Frostfang** (~95 rooms): Central hub, shops, services
- **Frostfang Slums** (~58 rooms): Early progression, PVP area
- **Catacombs** (~70 rooms): First dungeon, undead theme
- **Whispering Wastes** (~51 rooms): Gateway to outer zones

### OPTIONAL (Can be replaced/removed)
- **Frost Lake** (~125 rooms): Side content, lake exploration
- **Dark Forest** (~165 rooms): Mid-game content, spider theme
- **Mirror Caves** (~40 rooms): Late mid-game dungeon
- **Stormshards** (~40 rooms): Transition zone to end-game
- **Stormwatchers Keep** (~14 rooms): Small dungeon in wastes
- **Mystarion** (~130 rooms): End-game city

### SPECIAL/SYSTEM
- **Endless Trashheap** (~21 rooms): Self-contained special zone
- **Shadow Realm** (1 room): Death mechanic
- **Nowhere** (1 room): System void
- **Sun Anvil** (1 room): Work in progress

---

## Notes

1. **Cross-area connections** are primarily one-way conceptually (you progress outward from Frostfang) but rooms have bidirectional exits
2. **Room IDs are not sequential within zones** - some zones have gaps or outliers
3. **Endless Trashheap has no external connections** - likely teleport/quest access only
4. **Sun Anvil is WIP** - has no exits defined yet
5. **Shadow Realm and Nowhere** are system zones for death/error states
