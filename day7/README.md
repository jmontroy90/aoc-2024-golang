# Day 7: I am a brute

This one was conceptually pretty easy, although my initial attempt to do it recursively failed because I am not up to snuff on my recursion. I decided to brute-force it again (what else do people do here?), and frankly a bunch of frames on the stack is probably more wasteful than the Cartesian product I used. I openly stole the Cartesian product implementation from StackOverflow, but did my best to comment and understand it (see code).

Couple thoughts:

* Seeing this as a Cartesian product problem was clarifying - worth reviewing combinations, permutations, and Cartesian products here as a result.
* The operations are all essentially `foldLeft`, where the `concat` operation only works because the initial zero converts back to the first number once parsed, e.g. "043" -> 43.
* How do people solve this without brute force? There's some heuristics you could probably use, some early-exits and whatnot, but I'm not seeing the elegance right away. It'll be interesting to sit with these, possibly all year. 

* I wouldn't call this one hard, all things considered - I just feel like a bit of an ogre. 