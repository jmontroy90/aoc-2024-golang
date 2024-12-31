# Day 16: Hello old friend (e.g. algorithms)

Oh, what fun. We've finally hit the day where my deficits in DSA (data structures + algorithms) was an actual hindrance. This is why I'm here, right? Well, for the review of DSA stuff and the challenger, and today I got both.

So this was some kinda take on BFS / Dijkstra. Both things I know next to nothing about. I fiddled with a recursive attempt for a little bit initially, but that was misguided in two ways (impossible aggregation + recursion is only possible for BFS, not DFS). So I looked into BFS and learned how to do it in Golang, taking HEAVILY from existing solutions. Then the part 1 solution transforms it into Dijkstra, because our paths are weighted here. That re-introduced me to heaps, and how they're used in the priority queue that is needed for Dijkstra.

All great. And I'm proud that, although I didn't organically come up with BFS or Dijkstra (which is A-okay by me), I understood them completely by the time my part 1 answer was in. Then came part 2.

This is, I guess, also Dijkstra? I'll find out as I read more. But I'm really proud of myself here, because without any outside help, I did two things I think were crucial:

1) Implemented the algorithm for tracking paths during the search.
2) Figured out how to prune branches in this path-preserving world ("hit another path, only stay active if you have a lower or equal score").

I'm sure my code is messy as hell, and I would like to chunk it into discrete blocks. But I'll save that for another day. For now, I think I did pretty good.