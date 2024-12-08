# Day 5: Squinting at indexes

I feel like my approach here is pretty ugly and loopy, and so would be interested in other approaches. The basic idea for re-ordering (and also checking order):

1) The ordering rules are stored in a `map[int][]int`, where each value (slice) shows all the elements that must appear in our list **BEFORE** the key.
2) So for each element in our list, we look it up in the ordering rules, and for each "must appear before" (MAB) element in the rule, we scan all elements after our key element to see any violating MAB elements.
3) If we do find a MAB element, we insert it right before our key element, and remove it from its original location.
4) Each sorting over our input list MIGHT introduce other sorting issues, since it's a greedy, local algorithm. So we attempt to reorder until all rules pass. Usually this is 1 - 4 passes.

Key takeways:

* Indexes are a pain. Debuggers are great.
* I thought about doing a swap of the two elements for part (2), but my gut says this would wreck more havoc per sort and thus require more sorting iterations. Untested hypothesis. 
* Thank god for the generic `slices` package. Dunno what I'd do without it.
* There's redundant logic that can be cleaned up here. Many layers of looping to unpeel. If this were for a job, I'd spend probably another 2-4 hours refactoring, clarifying, and testing my existing solution. I might still come back and do that.

I'll update this as I learn about other approaches to this problem.
