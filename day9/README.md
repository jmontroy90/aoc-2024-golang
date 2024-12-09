# Day 9: I Scan, You Scan, We All Scan

It's a ScanFest! Part 2 of these was just maintaining scanning indexes from the left and right, and being careful to keep track of them properly. GoLand's debugger was very helpful here in helping me step through it and nail down bugs - you competition programmers who can just mentally visualize all this are nuts.

Main thoughts:

* I started off shafting myself by treating everything as runes. ID numbers are not runes - double-digit numbers are gonna kill ya! Always a good lesson - make your type match the actual meaning of the data. The input can be treated as runes, since it's not a number but essentially a cipher key. The second we introduce number IDs, we have to use numbers.
* Part 2's index management was fairly nuts. I lost some time not realizing that I'd have to scan for a place for every block **from the beginning**, every time. This means you end up scanning everything, right to left.
* There's definitely some shared logic here in both parts. Would be interesting to try to refactor this, but I'm done for the day. I'd try to share stuff, and probably make a `Block` type with methods if I were really doing this.
* There's also definitely some collapsible for-loop logic, instead of all those `for { ... }` fellers.

Overall, exhausting and focus-demanding, but not hard conceptually.