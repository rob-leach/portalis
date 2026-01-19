
// Invoked when the buff is first applied to the player.
function onStart(actor, triggersLeft) {
    SendUserMessage(actor.UserId(), 'You crack open the acorn and munch on the nutty interior.');
}

// Invoked every time the buff is triggered (see roundinterval)
function onTrigger(actor, triggersLeft) {
    healAmt = actor.AddHealth(UtilDiceRoll(1, 3));
    SendUserMessage(actor.UserId(), 'The acorn\'s energy heals you for <ansi fg="healing">'+String(healAmt)+' damage</ansi>.');
}

// Invoked when the buff has run its course.
function onEnd(actor, triggersLeft) {
    SendUserMessage(actor.UserId(), 'You finish digesting the acorn.');
}
