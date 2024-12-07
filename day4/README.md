# Day 4: Read your prompts

I wrote a working algorithm to recursively find every version of `XMAS` **IF you were allowed to change directions each time**. So I misread the prompt. I thought this was valid:
```text
X - - -
- M - -
- - A -
- S - -
```

This meant checking every direction for a valid next letter, every single time. Which (a) was a natural fit for recursion, and (b) generated way too many matches, and (c) literally isn't the problem, lol.

Takeaways today:

* Read your problem thoroughly. Read examples.
* If you're ever nesting loops very deeply, or reaching for loop labels, make a function instead.
* It's okay to be a bit redundant for the sake of simplicity. I over-check both part 1 and part 2, but being more efficient (fewer loops) would require managing more state for part 2. For part 1, it'd just be an early-exit, which isn't at all a big deal and is totally a low-hanging fruit win.