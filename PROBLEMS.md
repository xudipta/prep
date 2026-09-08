# Problem Index

Every problem in `dsa/*/problems/`. "Key Concept" is the one new idea the
problem teaches — if two problems would have the same answer here, the
duplicate was left out.

| # | Problem | Topic | Pattern | Difficulty | Key Concept |
|---|---------|-------|---------|------------|-------------|
| 1 | [Valid Palindrome](dsa/two-pointers/problems/valid-palindrome/README.md) | String | Two Pointers | Easy | Converging pointers skipping non-alphanumeric chars |
| 2 | [Container With Most Water](dsa/two-pointers/problems/container-with-most-water/README.md) | Array | Two Pointers | Medium | Greedy pointer move: always move the shorter side |
| 3 | [3Sum](dsa/two-pointers/problems/three-sum/README.md) | Array | Two Pointers | Medium | Fix one element + two pointers on the rest, with dedup |
| 4 | [Max Sum Subarray of Size K](dsa/sliding-window/problems/max-sum-subarray-size-k/README.md) | Array | Sliding Window (fixed) | Easy | Fixed-size window: add right, remove left |
| 5 | [Longest Substring Without Repeating Characters](dsa/sliding-window/problems/longest-substring-without-repeating/README.md) | String | Sliding Window (variable) | Medium | Shrink window on constraint violation using a seen-set/map |
| 6 | [Minimum Window Substring](dsa/sliding-window/problems/minimum-window-substring/README.md) | String | Sliding Window (variable) | Hard | Expand to satisfy, then shrink to minimize, with a "have vs need" counter |
| 7 | [Two Sum](dsa/hashing/problems/two-sum/README.md) | Array | Hashing | Easy | Complement lookup trades O(n²) for O(n) with a hash map |
| 8 | [Group Anagrams](dsa/hashing/problems/group-anagrams/README.md) | String | Hashing | Medium | Canonical key (sorted string) groups equivalent items |
| 9 | [Longest Consecutive Sequence](dsa/hashing/problems/longest-consecutive-sequence/README.md) | Array | Hashing | Medium | O(n) via set lookups by only starting counts at sequence heads |
| 10 | [Binary Search (Basic)](dsa/binary-search/problems/binary-search-basic/README.md) | Array | Binary Search | Easy | The canonical lo/hi/mid template and its invariants |
| 11 | [Search in Rotated Sorted Array](dsa/binary-search/problems/search-in-rotated-sorted-array/README.md) | Array | Binary Search on rotated space | Medium | Identify which half is sorted, then decide which half to keep |
| 12 | [Koko Eating Bananas](dsa/binary-search/problems/koko-eating-bananas/README.md) | Array | Binary Search on the answer | Medium | Binary search over a monotonic feasibility function, not the array |
| 13 | [Climbing Stairs](dsa/dynamic-programming/problems/climbing-stairs/README.md) | 1D DP | Dynamic Programming | Easy | Fibonacci-shaped recurrence from counting distinct paths |
| 14 | [House Robber](dsa/dynamic-programming/problems/house-robber/README.md) | 1D DP | Dynamic Programming | Medium | "Take or skip" state transition with a non-adjacency constraint |
| 15 | [0/1 Knapsack](dsa/dynamic-programming/problems/0-1-knapsack/README.md) | Knapsack DP | Dynamic Programming | Medium | 2D state (item, capacity) collapsible to a 1D rolling array |
| 16 | [Number of Islands](dsa/graphs/problems/number-of-islands/README.md) | Grid/Graph | Graph traversal (BFS/DFS) | Medium | Grid-as-implicit-graph flood fill, counting connected components |
| 17 | [Course Schedule](dsa/graphs/problems/course-schedule/README.md) | Graph | Topological Sort / Cycle Detection | Medium | Kahn's algorithm: a cycle exists iff not all nodes reach in-degree 0 |
| 18 | [Binary Tree Level Order Traversal](dsa/trees/problems/binary-tree-level-order-traversal/README.md) | Tree | BFS | Easy | Level-by-level BFS using the current queue length as a level boundary |
| 19 | [Lowest Common Ancestor of a BST](dsa/trees/problems/lowest-common-ancestor-bst/README.md) | Tree | BST property | Easy | Use BST ordering to decide direction instead of searching both subtrees |
| 20 | [Subsets](dsa/backtracking/problems/subsets/README.md) | Array | Backtracking | Medium | Include/exclude decision tree over every element |
| 21 | [N-Queens](dsa/backtracking/problems/n-queens/README.md) | Backtracking | Backtracking + pruning | Hard | Constraint tracking (columns/diagonals) to prune the search space early |
| 22 | [Subarray Sum Equals K](dsa/prefix-sum/problems/subarray-sum-equals-k/README.md) | Array | Prefix Sum + Hashing | Medium | Complement lookup on running prefix sums, seeded with `{0:1}` |
| 23 | [Range Sum Query - Immutable](dsa/prefix-sum/problems/range-sum-query-immutable/README.md) | Array | Prefix Sum | Easy | O(1) range queries after O(n) precomputation |
| 24 | [Jump Game](dsa/greedy/problems/jump-game/README.md) | Array | Greedy | Medium | Running "furthest reachable" bound, extended greedily |
| 25 | [Gas Station](dsa/greedy/problems/gas-station/README.md) | Array | Greedy | Medium | Running deficit resets the candidate start in one pass |
| 26 | [Valid Parentheses](dsa/stack/problems/valid-parentheses/README.md) | String | Stack | Easy | LIFO matching of nested opener/closer pairs |
| 27 | [Daily Temperatures](dsa/stack/problems/daily-temperatures/README.md) | Array | Monotonic Stack | Medium | Next-greater-element via a decreasing stack of indices |
| 28 | [Reverse Linked List](dsa/linked-list/problems/reverse-linked-list/README.md) | Linked List | Linked List | Easy | In-place pointer reversal with `prev/curr/next` |
| 29 | [Merge Two Sorted Lists](dsa/linked-list/problems/merge-two-sorted-lists/README.md) | Linked List | Linked List | Easy | Dummy head + pointer-per-list merge |
| 30 | [Single Number](dsa/bit-manipulation/problems/single-number/README.md) | Array | Bit Manipulation | Easy | XOR self-cancellation finds the unpaired element |
| 31 | [Counting Bits](dsa/bit-manipulation/problems/counting-bits/README.md) | Bit Manipulation | Bit Manipulation + DP | Easy | `n & (n-1)` recurrence reuses a smaller already-computed popcount |
| 32 | [Redundant Connection](dsa/union-find/problems/redundant-connection/README.md) | Graph | Union-Find | Medium | First edge whose endpoints are already connected creates the cycle |
| 33 | [Number of Provinces](dsa/union-find/problems/number-of-provinces/README.md) | Graph | Union-Find | Medium | Component count decreases by one on each successful union |
| 34 | [Network Delay Time](dsa/graphs/problems/network-delay-time/README.md) | Graph | Dijkstra | Medium | Min-heap finalizes the closest unvisited node first |
| 35 | [Longest Common Subsequence](dsa/dynamic-programming/problems/longest-common-subsequence/README.md) | String DP | Dynamic Programming | Medium | 2D DP comparing two sequences position by position |
| 36 | [Longest Increasing Subsequence](dsa/dynamic-programming/problems/longest-increasing-subsequence/README.md) | Sequence DP | DP + Binary Search | Medium | Patience sorting: track the smallest tail per achievable length |

More problems are added incrementally — see `PROGRESS.md` for what's planned
next. Every addition must answer "what new idea does this teach?" before it's
accepted.
