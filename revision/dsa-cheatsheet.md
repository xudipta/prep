# DSA Cheatsheet

One screen, top-level. See `PATTERN-MAP.md` for the full decision framework
and each `dsa/<technique>/notes.md` for technique-specific detail.

| Technique | Core idea | Complexity win |
|---|---|---|
| Two Pointers | Converge/scan two indices exploiting sorted order | O(n²) → O(n) |
| Sliding Window | Incrementally update a contiguous range's state | O(n·k) → O(n) |
| Hashing | O(1) average lookup/membership via map/set | O(n²) → O(n) |
| Binary Search | Halve a monotonic search space | O(n) → O(log n) |
| Backtracking | Choose → recurse → undo, with pruning | Exhaustive, pruned |
| Dynamic Programming | Cache/tabulate overlapping subproblems | O(2ⁿ) → O(poly) |
| Graphs (BFS/DFS) | Traverse explicit or implicit graphs | O(V+E) |
| Trees | Recursive traversal (pre/in/post/level order) | O(n) |

**Before reaching for a technique**, ask:
- Sorted/sortable + pair/triple problem? → Two Pointers.
- Contiguous subarray/substring + incremental update? → Sliding Window.
- Need O(1) lookup/membership/grouping? → Hashing.
- Monotonic predicate over a value space? → Binary Search (possibly on the
  answer, not the array).
- Need all valid configurations, prunable? → Backtracking.
- Optimal value, overlapping subproblems? → Dynamic Programming.
- Implicit or explicit graph, connectivity/shortest-path/ordering? →
  Graphs.

**Complexity vs. input size**: n≤12 → exponential OK; n≤1,000 → O(n²)/O(n³)
OK; n≤10⁶ → O(n log n); n≤10⁸ → O(n) required.

See `revision/algorithms-cheatsheet.md` and
`revision/data-structures-cheatsheet.md` for adjacent quick references, and
`tips-tricks/README.md` for the full problem-solving checklist.
