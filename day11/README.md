# Day 11: Just cache!

So this first part went quite well recursively - I read the directions well enough to realize immediately that red herrings like "order matters" didn't contribute to our answer, which didn't care about summing the numbers or ordering them or anything. So that was nice. Then the recursive solution was easy and I'm sure some classic algorithm.

Part 2 though - like everyone, I tried my naive implementation with a big number, and it did not want to finish. I tried a (I thought!) good optimization (maybe still is) - "if it's even, you can always degenerate it to individual digits after log2(len(num)) blinks, so just skip ahead to that". Then I tried caching in a way that maintained state on the recursive call, because I didn't see them much-more-obvious "cache results per if-branch" solution. Trying to cache results on the recursive-exit branch was hell, because every recursive call only has a bit of the state, and there's probably no way to cleanly track (or if there is, it's dumb and laborious).

Once I cached, part 2 blazed in like a second.

I'm looking forward to shoring up all this knowledge with fundamentals and other review, frankly!


