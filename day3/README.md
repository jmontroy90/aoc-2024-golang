# Day 3: Sometimes subproblems sustain state

Okay, not too bad here. My first pass has a lot of state management and object conversion that should be placed on methods, but it works. I'm not gonna write table tests for this, since it wasn't really helpful or needed, but I will take the chance to write some benchmarks and golden-file testing. So - still to do:

* Convert implied object conversions to actual methods on structs.
* Write some benchmark tests, just to get the practice.
* Write some Goldie tests.
* Start accepting input files per day (defaults being mine).

One fun lesson from today - I tried to do part 2 line by line, but since the state of a line ("still multiplying") carries over to the next, it was easier just to concatenate all the line together rather than pass state between iterations. This is a fun inversion on day 2's problem, where each line was an independent problem ("report").
