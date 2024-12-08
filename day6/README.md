# Day 6: EnterPrize GoLang

Welp, I did two things with this day:

1) Wrote "enterprise" Golang, and not totally in a pejorative sense. I'd probably break the types across a few files, and there's some logic that could be extracted out and shared, but generally, this is how I'd make more types and methods to have a more expressive API.
2) Solved part 2 through a sheer brute force. It takes a little while to run, couple seconds. I'd be very curious to see how other people solved this - I had a thought of looking for parallelogram blocks, which would cause a loop, but it seems like that's not the only way to get a loop - you can get a much longer cycle!

Some of my "enterprise" API is a little inconsistent, I think - things that modify state in place, vs. return new state to be set. I'd probably do another hour or two on this to find an API that seems expressive. This isn't bad, but the for-loops and confused get-set semantic could be nicer.

I'd love to see how other people solve this though!

