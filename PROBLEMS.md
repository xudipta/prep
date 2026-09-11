# Problem Index

Every problem in `dsa/*/problems/`, grouped by technique in the same order
as the sidebar, `README.md`'s learning path, and `PROGRESS.md` — read top to
bottom to move through the curriculum technique by technique. "Key Concept"
is the one new idea the problem teaches — if two problems would have the
same answer here, the duplicate was left out. "LeetCode" links straight to
the problem to try it yourself; a couple of foundational/textbook problems
aren't hosted on LeetCode, so those link to the closest official equivalent
instead (noted in the problem's own README).

| # | Problem | Topic | Pattern | Difficulty | LeetCode | Key Concept |
|---|---------|-------|---------|------------|----------|-------------|
| 1 | [Two Sum](dsa/hashing/problems/two-sum/README.md) | Array | Hashing | Easy | [1](https://leetcode.com/problems/two-sum/) | Complement lookup trades O(n²) for O(n) with a hash map |
| 2 | [Group Anagrams](dsa/hashing/problems/group-anagrams/README.md) | String | Hashing | Medium | [49](https://leetcode.com/problems/group-anagrams/) | Canonical key (sorted string) groups equivalent items |
| 3 | [Longest Consecutive Sequence](dsa/hashing/problems/longest-consecutive-sequence/README.md) | Array | Hashing | Medium | [128](https://leetcode.com/problems/longest-consecutive-sequence/) | O(n) via set lookups by only starting counts at sequence heads |
| 4 | [Valid Palindrome](dsa/two-pointers/problems/valid-palindrome/README.md) | String | Two Pointers | Easy | [125](https://leetcode.com/problems/valid-palindrome/) | Converging pointers skipping non-alphanumeric chars |
| 5 | [Container With Most Water](dsa/two-pointers/problems/container-with-most-water/README.md) | Array | Two Pointers | Medium | [11](https://leetcode.com/problems/container-with-most-water/) | Greedy pointer move: always move the shorter side |
| 6 | [3Sum](dsa/two-pointers/problems/three-sum/README.md) | Array | Two Pointers | Medium | [15](https://leetcode.com/problems/3sum/) | Fix one element + two pointers on the rest, with dedup |
| 7 | [Max Sum Subarray of Size K](dsa/sliding-window/problems/max-sum-subarray-size-k/README.md) | Array | Sliding Window (fixed) | Easy | [643](https://leetcode.com/problems/maximum-average-subarray-i/)* | Fixed-size window: add right, remove left |
| 8 | [Longest Substring Without Repeating Characters](dsa/sliding-window/problems/longest-substring-without-repeating/README.md) | String | Sliding Window (variable) | Medium | [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | Shrink window on constraint violation using a seen-set/map |
| 9 | [Minimum Window Substring](dsa/sliding-window/problems/minimum-window-substring/README.md) | String | Sliding Window (variable) | Hard | [76](https://leetcode.com/problems/minimum-window-substring/) | Expand to satisfy, then shrink to minimize, with a "have vs need" counter |
| 10 | [Subarray Sum Equals K](dsa/prefix-sum/problems/subarray-sum-equals-k/README.md) | Array | Prefix Sum + Hashing | Medium | [560](https://leetcode.com/problems/subarray-sum-equals-k/) | Complement lookup on running prefix sums, seeded with `{0:1}` |
| 11 | [Range Sum Query - Immutable](dsa/prefix-sum/problems/range-sum-query-immutable/README.md) | Array | Prefix Sum | Easy | [303](https://leetcode.com/problems/range-sum-query-immutable/) | O(1) range queries after O(n) precomputation |
| 12 | [Valid Parentheses](dsa/stack/problems/valid-parentheses/README.md) | String | Stack | Easy | [20](https://leetcode.com/problems/valid-parentheses/) | LIFO matching of nested opener/closer pairs |
| 13 | [Daily Temperatures](dsa/stack/problems/daily-temperatures/README.md) | Array | Monotonic Stack | Medium | [739](https://leetcode.com/problems/daily-temperatures/) | Next-greater-element via a decreasing stack of indices |
| 14 | [Reverse Linked List](dsa/linked-list/problems/reverse-linked-list/README.md) | Linked List | Linked List | Easy | [206](https://leetcode.com/problems/reverse-linked-list/) | In-place pointer reversal with `prev/curr/next` |
| 15 | [Merge Two Sorted Lists](dsa/linked-list/problems/merge-two-sorted-lists/README.md) | Linked List | Linked List | Easy | [21](https://leetcode.com/problems/merge-two-sorted-lists/) | Dummy head + pointer-per-list merge |
| 16 | [Binary Search (Basic)](dsa/binary-search/problems/binary-search-basic/README.md) | Array | Binary Search | Easy | [704](https://leetcode.com/problems/binary-search/) | The canonical lo/hi/mid template and its invariants |
| 17 | [Search in Rotated Sorted Array](dsa/binary-search/problems/search-in-rotated-sorted-array/README.md) | Array | Binary Search on rotated space | Medium | [33](https://leetcode.com/problems/search-in-rotated-sorted-array/) | Identify which half is sorted, then decide which half to keep |
| 18 | [Koko Eating Bananas](dsa/binary-search/problems/koko-eating-bananas/README.md) | Array | Binary Search on the answer | Medium | [875](https://leetcode.com/problems/koko-eating-bananas/) | Binary search over a monotonic feasibility function, not the array |
| 19 | [Subsets](dsa/backtracking/problems/subsets/README.md) | Array | Backtracking | Medium | [78](https://leetcode.com/problems/subsets/) | Include/exclude decision tree over every element |
| 20 | [N-Queens](dsa/backtracking/problems/n-queens/README.md) | Backtracking | Backtracking + pruning | Hard | [51](https://leetcode.com/problems/n-queens/) | Constraint tracking (columns/diagonals) to prune the search space early |
| 21 | [Binary Tree Level Order Traversal](dsa/trees/problems/binary-tree-level-order-traversal/README.md) | Tree | BFS | Easy | [102](https://leetcode.com/problems/binary-tree-level-order-traversal/) | Level-by-level BFS using the current queue length as a level boundary |
| 22 | [Lowest Common Ancestor of a BST](dsa/trees/problems/lowest-common-ancestor-bst/README.md) | Tree | BST property | Easy | [235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | Use BST ordering to decide direction instead of searching both subtrees |
| 23 | [Number of Islands](dsa/graphs/problems/number-of-islands/README.md) | Grid/Graph | Graph traversal (BFS/DFS) | Medium | [200](https://leetcode.com/problems/number-of-islands/) | Grid-as-implicit-graph flood fill, counting connected components |
| 24 | [Course Schedule](dsa/graphs/problems/course-schedule/README.md) | Graph | Topological Sort / Cycle Detection | Medium | [207](https://leetcode.com/problems/course-schedule/) | Kahn's algorithm: a cycle exists iff not all nodes reach in-degree 0 |
| 25 | [Network Delay Time](dsa/graphs/problems/network-delay-time/README.md) | Graph | Dijkstra | Medium | [743](https://leetcode.com/problems/network-delay-time/) | Min-heap finalizes the closest unvisited node first |
| 26 | [Redundant Connection](dsa/union-find/problems/redundant-connection/README.md) | Graph | Union-Find | Medium | [684](https://leetcode.com/problems/redundant-connection/) | First edge whose endpoints are already connected creates the cycle |
| 27 | [Number of Provinces](dsa/union-find/problems/number-of-provinces/README.md) | Graph | Union-Find | Medium | [547](https://leetcode.com/problems/number-of-provinces/) | Component count decreases by one on each successful union |
| 28 | [Single Number](dsa/bit-manipulation/problems/single-number/README.md) | Array | Bit Manipulation | Easy | [136](https://leetcode.com/problems/single-number/) | XOR self-cancellation finds the unpaired element |
| 29 | [Counting Bits](dsa/bit-manipulation/problems/counting-bits/README.md) | Bit Manipulation | Bit Manipulation + DP | Easy | [338](https://leetcode.com/problems/counting-bits/) | `n & (n-1)` recurrence reuses a smaller already-computed popcount |
| 30 | [Jump Game](dsa/greedy/problems/jump-game/README.md) | Array | Greedy | Medium | [55](https://leetcode.com/problems/jump-game/) | Running "furthest reachable" bound, extended greedily |
| 31 | [Gas Station](dsa/greedy/problems/gas-station/README.md) | Array | Greedy | Medium | [134](https://leetcode.com/problems/gas-station/) | Running deficit resets the candidate start in one pass |
| 32 | [Climbing Stairs](dsa/dynamic-programming/problems/climbing-stairs/README.md) | 1D DP | Dynamic Programming | Easy | [70](https://leetcode.com/problems/climbing-stairs/) | Fibonacci-shaped recurrence from counting distinct paths |
| 33 | [House Robber](dsa/dynamic-programming/problems/house-robber/README.md) | 1D DP | Dynamic Programming | Medium | [198](https://leetcode.com/problems/house-robber/) | "Take or skip" state transition with a non-adjacency constraint |
| 34 | [0/1 Knapsack](dsa/dynamic-programming/problems/0-1-knapsack/README.md) | Knapsack DP | Dynamic Programming | Medium | [416](https://leetcode.com/problems/partition-equal-subset-sum/)* | 2D state (item, capacity) collapsible to a 1D rolling array |
| 35 | [Longest Common Subsequence](dsa/dynamic-programming/problems/longest-common-subsequence/README.md) | String DP | Dynamic Programming | Medium | [1143](https://leetcode.com/problems/longest-common-subsequence/) | 2D DP comparing two sequences position by position |
| 36 | [Longest Increasing Subsequence](dsa/dynamic-programming/problems/longest-increasing-subsequence/README.md) | Sequence DP | DP + Binary Search | Medium | [300](https://leetcode.com/problems/longest-increasing-subsequence/) | Patience sorting: track the smallest tail per achievable length |

\* Not directly on LeetCode as framed here — linked to the closest official
equivalent; see that problem's own README for the specific relationship.

More problems are added incrementally — see `PROGRESS.md` for what's planned
next. Every addition must answer "what new idea does this teach?" before it's
accepted.
