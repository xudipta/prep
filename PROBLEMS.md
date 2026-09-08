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

More problems are added incrementally — see `PROGRESS.md` for what's planned
next. Every addition must answer "what new idea does this teach?" before it's
accepted.
