# Session 4: The Crafting Prophecy (2026-01-18)

*Galstaff adjusts his spectacles and unfurls ancient scrolls of MUD lore*

"HARK! The council has asked me to divine the future of crafting in portalis. I have SCRIED through the histories of a dozen MUDs and consulted the wisdom of the ages!"

## Quest: Research MUD Crafting Systems for Issue #14

**Status:** ✓ RESEARCH COMPLETE - 7 major sources consulted

**Output:**
- `/Users/i2pi/h/fun/bday2026/portalis/CRAFTING_RESEARCH.md` - Full research document
- `/Users/i2pi/h/fun/bday2026/portalis/ISSUE_14_COMMENT.md` - GitHub issue comment (ready to post)

---

## Key Prophecies Revealed

### The Central Binding Force

From my research, ONE TRUTH stands above all others: **item durability/decay is the skeleton key** that unlocks a balanced crafting economy.

**The Problem:** If crafted items don't break/decay, players eventually have "enough" gear and stop buying from crafters. When demand drops but supply is unlimited (crafters can make infinite items), the economy stagnates.

**The Solution:** Implement item durability. Players use items → durability decreases → items need repair → repair costs create money drain → ongoing demand for crafters' services.

This single mechanic is more important than anything else in balancing a crafting system.

### The Power Balance Prophecy

The ancients have spoken: **crafted items should be ROUGHLY EQUAL in power to dropped and quest-obtained gear at the same tier.**

**Bad approach:** Superior crafted items
- Creates mandatory crafting (everyone must craft or buy)
- Breaks economy (drops become worthless)
- Unfair to combat-focused players

**Recommended distribution at each level tier:**
- Dropped loot: 30-40% of best gear
- Crafted items: 30-40% of best gear
- Quest rewards: 20-30% of best gear

**Result:** Players choose based on playstyle and availability, not because one is obviously broken.

### The Three Common Curses

**1. The Grind Curse** (Example: CLOK)
- Problem: Skills feel tedious without visible progress
- Symptom: "Crafting for hours without seeing improvement"
- Impact: Players abandon the system
- Antidote:
  - Make progression visible (show XP %)
  - Vary the mechanics (not all button-mashing)
  - Keep time investment reasonable (minutes, not hours)
  - Give feedback ("Excellent work!" vs. "Crude result")

**2. The Hidden Knowledge Curse** (Example: Armageddon)
- Problem: Recipes impossible to discover
- Symptom: Players must ask others to learn ANY recipe
- Impact: Gatekeeping friction for new players
- Antidote:
  - Make recipes discoverable through NPCs
  - Document recipes in books/guides
  - Allow experimentation ("put items in container, see what happens")
  - Example: Discworld MUD got this RIGHT

**3. The Economic Stagnation Curse**
- Problem: Permanent items + unlimited supply = worthless crafting
- Math: Eventually everyone has enough gear → demand drops to zero → crafters can't sell anything → crafting becomes worthless
- Antidote:
  - Item decay/durability (creates ongoing demand)
  - Material scarcity (not infinite supply)
  - Crafting costs (time + resources)
  - Faucet/drain balance (money entering ≈ money leaving)

---

## Portalis's Advantages

✓ **Container-based recipe system** - Elegant, scriptable, discoverable (not hidden)
✓ **NPC system** - Can teach recipes and manage crafting services
✓ **JavaScript scripting** - Allows complex crafting mechanics and item quality tiers
✓ **Item and buff systems** - Can support durability tracking and repair mechanics

The foundation is already good! The key is avoiding pitfalls in implementation.

---

## Critical Success Factors (Pre-Launch Checklist)

