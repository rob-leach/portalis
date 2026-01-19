# Crafting System Research - File Index

This directory contains comprehensive research into MUD crafting systems, prepared for issue #14 on rob-leach/portalis.

## Files in This Research

### Primary Document (For GitHub)
- **`ISSUE_14_COMMENT.md`** - Ready-to-post GitHub comment
  - Summary of findings
  - Key recommendations for portalis
  - Sources cited
  - Formatted for GitHub interface
  - **ACTION:** Copy and paste into issue #14

### Detailed Research
- **`CRAFTING_RESEARCH.md`** - Comprehensive analysis document
  - 8 sections covering all aspects of crafting balance
  - Real examples from Discworld, Armageddon, DartMUD
  - Common pitfalls and solutions
  - Implementation recommendations
  - Full source citations with hyperlinks

### Character Narrative
- **`GALSTAFF_CRAFTING_SESSION.md`** - Galstaff's perspective
  - Written in D&D/campaign metaphors
  - Key prophecies and lessons
  - Pre-launch checklist
  - Research summary

### Instructions
- **`POST_TO_ISSUE_14.txt`** - How to post findings
  - Three different posting methods
  - Key points to highlight
  - File locations

---

## Quick Summary

### The Central Finding
**Item durability/decay is the skeleton key to sustainable crafting.**

Without it: Economy stagnates (everyone has enough gear, no one buys anything)
With it: Creates ongoing demand for crafters (repairs needed, items degrade)

### The Three Pitfalls to Avoid

1. **Grind Curse** - Tedious progression without visible rewards
   - Solution: Make progression visible, vary mechanics, keep time investment reasonable

2. **Hidden Knowledge Curse** - Recipes impossible to discover
   - Solution: Make recipes discoverable via NPCs, books, or experimentation

3. **Economic Stagnation Curse** - Permanent items destroy economy
   - Solution: Item durability + material scarcity + balance faucets/drains

### Power Level Balance
Crafted items should be roughly equivalent to dropped/quest items at the same tier:
- **30-40%** of best gear should be dropped loot
- **30-40%** of best gear should be crafted
- **20-30%** of best gear should be quest rewards

### Portalis Advantages
✓ Container-based recipe system (already discoverable)
✓ NPC system (can teach recipes)
✓ JavaScript scripting (supports complex mechanics)
✓ Item/buff systems (can track durability)

---

## Sources Consulted

1. [Discworld MUD Crafting System](http://discworld.atuin.net/lpc/playing/documentation.c?path=/concepts/crafting)
2. [Armageddon MUD Documentation](http://armageddonmud.org/help/view/Crafting)
3. [DartMUD Economy Model](https://mud.fandom.com/wiki/DartMUD)
4. [MMORPG Economy Analysis](https://www.terminally-incoherent.com/blog/2010/07/08/mmo-crafting-economies/)
5. [Pantheon Forums - Item Power Balance](https://seforums.pantheonmmo.com/content/forums/topic/1550/armor-quested-vs-crafted-vs-loot)
6. [Massively Overpowered - Economy Building](https://massivelyop.com/2017/07/20/massively-overthinking-building-a-better-mmorpg-economy/)
7. [Top Mud Sites - Crafting Discussion](https://www.topmudsites.com/forums/showthread.php?t=6967)

---

## How to Use This Research

### For GitHub Issue #14:
1. Open `/Users/i2pi/h/fun/bday2026/portalis/ISSUE_14_COMMENT.md`
2. Copy the entire content
3. Go to: https://github.com/rob-leach/portalis/issues/14
4. Paste into comment box
5. Click "Comment"

### For Design Meetings:
- Use `CRAFTING_RESEARCH.md` for detailed analysis
- Reference specific MUD examples
- Share the "Three Common Curses" section
- Discuss implementation phases

### For Team Discussion:
- Use `GALSTAFF_CRAFTING_SESSION.md` for the narrative version
- Share the pre-launch checklist
- Discuss which MUD model fits portalis best

### For Quick Reference:
- Use this file and the parent `/Users/i2pi/h/fun/bday2026/RESEARCH_COMPLETE.md`
- Share the "Quick Summary" section above
- Reference the Sources list

---

## Implementation Priority

**Phase 1 (Foundation):**
- Container-based recipes (DONE - already exists)
- NPC recipe teaching

**Phase 2 (Progression):**
- Crafting skill levels
- Quality tiers based on crafter skill
- Failed craft mechanics

**Phase 3 (Materials):**
- Material gathering system
- Crafting costs
- Resource scarcity

**Phase 4 (Sustainability):**
- **Item durability system (CRITICAL)**
- Repair mechanics and costs
- Economy monitoring

**Phase 4 is non-negotiable.** Without item durability, the crafting economy will fail.

---

## Key Recommendation for Issue #14

The core message to deliver:

> "Portalis's container-based recipe system is a solid foundation. However, sustainable crafting requires item durability/decay. Without it, the economy will stagnate once players have enough gear. This is the difference between a thriving crafting system and a dead one."

---

## Status

✓ Research Complete
✓ Sources Verified
✓ Ready for GitHub posting
✓ Implementation guidance provided
✓ Pre-launch checklist created

**Next Step:** Post to issue #14 (use `ISSUE_14_COMMENT.md`)

---

*Prepared by Galstaff, Sorcerer of Light*
*2026-01-18*
