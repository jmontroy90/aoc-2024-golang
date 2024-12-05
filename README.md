# AOC 2024 Golang

* Solve some AOC in Golang.
* Finish it for once!
* Tie it to algorithms and other resources.

# Progress

|                     Day                      | Part 1 | Part 2 |                Comments                |
|:--------------------------------------------:|:------:|:------:|:--------------------------------------:|
| [Day 1](https://adventofcode.com/2024/day/1) |   ⭐    |   ⭐    |            [Easy!](#Day 1)             |
| [Day 2](https://adventofcode.com/2024/day/2) |   ⭐    |   ⭐    | [Problems should be composed.](#Day 2) |


# Comments

### Day 1: Easy!

Overall, an easy start. I spent more time setting up the repo structure and the file loader than anything else. The problem 
itself was really straight-forward - sort and compare. The second part could be done with sorted lists, but making a map and 
looking up its counts reads cleaner to me.

### Day 2: Problems should be composed.

Well, that's more like it! Part 1 was pretty easy, but part 2 messed with me. Here are some notes:

* I tried to re-program part 2 from scratch by just taking out the element at the index it occurred. But this doesn't work, because you need to try taking out each of the three relevant elements - any one of their removals might pop things into working condition.
* Good, clean code mean finding nice boundaries; here, the part 1 boundary clearly was testing each report. But I modified this boundary to return the trouble index, which is required for part 2 to be clean and useful itself. I tried to think of a way to compose the two requirements (just returning "is report safe" bool + returning where the problem was), but it wasn't obvious. The cheap-y way to do it is to just create a `SafeReportResult` struct and return all relevant data that way - pay me big bucks please!
* Once I started using `IsReportSafe` directly, Part 2's problem become what it should be - problem index removal. Compose your problems!!
* This whole thing was also a nice exercise in recalling slice behavior. I totally botched my first attempting at using `append` because appending to a slice modifies its backing array at the slice point indicated. Total nonsense result. Using `slices.Concat` was nicer, but does too much allocation. Using `copy` - a rarity for me! - on a preallocated slice was nice to just pop the elements into where they're needed.
* I looked at other AOC2024 Go solutions to figure out whether I was undercounting or overcounting once my first attempt was wrong. To be precise, I ran my input through their programs and then figured out which reports we disagreed on. This was really useful to pop into my table tests and start to realize how I'd misconceived the solution.

Overall, this took me a little while longer than expected, but that's fair - we're getting back into it!