# Day 13: MATH!!!

(side note: I've had a lot going on the last few days, and so am WAY behind. We'll see if I can catch up!)

Okay, I don't feel bad about failing at part 2 of this, because answer involved working knowledge of basic linear algebra. Which I most certainly no longer have.

My thought process before looking up the answer (this [Youtube video here](https://www.youtube.com/watch?v=-5J-DAsWuJc) was really excellent, bravo) was as follows:

* I see I can no longer brute-force this. I already knew two loops was really dumb.
* Can I rewrite this with one loop and some math? Yes! `FindCheapestWithMath` shows how. Let's say this solution is canon for me in part 1.
* Does one loop help for our scale? Nope, not at all.
* Can I subdivide the problem by solving for 10e+13 first and then re-using the original solution? Nope, not how numbers work.
* What about solving for the 10e+13 number first by dividing it by 10e+12, solving for 1000, and then multiplying the pushes by 10e+12? Nope, also not how numbers work.
* Well then, this solution is either:
  * Some math I don't know that makes things trivial. (<-ding ding ding!)
  * Some divide-and-conquer algorithm for pruning branches of inputs to test, that aggressively refines itself down to a workable input. This seemed intriguing, but frankly, beyond my immediate skillset or patience to conceive of.

So after all that thinkin' and thonkin', I felt fine looking up the answer. And lo, it was math. That's fine.