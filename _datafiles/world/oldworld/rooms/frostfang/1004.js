// Snarky Squirrel Commons - Time-based NPC spawning
// Day: Janitor cleans and mutters jokes
// Night: Performers and hecklers do their thing

const COMEDIAN_MOB_ID = 63;
const HECKLER_MOB_ID = 64;
const JANITOR_MOB_ID = 72;

function onIdle(room) {

    var isDay = UtilIsDay();

    // Get current mobs in room
    var comedians = room.GetMobs(COMEDIAN_MOB_ID);
    var hecklers = room.GetMobs(HECKLER_MOB_ID);
    var janitors = room.GetMobs(JANITOR_MOB_ID);

    if (isDay) {
        // Daytime - janitor works, performers leave

        // Despawn performers if present
        for (var i in comedians) {
            room.SendText("The failed comedian mutters something about needing sleep and shuffles out.");
            comedians[i].Command("despawn");
        }
        for (var i in hecklers) {
            room.SendText("The enthusiastic heckler yawns and wanders off.");
            hecklers[i].Command("despawn");
        }

        // Spawn janitor if not present
        if (janitors.length == 0) {
            room.SendText("The janitor shuffles in with his mop, ready for another day of cleaning.");
            room.SpawnMob(JANITOR_MOB_ID);
        }

    } else {
        // Nighttime - performers arrive, janitor leaves

        // Despawn janitor if present
        for (var i in janitors) {
            room.SendText("The janitor puts away his mop and heads home for the night.");
            janitors[i].Command("despawn");
        }

        // Spawn comedian if not present
        if (comedians.length == 0) {
            room.SendText("A dejected-looking performer shuffles in from backstage.");
            room.SpawnMob(COMEDIAN_MOB_ID);
        }

        // Spawn heckler if not present (slight delay feels more natural)
        if (hecklers.length == 0 && comedians.length > 0) {
            room.SendText("Someone pushes through the crowd and finds a seat near the stage.");
            room.SpawnMob(HECKLER_MOB_ID);
        }
    }

    return true;
}

function onLoad(room) {
    // On server start, spawn appropriate NPCs based on time
    onIdle(room);
}
