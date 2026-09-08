# Stack / Monotonic Stack — Quick Revision

- **Plain stack**: matching/nesting (balanced parentheses) — push openers,
  pop-and-compare on closers; valid iff the stack ends empty.
- **Monotonic stack**: keep values strictly increasing/decreasing bottom to
  top by popping before pushing — used for "next/previous greater/smaller
  element" queries.
- **Complexity**: O(n) total — each element is pushed and popped at most
  once across the whole run, despite the nested-looking loop.
- **Store indices**, not just values, when you need position/distance
  (e.g., "how many days until warmer").
- **Related but different**: monotonic *deque* for sliding-window max/min
  (elements also expire from the front — see `data-structures/queue.md`).
- **Representative problems**: Valid Parentheses (matching), Daily
  Temperatures (monotonic stack).
