
// Electrocuted - periodic lightning damage from the construction robot

// Invoked when the buff is first applied to the player.
function onStart(actor, triggersLeft) {
    SendUserMessage(actor.UserId(), '<ansi fg="yellow">Lightning arcs through your body!</ansi>');
    SendRoomMessage(actor.GetRoomId(), '<ansi fg="yellow">Lightning arcs through ' + actor.GetCharacterName(true) + '!</ansi>', actor.UserId());
}

// Invoked every time the buff is triggered
function onTrigger(actor, triggersLeft) {
    dmgAmt = Math.abs(actor.AddHealth(-1 * UtilDiceRoll(2, 8)));

    SendUserMessage(actor.UserId(), '<ansi fg="yellow">ZZZZAP!</ansi> Electricity surges through you for <ansi fg="damage">' + String(dmgAmt) + ' damage</ansi>!');
    SendRoomMessage(actor.GetRoomId(), '<ansi fg="yellow">ZZZZAP!</ansi> Sparks fly from ' + actor.GetCharacterName(true) + '.', actor.UserId());
}

// Invoked when the buff has run its course.
function onEnd(actor, triggersLeft) {
    SendUserMessage(actor.UserId(), 'The electrical current finally dissipates from your body.');
    SendRoomMessage(actor.GetRoomId(), 'The sparks around ' + actor.GetCharacterName(true) + ' fade away.', actor.UserId());
}
