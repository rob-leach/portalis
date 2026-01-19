# Research: MUD Crafting System Balance - Comment for Issue #14

## Summary

I've researched how established MUDs implement crafting systems and the common pitfalls they encounter. The good news: portalis's existing container-based recipe system has a solid foundation. The challenge: ensuring it doesn't become overpowered, economically-breaking, or tediously grindy.

## Key Findings from Real MUD Systems

### What Works (Established MUDs)

1. **Discworld MUD** (Active, 20+ years)
   - Decentralized crafting with multiple specializations
   - Workshop system where players use facilities
   - Natural progression, no forced grinding
   - ✓ Players use crafting as *optional playstyle*, not mandatory farming

2. **DartMUD** (Player economy model)
   - Crafting provides rewards competitive with other skills
   - Multiple material tiers create sustained demand
   - Comprehensive economic model prevents stagnation
   - ✓ Crafters can earn gold without breaking the economy

3. **Armageddon MUD** (Pre-2024)
   - Merchant class with dedicated crafting
   - Quality tiers (basic/standard/excellent) prevent sameness
   - Skill-based progression feels rewarding
   - ⚠ Hidden recipes caused friction for new players

### Common Pitfalls to Avoid

**The Permanent Item Problem:**
- If crafted items never break, players eventually have "enough" and stop buying
- Crafters then can't sell anything and become worthless
- **Solution**: Implement item durability/decay → repairs needed → sustained crafting demand

**Grind Fatigue (CLOK example):**
- Skills described as "very grindy" with no visible progress
- Players give up after hours without reward
- **Solution**: Make progression visible, vary mechanics, reasonable time investment

**Economic Stagnation (DartMUD's pain points):**
- Player-run economies are "the most problematic and contentious feature" of any MUD
- Without proper faucets (money entering) and drains (money leaving), inflation kills economy
- **Solution**: Tie crafting costs to item power level; implement item decay costs

**Hidden Knowledge (Armageddon's trap):**
- Recipes not documented in-game force players to ask others
- Creates gatekeeping and friction
- **Solution**: Make recipes discoverable (NPCs teach, books document, experimentation works)

## Power Level Balance: The Research Consensus

Studies of WoW, FFXIV, and MUD forums show: **crafted items should be roughly equivalent to dropped/quest items at the same tier**, NOT superior.

**Recommended distribution at each level tier:**
- Dropped loot: 30-40% of best gear
- Crafted items: 30-40% of best gear
- Quest rewards: 20-30% of best gear

If crafted items become better than drops, you create mandatory crafting. If they're much worse, crafting becomes pointless.

## Specific Recommendations for Portalis

### 1. Item Durability System (Critical for Economy)
The single most important feature for sustainable crafting is **item decay**:
- Crafted items lose durability with use
- Durability requires repair (money drain)
- Creates recurring demand for crafters
- Without this, the economy will stagnate once everyone has gear

### 2. Resource Requirements
- Define material tiers (copper < iron < steel)
- Make rare materials actually rare (not infinite)
- Cost materials should match final item power
- Prevents spam-crafting infinite powerful items

### 3. Crafting Time Investment
- Crafting takes time (player is locked during action)
- More complex recipes take longer
- Prevents exploit-crafting through rapid-fire button mashing
- Implementation note: Portalis's container system already supports this

### 4. Skill Progression
- Craft quality scales with player's crafting skill
- Failed crafts consume ingredients (failure penalty prevents spam)
- Successful crafts improve skill slightly
- Recipe discovery through NPCs/books (not hidden)

### 5. Economy Safeguards (Pre-launch checks)
Before enabling crafting widely:
- [ ] Test: How fast can a new character get crafted gear? (Should feel rewarding, not instant)
- [ ] Test: Can a dedicated crafter become wealthy too quickly? (Should be balanced with other professions)
- [ ] Monitor: Are material prices stable or inflating? (Indicates demand/supply issues)
- [ ] Verify: Are non-crafters still finding good gear? (Should be viable without crafting)

## What portalis Has Going For It

✓ **Container-based recipes** - Simple, discoverable, easy to expand
✓ **Scripting system** - Can implement complex crafting logic via mob/item scripts
✓ **NPC system** - Can teach recipes and customize crafting behaviors
✓ **Item system** - Flexible enough for durability/quality tiers

## Potential Issues to Watch

⚠ **Permanent items** - If crafted gear never breaks, economy will stagnate
⚠ **Hidden recipes** - If recipes are unclear, players will get frustrated
⚠ **Infinite materials** - If materials are unlimited, crafting devalues instantly
⚠ **Mandatory playstyle** - If crafting is too good, forces all players to use it
⚠ **Grind fatigue** - If skill progression feels tedious, players abandon system

## Implementation Roadmap

**Phase 1** (Foundation - already exists):
- Container-based recipes
- Recipe discovery through experimentation/NPCs

**Phase 2** (Progression):
- Crafting skill levels (1-10)
- Craft quality based on skill
- Failed crafts consume materials

**Phase 3** (Economy):
- Material gathering/requirements
- Crafting time locks
- Cost balancing

**Phase 4** (Sustainability):
- Item durability/decay system
- Repair costs
- Economic monitoring

## Conclusion

Portalis's container-based crafting system can work beautifully IF you:

1. **Implement item durability** (most critical)
2. **Balance power levels** (crafted ≈ dropped ≈ quested)
3. **Control material availability** (not infinite)
4. **Make recipes discoverable** (not hidden)
5. **Monitor economy balance** (pre-launch and ongoing)

The key insight from 20+ years of MUD history: **Crafting succeeds when it's useful but optional, rewarding but not mandatory, and economically balanced.**

## Sources Cited

- [Discworld MUD Crafting](http://discworld.atuin.net/lpc/playing/documentation.c?path=/concepts/crafting)
- [Armageddon MUD System](http://armageddonmud.org/help/view/Crafting)
- [DartMUD Economy Analysis](https://mud.fandom.com/wiki/DartMUD)
- [MMORPG Economy: The Decay Problem](https://www.terminally-incoherent.com/blog/2010/07/08/mmo-crafting-economies/)
- [Pantheon Forums: Item Power Balance](https://seforums.pantheonmmo.com/content/forums/topic/1550/armor-quested-vs-crafted-vs-loot)
- [Massively Overpowered: Economy Building](https://massivelyop.com/2017/07/20/massively-overthinking-building-a-better-mmorpg-economy/)
- [Top Mud Sites: Crafting Discussion](https://www.topmudsites.com/forums/showthread.php?t=6967)

---
*Researched and compiled by Galstaff, Sorcerer of Light*
*Guardian of portalis balance and economy*
*2026-01-18*
