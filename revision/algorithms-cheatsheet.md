# Algorithms Cheatsheet

## Complexity Classes (fastest to slowest)

`O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(n³) < O(2ⁿ) < O(n!)`

## Sorting

| Algorithm | Time (avg) | Time (worst) | Space | Stable? |
|---|---|---|---|---|
| Quicksort | O(n log n) | O(n²) | O(log n) | No |
| Mergesort | O(n log n) | O(n log n) | O(n) | Yes |
| Heapsort | O(n log n) | O(n log n) | O(1) | No |
| Insertion Sort | O(n²) | O(n²) | O(1) | Yes (good for nearly-sorted/small n) |
| Counting Sort | O(n+k) | O(n+k) | O(k) | Yes (k = value range, non-comparison) |

Go's `sort.Sort` uses an introsort variant (quicksort + heapsort fallback);
`sort.Stable` uses an insertion-sort/mergesort hybrid for stability.

## Searching

- **Linear search**: O(n), no ordering required.
- **Binary search**: O(log n), requires sorted/monotonic structure — see
  `dsa/binary-search`.

## Graph Algorithms — At a Glance

| Algorithm | Purpose | Complexity |
|---|---|---|
| BFS | Shortest path, unweighted | O(V+E) |
| DFS | Components, cycles, topo sort | O(V+E) |
| Dijkstra | Shortest path, non-negative weights | O((V+E) log V) |
| Bellman-Ford | Shortest path, negative weights OK | O(V·E) |
| Floyd-Warshall | All-pairs shortest path | O(V³) |
| Kruskal | MST (edge-sorted + DSU) | O(E log E) |
| Prim | MST (grow from a node + heap) | O(E log V) |

## Recursion → DP

Every DP problem starts as a recursive brute force. If the same
subproblem is computed more than once, cache it (top-down) or fill a table
in dependency order (bottom-up). See `dsa/dynamic-programming/README.md`.

## Master Theorem (Divide & Conquer)

For `T(n) = a·T(n/b) + O(nᶜ)`:
- If `c < log_b(a)`: `T(n) = O(n^log_b(a))`.
- If `c == log_b(a)`: `T(n) = O(nᶜ log n)`.
- If `c > log_b(a)`: `T(n) = O(nᶜ)`.

Example: mergesort — `T(n) = 2T(n/2) + O(n)` → `a=2, b=2, c=1`,
`log_b(a) = 1 = c` → `T(n) = O(n log n)`.

## Amortized Analysis

"Amortized O(1)" means the *average* cost per operation over a sequence is
O(1), even if occasional individual operations cost more (e.g., a slice
`append` that occasionally triggers an O(n) reallocation, but doubles
capacity each time — so reallocations become exponentially rarer as the
slice grows, averaging out to O(1) per append).

See also: `revision/data-structures-cheatsheet.md`,
`dsa/dynamic-programming/dp-pattern-recognition.md`.
