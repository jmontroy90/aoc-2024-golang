# Day 3: Sometimes subproblems sustain state

Okay, not too bad here. My first pass has a lot of state management and object conversion that might be placed on methods, but there's not a nice boundary that I can immediately find. I'm not gonna write table tests for this, since it wasn't really helpful or needed, but I took the chance to write a benchmark. Simple. Side note - the `-8` or whatever you get appended on your benchmark test name is the number of CPUs used to run the benchmark. By default, it's `GOMAXPROCS`. You can try different combos on the command line via:

```shell
go test ./... -bench=. -cpu=1,2,4,8
# BenchmarkSumWithStop     	   51266	     21286 ns/op
# BenchmarkSumWithStop-2   	   62448	     19139 ns/op
# BenchmarkSumWithStop-4   	   62978	     19035 ns/op
# BenchmarkSumWithStop-8   	   62290	     19211 ns/op
```

Takeaways:

* The theme of this one was clearly regular expressions, parsing out huge chunks of text. Taking the chance to explore the regexp library a bit was useful - capture groups and NAMED capture groups are useful, although they left me craving some wrapper library for some of the bare-bones functionality of the core regexp library.
* A lot of the state transition code smells to me, it's a lot of object conversion in ways that feels inelegant. The thought has to be - is this a one-to-one object conversion? For `changePoint`, it is; for `scanRange`, it actually takes many `changePoint` objects for one `scanRange` object, since a lot of `changePoint` objects don't actually change the scan range state. An interesting pattern to capture!
* Sometimes subproblems sustain state, e.g. "don't decompose a problem if it requires global state". I tried to do part 2 line by line, but since the state of a line ("still multiplying") carries over to the next, it was easier just to concatenate all the line together rather than pass state between iterations. This is a fun inversion on day 2's problem, where each line was an independent problem ("report").
