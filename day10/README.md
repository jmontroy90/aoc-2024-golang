# Day 10: Beautiful Code

Well, I finally took my `Grid` code and put it into a separate module (painfully named `util`). This is a form of premature optimization, but for learning purposes, I think it'll be interesting to try to create a clean enough API boundary such that multiple days can re-use the `util.Grid` code (because right now they all manage state slightly differently). The good news is that my `Grid` code once again made my logic quite clean and readable.

This recursive solution is, in fact, the recursive solution I thought I needed for part 1 of the `XMAS` challenge on Day 4, and is (I suspect) some kind of depth-first search (might not be kind of, might be exactly). Regardless, it was a fun exercise and since the core logic knocked down pretty quick, I was able to spend some time on structure and maintainability. This is now code that looks fairly close - not 100%! - to what I'd be happy to put into production at a job.

Takeaways:
* This is, as mentioned, probably a depth-first search, so it'll be good to take this opportunity to re-learn those algorithms and practice implementing them.
* Struct state vs. method state - is the state relevant per method call? Or do you just need over time through a whole set of calls? The `visited` map is perfect for structs, just as long as it gets reset when needed. The other option is to just carry it around your method calls - valid, but tricky when you have lots of state to potentially add over time.
* My `Grid` code now offers a `Scan()` that emulates the API of `bufio.Scanner` - I really like the `for buf.Scan() { buf.ScanText() }` pattern, I think it reads intuitively and works well with language constructs. I haven't technically tested the `ScanOnce` code yet - that'll be useful for refactors using this `Grid` code in other days.
* By the end of this, I'm gonna have a lot of code for grid management!