- [ ] **Implement item durability/repair system** (do this FIRST - it's THE KEY)
- [ ] Test power level equivalence (crafted items ≈ dropped items ≈ quest items)
- [ ] Define material tiers and scarcity (copper < iron < steel, and not infinite)
- [ ] Make recipes discoverable (test with new players - can they find recipes intuitively?)
- [ ] Monitor economy balance (track faucets vs drains - are crafters becoming too wealthy?)
- [ ] Playtest with new players (does it feel grindy? Are recipes clear?)

---

## Implementation Roadmap

**Phase 1: Container Recipes** (Foundation exists)
- Container-based recipes working
- Recipe discovery through experimentation
- No skill progression yet

**Phase 2: Crafting Skills & Quality** (Next)
- Add crafting skill progression (1-10)
- Craft quality depends on crafter's skill level
- Recipe discovery through NPCs/books
- Failed crafts consume materials (prevents spam)

**Phase 3: Material Requirements & Gathering** (Later)
- Define material sources (mobs, gathering locations, merchants)
- Material tiers with scarcity
- Resource costs balanced to item power level

**Phase 4: Item Durability & Repair** (Critical for sustainability)
- Add durability system to crafted items
- Implement repair mechanics and costs
- Create money drain for economy stabilization

---

## Examples from Real MUDs

### Discworld MUD (Active, 20+ years) ✓ SUCCESSFUL

**What they got right:**
- Decentralized crafting (anyone can craft with proper skills)
- Multiple specializations (finesmithing, pottery, jewellery, music, healing)
- Workshop system (artisans maintain workshops; players pay fees)
- Natural progression (skills improve through meaningful use, not grinding)
- Optional playstyle (crafting is useful but not mandatory)

**Key insight:** "No forced grinding. Skills progress at a reasonable pace with meaningful use."

### Armageddon MUD (Closed 2024) ⚠ LESSONS LEARNED

**What they got right:**
- Merchant class with dedicated crafting
- Quality tiers (basic/standard/excellent) prevent sameness
- Skill-based progression feels rewarding

**What they got wrong:**
- Recipes were hidden (not documented in-game)
- New players couldn't discover recipes (had to ask other players)
- Gatekeeping friction for new player experience

### DartMUD (Player-run economy) ✓ BALANCED

**What they got right:**
- Crafting provides rewards competitive with other skills
- Comprehensive economic model (wood, metal, cloth, stone, ceramic)
- Multiple production tiers create sustained crafting demand
- Player-driven market (players source materials and sell finished goods)

**What they struggled with:**
- "Player-run economy has been the most problematic and contentious feature"
- Multiple tweaks needed to balance realism vs. gameplay
- Key lesson: Item decay/durability is critical for sustainability

---

## The Ultimate Wisdom

*"Crafting succeeds when it's useful but optional, rewarding but not mandatory, and economically balanced."*

### From the Faucet-Drain Principle

**Money Flow Model:**
- **Faucets** (money enters): Quest rewards, mob drops, NPC payments
- **Drains** (money exits): NPC vendors, taxes, repair costs
- **Goal**: Keep faucets ≈ drains to prevent inflation or deflation

**Application to crafting:**
- Crafting item costs are drains (reduce player gold)
- Repair costs are drains (reduce player gold)
- Crafting profit margins are faucets (give players gold)
- These must balance or economy stagnates

---

## Next Steps for Issue #14

When commenting on issue #14, highlight:

1. **Item durability is THE KEY** (most important finding from research)
   - Without it, economy will stagnate
   - With it, crafting becomes sustainable and valuable

2. **Container-based recipes are a good foundation**
   - Already discoverable (unlike Armageddon's hidden recipes)
   - Scriptable for complex mechanics
   - Can expand to NPC training and quality tiers

3. **Pre-launch testing is critical**
   - Playtest power level balance
   - Verify recipes are discoverable
   - Monitor economy for stagnation signs

4. **Avoid the three common curses**
   - Grind curse → visible progression, varied mechanics
   - Hidden knowledge curse → discoverable recipes
   - Economic stagnation curse → item decay + material scarcity

---

## Research Sources

1. [Discworld MUD Crafting System](http://discworld.atuin.net/lpc/playing/documentation.c?path=/concepts/crafting) - Active, well-designed
2. [Armageddon MUD Documentation](http://armageddonmud.org/help/view/Crafting) - Merchant class example
3. [DartMUD Economy Design](https://mud.fandom.com/wiki/DartMUD) - Comprehensive medieval economy
4. [MMORPG Economy Analysis](https://www.terminally-incoherent.com/blog/2010/07/08/mmo-crafting-economies/) - Item decay problem
5. [Pantheon Forums - Item Balance](https://seforums.pantheonmmo.com/content/forums/topic/1550/armor-quested-vs-crafted-vs-loot) - Power tier distribution
6. [Massively Overpowered - Economy Building](https://massivelyop.com/2017/07/20/massively-overthinking-building-a-better-mmorpg-economy/) - Faucet-drain model
7. [Top Mud Sites - Crafting Discussion](https://www.topmudsites.com/forums/showthread.php?t=6967) - Community perspectives

---

*"The prophecy is complete. The path is clear. Now comes the EXECUTION!"*

*Galstaff rolls a natural 20 on his Divination check and closes his spellbook with satisfaction.*

*"Item durability is the key. Never forget this. It is the difference between a thriving crafting economy and a stagnant wasteland of worthless items."*

---
