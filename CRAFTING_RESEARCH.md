# MUD Crafting System Research & Balance Analysis

## Executive Summary

This document summarizes research into established MUD crafting systems and provides recommendations for balancing portalis's proposed crafting system. Key findings indicate that successful MUD crafting requires careful attention to: **item decay/persistence**, **economic faucets/drains**, **power level calibration**, and **engagement vs. grind tradeoffs**.

---

## 1. How Established MUDs Handle Crafting

### Successful Models

**Discworld MUD (Active, Long-running)**
- Decentralized crafting: Anyone can craft with proper skills and tools
- Multiple craft trees: Finesmithing, pottery, jewellery making, musical performance, medical healing
- Workshop system: Artisans maintain workshops; players pay fees to use facilities
- Skill-based progression: Crafting skills improve through use with natural difficulty curves
- No forced grinding: Skills progress at reasonable pace with meaningful use
- Source: [Discworld MUD Crafting Documentation](http://discworld.atuin.net/lpc/playing/documentation.c?path=/concepts/crafting)

**Armageddon MUD (Closed 2024, but historical lessons remain)**
- Merchant class: Dedicated crafting profession with multiple specialties
- Recipe system: Players must learn recipes (though controversial - recipes not documented in-game)
- Difficulty levels: Crafting same item with varying quality/difficulty options
- Tool requirements: Specific tools needed for different crafts
- Limitations: Hidden recipe knowledge created friction for newer players
- Source: [Armageddon MUD Crafting System](http://armageddonmud.org/help/view/Crafting)

**DartMUD (Player-run Economy)**
- Medieval production chain: Wood → Metal → Cloth → Stone → Ceramic skills
- Diversified economy: Crafting provides better rewards than thief skills
- Player-driven market: Players source raw materials and sell finished goods
- Scaling complexity: Comprehensive economic model with multiple production tiers
- Source: [DartMUD Economy Design](https://mud.fandom.com/wiki/DartMUD)

### Lesser-Successful Models

**High-Grind Crafting (CLOK example)**
- Problem: Crafting skills described as "very grindy" by the community
- Issue: Progression felt tedious without visible rewards
- Lesson: Grinding for hours without progress discourages continued engagement

**Hidden Recipes Pattern (Armageddon)**
- Problem: Recipes not available in-game or documentation
- Impact: New players forced to ask other players (creates gating friction)
- Better approach: Make recipes discoverable through gameplay or documentation

---

## 2. Common Pitfalls & Balance Issues

### The Core Economy Problem: Permanent Items

**The Central Challenge:**
If crafted items don't break/decay, players eventually have "enough" gear and stop buying. When demand drops but supply is unlimited (crafters can make infinite items), the economy stagnates.

- **DartMUD solution**: Attempted comprehensive medieval model with multiple material types to sustain crafting demand
- **WoW/FFXIV approach**: Gear tiers invalidate old gear; new expansion forces new crafting demand
- **MUD-specific**: Most text-based games lack expansion cycles; need other solutions

**Research finding:** From [MMORPG Economy Analysis](https://www.terminally-incoherent.com/blog/2010/07/08/mmo-crafting-economies/):
> "The real reason why economies are hard to balance is that things don't decay online. In the real world, everything decays, breaks, needs maintenance."

### Exploitation Through Repetition

**Problem:** Computer systems can't judge "smart" vs "exploit" skill use.
- Players can spam the same action repeatedly to level skills artificially
- No penalty for inefficient or redundant crafting

**Solutions found in MUDs:**
1. **Concentration system** (some MUDs): Each skill use reduces success chance for subsequent uses, preventing rapid-fire farming
2. **Failure-based forgetting** (proposed): Failed crafts have chance to reduce skill, discouraging lazy spamming
3. **Diminishing returns**: Crafting same item repeatedly yields less XP

### Jack-of-All-Trades Syndrome

**Problem:** Without limits, players eventually learn ALL crafting skills at max level.
- Removes specialization and interdependence
- Eliminates player trading/cooperation

**Solutions:**
1. **Skill cap per character** (e.g., max 3 crafting specialties)
2. **Time investment gates** (learning new craft resets existing skill)
3. **Resource constraints** (rare materials limited availability)

---

## 3. Power Level Balance: Crafted vs. Dropped vs. Quest Items

### Recommended Power Distribution

Research from [Pantheon Forums](https://seforums.pantheonmmo.com/content/forums/topic/1550/armor-quested-vs-crafted-vs-loot) suggests:

**Tier-based approach (Most successful):**
- At each broad tier (e.g., levels 1-5, 6-10, 11-15):
  - **Dropped loot**: 30-40% of best gear (from mob kills)
  - **Crafted items**: 30-40% of best gear (from crafting)
  - **Quest rewards**: 20-30% of best gear (from quests)
  - These should be **roughly equivalent** in power, forcing choice based on availability/playstyle

**Balance principle:**
- Crafted gear power ≤ equivalent-tier dropped/quest gear
- BUT crafted items should have unique properties (e.g., customization, specific stats)
- Crafted items take **time/resources** vs. dropped items take **luck/combat**

### Anti-Pattern Warning

**Avoid:** Making crafted items strictly better than drops
- Creates mandatory crafting
- Breaks economy (everyone becomes crafter or buys)
- Unfair to combat-focused players

**Avoid:** Making crafted items significantly weaker
- Crafting becomes useless/ignored
- Wastes development effort

---

## 4. Economy Preservation: Preventing Breakdown

### The Faucet-Drain Model

From [Massively Overpowered Economy Analysis](https://massivelyop.com/2017/07/20/massively-overthinking-building-a-better-mmorpg-economy/):

**Money Flow Model:**
- **Faucets** (money enters): Quest rewards, mob drops, NPC payments
- **Drains** (money exits): NPC vendors, taxes, repair costs
- **Goal**: Keep faucets ≈ drains to prevent inflation/deflation

### Recommended Safeguards for Portalis Crafting

1. **Item Decay System**
   - Crafted items degrade with use and require repairs
   - Repair costs are money drain
   - Creates ongoing demand for crafters
   - Alternative: Items break permanently after X uses

2. **Crafting Cost Structure**
   - Material costs: Raw materials required (gathered/looted)
   - Tool costs: Wear and tear on crafting tools
   - Time cost: Crafting locks player action for duration
   - Gold cost: NPC fees for using crafting services

3. **Limited Resources**
   - Cap daily craftable quantity per player
   - Make rare materials genuinely rare
   - Require multiple specialized resources

4. **Economic Checks**
   - Monitor average item prices
   - Track crafter profit margins (should be modest)
   - Watch for inflation in crafted goods
   - Alert if player-run economy stagnates

---

## 5. Time Sinks vs. Fun Factor

### The Grind Problem

**Finding:** MUDs that made crafting "very grindy" (CLOK example) saw player complaints.
- Hours of repetitive actions without visible progress discourages continuation
- Different from healthy progression curves

### Design Principles for Engaging Crafting

**Make progression visible:**
- Show skill XP percentage toward next level
- Give feedback ("Excellent work!" vs. "Crude result")
- Unlock new recipes as skills increase

**Vary the mechanics:**
- Different crafts use different systems (not all button-mashing)
- Puzzle-like crafting (alchemy with ingredient mixing)
- Timed crafting (quality increases with patience)

**Reward experimentation:**
- Discovering recipes through trial
- Crafting "specialty" variants with different resources
- Quality tiers (crude/standard/excellent items)

**Keep time investment reasonable:**
- Crafting same item should take seconds-to-minutes, not hours
- Complex recipes can take longer, but rarer
- Skill training shouldn't require days of repetition for meaningful progress

---

## 6. Specific Recommendations for Portalis's Crafting Proposal

### System Design

Based on portalis's existing architecture (container-based recipes):

1. **Container Recipe System (Already Implemented Foundation)**
   - Recipes as item combinations in containers
   - Simple and discoverable: players learn by experimenting
   - Can scale to complexity with different recipe combinations

2. **Proposed Additions**

   **a) Recipe Discovery**
   ```
   - Make recipes learnable through NPCs (quest rewards, skill trainers)
   - Document recipes in-game (recipe books, NPCs explain)
   - Allow experimental discovery (put items in container, see what works)
   ```

   **b) Crafting Skills**
   - Tie crafting quality to player skill level
   - Example: Blacksmithing skill 1-5 determines final weapon quality
   - Failed crafts consume ingredients but produce nothing (failure penalty)
   - Successful crafts improve skill

   **c) Resource Requirements**
   - Define material tiers (copper < iron < steel)
   - Require gathering/hunting to obtain materials
   - Make rare materials actually rare
   - Cost should match item power level

   **d) Time Investment**
   - Crafting takes time (player action locked during crafting)
   - More complex recipes take longer
   - Encourages strategic crafting timing

   **e) Item Decay (Critical for Economy)**
   - Crafted items lose durability with use
   - Repair costs create ongoing demand
   - Repair costs are money drain (economy stabilizer)
   - Alternative: Items last X uses then break permanently

3. **Avoid These Pitfalls**
   - [ ] Don't make recipes hidden or impossible to discover
   - [ ] Don't allow infinite fast crafting (prevents spam exploitation)
   - [ ] Don't make all items craftable without resource limits
   - [ ] Don't ignore economy faucet/drain balance
   - [ ] Don't tie power progression solely to crafting
   - [ ] Don't allow crafters to become rich too quickly

### Phased Implementation Approach

**Phase 1: Basic Container Recipes** (Already implemented)
- Simple item combinations in containers
- Manual discovery through experimentation
- No skill progression yet

**Phase 2: Crafting Skills & Quality**
- Add crafting skill progression
- Craft quality depends on crafter skill
- Recipe discovery through NPCs/books

**Phase 3: Material Requirements & Gathering**
- Define material sources (mobs, gathering locations)
- Implement gathering skills or merchant system
- Balance material availability

**Phase 4: Item Durability & Repair**
- Add durability system
- Implement repair mechanics and costs
- Create money drain for economy balance

### Testing & Monitoring

**Before Launch:**
1. Playtest crafting with new players
   - Can they discover recipes intuitively?
   - Does progression feel rewarding?
   - Is time investment reasonable?

2. Economy simulation
   - How quickly can new characters obtain crafted gear?
   - Does a dedicated crafter become wealthy too quickly?
   - Are there incentives for non-crafters?

3. Balance checks
   - Compare crafted item power to drops/quests
   - Verify resource costs match item value
   - Test that failures don't feel punishing

**Post-Launch Monitoring:**
- Track which recipes are used most
- Monitor crafter wealth accumulation
- Watch for "everyone is crafting" vs. "nobody crafts"
- Adjust material rarity/costs based on player behavior

---

## 7. Key Sources & References

1. [Discworld MUD Crafting System](http://discworld.atuin.net/lpc/playing/documentation.c?path=/concepts/crafting) - Active, well-designed system
2. [Armageddon MUD Documentation](http://armageddonmud.org/help/view/Crafting) - Merchant class crafting
3. [DartMUD Economy Design](https://mud.fandom.com/wiki/DartMUD) - Comprehensive medieval economy model
4. [MMORPG Economy Analysis](https://www.terminally-incoherent.com/blog/2010/07/08/mmo-crafting-economies/) - Item decay problem
5. [Pantheon Forums - Item Balance](https://seforums.pantheonmmo.com/content/forums/topic/1550/armor-quested-vs-crafted-vs-loot) - Power tier distribution
6. [Massively Overpowered - Economy Building](https://massivelyop.com/2017/07/20/massively-overthinking-building-a-better-mmorpg-economy/) - Faucet-drain model
7. [Top Mud Sites - Crafting System Discussion](https://www.topmudsites.com/forums/showthread.php?t=6967) - Community perspectives

---

## 8. Summary: The Crafting Balance Act

**The goal of a good MUD crafting system:**
- Useful but not mandatory
- Rewarding without being broken
- Engaging without being tedious
- Economically sustainable

**Key success factors:**
1. ✓ Discoverable recipes (not hidden)
2. ✓ Meaningful progression (skills improve visibly)
3. ✓ Balanced power (crafted ≈ dropped ≈ quested)
4. ✓ Resource constraints (costs prevent spam)
5. ✓ Economy safeguards (decay + drains + faucet balance)
6. ✓ Optional playstyle (crafting shouldn't be mandatory)

If portalis's crafting system addresses these points, it will avoid the pitfalls that plague other MUDs and create a sustainable, engaging system that enriches the game world without breaking it.

---

*Research compiled by Galstaff, Sorcerer of Light*
*2026-01-18*
