
// Belligerent Construction Robot
// A challenging mid-game mob that guards blocked construction zones
// Special abilities: Lightning strikes, fire effects, electromagnetic push
// Weakness: Water/ice effects neutralize the push (cancel-on-water flag on buff 40)

// Track which direction to push players (set when spawned in specific rooms)
var pushDirection = {
    296: 295,  // Room 296 pushes to 295 (northwest)
    299: 298,  // Room 299 pushes to 298 (west)
    217: 216,  // Room 217 pushes to 216 (north)
    202: 201   // Room 202 pushes to 201 (southwest)
};

// Track combat round for ability rotation
var combatRound = 0;

function onIdle(mob, room) {
    var random = Math.floor(Math.random() * 10);

    switch (random) {
        case 0:
            mob.Command("emote 's optical sensor sweeps the area with a beam of red light.");
            return true;
        case 1:
            mob.Command("say ATTENTION: THIS AREA IS UNDER CONSTRUCTION. TURN BACK IMMEDIATELY.");
            return true;
        case 2:
            mob.Command("emote adjusts its stance with a grinding of gears, blocking the path ahead.");
            return true;
        case 3:
            mob.Command("emote 's warning lights intensify as it detects movement.");
            return true;
        case 4:
            room.SendText("A <ansi fg=\"yellow\">crackle of electricity</ansi> arcs between the robot's tesla coil fingers.");
            return true;
        default:
            return false;
    }
}

function onCombatStart(mob, room, eventDetails) {
    combatRound = 0;
    mob.Command("say HOSTILE ACTION DETECTED. INITIATING REMOVAL PROTOCOL.");
    room.SendText("<ansi fg=\"red\">The construction robot's eye blazes bright red as combat systems engage!</ansi>");
}

function onCombatRound(mob, room, eventDetails) {
    combatRound++;

    // Get current target
    var target = mob.GetTarget();
    if (target == null) {
        return false;
    }

    // Ability rotation based on combat round
    var abilityRoll = UtilDiceRoll(1, 100);

    // Every 4 rounds: Electromagnetic push (can be canceled by water/ice)
    if (combatRound % 4 == 0) {
        mob.Command("emote 's chest plates slide open, revealing a pulsing electromagnetic core!");
        room.SendText("<ansi fg=\"cyan\">VRRRRRMMMMM!</ansi> The robot charges its electromagnetic repulsor!");

        // Check if target has water/ice protection (HasBuffFlag checks for cancel-on-water)
        // Water spells like Drench (buff 23) or ice effects will have this flag
        var hasWaterProtection = target.HasBuffFlag("cancel-on-water");

        if (!hasWaterProtection) {
            // Apply the EM repulsion buff
            target.GiveBuff(40, "construction_robot");

            // Push the player out of the room
            var currentRoom = mob.GetRoomId();
            var targetRoom = pushDirection[currentRoom];

            if (targetRoom) {
                target.SendText("<ansi fg=\"cyan\">The electromagnetic pulse hurls you backward!</ansi>");
                room.SendText("<ansi fg=\"cyan\">" + target.GetCharacterName(true) + " is blasted away by the electromagnetic pulse!</ansi>", target.UserId());
                target.MoveRoom(targetRoom);
                mob.Command("say INTRUDER REMOVED. RESUMING PATROL.");
            }
        } else {
            // Water/ice effect protected them - robot short circuits!
            room.SendText("<ansi fg=\"cyan\">SPUTTER! CRACKLE!</ansi> The electromagnetic pulse fizzles ineffectively!");
            room.SendText("<ansi fg=\"blue\">The water/ice on " + target.GetCharacterName(true) + " short-circuits the robot's systems!</ansi>");
            mob.Command("emote 's systems spark and short-circuit momentarily!");
            // Robot takes some feedback damage
            mob.AddHealth(-1 * UtilDiceRoll(2, 6));
        }
        return true;
    }

    // Every 3 rounds: Lightning strike (electrocuted debuff)
    if (combatRound % 3 == 0 && abilityRoll <= 70) {
        mob.Command("emote extends its tesla coil arm, electricity building to a blinding crescendo!");
        target.GiveBuff(41, "construction_robot");

        // Immediate lightning damage
        var dmg = UtilDiceRoll(3, 6);
        target.AddHealth(-1 * dmg);
        target.SendText("<ansi fg=\"yellow\">CRACKOOM!</ansi> A bolt of lightning strikes you for <ansi fg=\"damage\">" + String(dmg) + " damage</ansi>!");
        room.SendText("<ansi fg=\"yellow\">CRACKOOM!</ansi> Lightning arcs from the robot into " + target.GetCharacterName(true) + "!", target.UserId());
        return true;
    }

    // 25% chance: Steam vent (fire damage to everyone in room)
    if (abilityRoll <= 25) {
        mob.Command("emote 's pressure valves release with a deafening HISSSSS!");
        room.SendText("<ansi fg=\"red\">Superheated steam billows from the robot in all directions!</ansi>");

        // Apply on_fire buff (buff 22 - already exists in GoMud)
        target.GiveBuff(22, "construction_robot");
        return true;
    }

    // Otherwise, intimidating combat flavor
    var flavorRoll = UtilDiceRoll(1, 4);
    switch (flavorRoll) {
        case 1:
            mob.Command("emote 's pile-driver arm retracts with a hydraulic hiss!");
            break;
        case 2:
            mob.Command("emote grinds forward on treaded feet, metal screeching against stone!");
            break;
        case 3:
            room.SendText("The robot's speaker crackles: <ansi fg=\"red\">\"RESISTANCE IS INADVISABLE.\"</ansi>");
            break;
        case 4:
            mob.Command("emote 's chest panel glows with building heat!");
            break;
    }

    return false;
}

function onDeath(mob, room, eventDetails) {
    room.SendText("<ansi fg=\"yellow\">CRITICAL SYSTEM FAILURE</ansi>");
    room.SendText("The construction robot shudders violently, gears grinding and sparks flying!");
    room.SendText("<ansi fg=\"red\">KABOOM!</ansi> The robot explodes in a shower of cogs, springs, and molten metal!");
    room.SendText("A tinny voice crackles from the wreckage: <ansi fg=\"cyan\">\"Maintenance... required...\"</ansi>");
}
