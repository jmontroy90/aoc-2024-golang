# Day 2: Problems should be composed.

Well, that's more like it! Part 1 was pretty easy, but part 2 messed with me. Here are some notes:

* I tried to re-program part 2 from scratch by just taking out the element at the index the problem first occurred. But this doesn't work, because you need to try taking out each of the three relevant elements - any one of their removals might pop things into working condition.
* Good, clean code mean finding nice boundaries; here, the part 1 boundary clearly was testing each report. But I modified this boundary to return the trouble index, which is required for part 2 to be clean and useful itself. I tried to think of a way to compose the two requirements (just returning "is report safe" bool + returning where the problem was), but it wasn't obvious. The cheap-y way to do it is to just create a `SafeReportResult` struct and return all relevant data that way - pay me big bucks please!
* Once I started using `IsReportSafe` directly, Part 2's problem become what it should be - problem index removal. Compose your problems!!
* This whole thing was also a nice exercise in recalling slice behavior. I totally botched my first attempting at using `append` because appending to a slice modifies its backing array at the slice point indicated. Total nonsense result. Using `slices.Concat` was nicer, but does too much allocation. Using `copy` - a rarity for me! - on a preallocated slice was nice to just pop the elements into where they're needed.
* I looked at other AOC2024 Go solutions to figure out whether I was undercounting or overcounting once my first attempt was wrong. To be precise, I ran my input through their programs and then figured out which reports we disagreed on. This was really useful to pop into my table tests and start to realize how I'd misconceived the solution.

Overall, this took me a little while longer than expected, but that's fair - we're getting back into it!

## Notes from Others

- Creating pairs with two pointers, or an off-by-one zipping, makes it easier to iterate.
- The core data element here is the pair differences. Focus on that.
- If you have a slice of pair-wise differences, you can evaluate if all of them are negative vs. all positive.
  - This way we check all conditions every time for a global pass / fail, rather than some extra state one time.
