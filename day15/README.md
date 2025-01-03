# Day 15: There's a data structure somewhere

Okay, yes, I see it. So much repeated logic, so much nested looping. I think I could, with a couple hours, get pretty far in cleaning this up. That's maybe one round of work. More interesting would be a second round where I study solutions from others and they probably just use a friggin' deque or something and call it a day for things. I'm also interested in a recursive solution - part 2 is very BFS-y. I need to develop more of a familiarity with "algorithmic situations", e.g. pattern matching a problem to an existing solution (without locking myself in!).

But I'm glad I got this to work. Let's call a todo here:

* Consolidate massively repeated logic.
* Make some kind of generic "Moveable" type that encompasses both `O` and `[]`.
* Look up existing solutions and follow whatever algorithms they do.

## Notes from Others
* There's no need to have different logic in part 2 for the `left / right` vs. `up / down` cases - it just was easier to think about.
* HyperNeutrino does de-dupe targets for part 2, but mostly to avoid an exponential size blow-up.
* You don't need to de-dupe targets to avoid bad moves if you make a copy of the unmodified grid each time, and take the move inputs from that unmodified grid. But this does, well, make a full grid copy every move. Not awful if you replace it in place.
* I clearly can consolidate so much logic.