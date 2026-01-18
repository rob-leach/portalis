
// Electromagnetic Repulsion - pushes the player back
// Canceled by water effects (shorts out the robot's systems)

// Invoked when the buff is first applied to the player.
function onStart(actor, triggersLeft) {
    SendUserMessage(actor.UserId(), '<ansi fg="cyan">BZZZZT!</ansi> A massive electromagnetic pulse slams into you!');
    SendRoomMessage(actor.GetRoomId(), '<ansi fg="cyan">BZZZZT!</ansi> ' + actor.GetCharacterName(true) + ' is hit by an electromagnetic pulse!', actor.UserId());
}

// Invoked every time the buff is triggered
function onTrigger(actor, triggersLeft) {
    // Push is handled by the mob script via MoveRoom
}

// Invoked when the buff has run its course.
function onEnd(actor, triggersLeft) {
    SendUserMessage(actor.UserId(), 'The electromagnetic force dissipates.');
}
