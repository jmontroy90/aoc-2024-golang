# Day 12: Complexity is raised

This one felt like a step up complexity-wise - I'm late to completing it in part because of prior commitments, but also because of said complexity. Completions fell off by like a third between day 11 and day 12, and most posts on the AOC subreddit were memes about people failing to complete it.

That said, the main hurdle was figuring out the right data structures to solve this, and how to tackle the iterations. I spent a lot longer just thinking on this one, rather than just getting straight to coding with a solution and then struggling with indexes or something.

Ultimately the algorithm is simple enough - I think it's probably some kind of "flood fill" scanning algorithm, though I don't know the exact name. It could be done concurrently with how visited locations are stored in sets, but just recursively you still have to make sure not to double-count visited locations.

The algorithm, in short:

1) Find an unscanned rune.
2) Flood-fill scan that rune's area by recursively trying every direction at every location and waiting for a change before stopping.
3) If we find some place that we haven't been to that's also not our region, it must be something we want to fence, so add a directional fence.

Then we count how many unique positions and fences we have, and multiply! Part 2 is the same, but we actually use the fence data structure to find a fence, scan where it might continue, deleting as we go, and then return 1 for a side when that side's fences are all deleted (so we don't rescan).

Overall, I'm happy with how my code is reading. It's not competitive code, but it's something I'd be happy to present to anyone.