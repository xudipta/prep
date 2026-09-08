# Pattern Map — From Problem Characteristics to Technique

A decision framework: read the left column (what you notice about the
problem), follow it to a likely pattern, then to representative problems in
this repo. This is a heuristic map, not a strict algorithm — many problems
combine two or three of these.

```
Problem characteristic
        ↓
Likely pattern
        ↓
Recommended technique
        ↓
Representative problems
```

## Arrays & Strings

| Characteristic | Likely pattern | Technique | Representative problems |
|---|---|---|---|
| Need pairs/triples summing to a target, array sortable | Fix + converge | Two Pointers | Three Sum |
| Two ends of a sorted/monotonic structure, want max/min of a function of both ends | Converging pointers | Two Pointers | Container With Most Water |
| Contiguous subarray/substring, fixed length K | Fixed window | Sliding Window | Max Sum Subarray of Size K |
| Contiguous subarray/substring, "longest/shortest satisfying a constraint" | Variable window | Sliding Window | Longest Substring Without Repeating Characters, Minimum Window Substring |
| Need O(1) lookup of "have I seen this / what's its count" | Hash map/set | Hashing | Two Sum, Longest Consecutive Sequence |
| Need to group items by an equivalence relation | Canonical key hashing | Hashing | Group Anagrams |
| Sorted (or rotated-sorted) array, need an index/value | Binary Search | Binary Search | Binary Search (Basic), Search in Rotated Sorted Array |
| "Minimize the maximum" / "maximize the minimum", answer space is monotonic | Binary search on the answer | Binary Search | Koko Eating Bananas |
| Need every subset/combination/permutation | Decision tree with undo | Backtracking | Subsets, N-Queens |
| Constraints prune a large search space (board positions, sudoku) | Backtracking + pruning | Backtracking | N-Queens |
| Need the sum/count of a contiguous range, queried repeatedly on a static array | Precompute cumulative sums | Prefix Sum | Range Sum Query - Immutable |
| Number of subarrays with sum equal to K, negative numbers present | Prefix sum + hashing | Prefix Sum | Subarray Sum Equals K |
| Minimize/maximize X to achieve Y, a locally-best choice provably safe (exchange argument) | Sort + greedy scan | Greedy | Jump Game, Gas Station |
| Matching/nesting of paired tokens | Stack (LIFO) | Stack | Valid Parentheses |
| Next/previous greater or smaller element, days until warmer | Monotonic stack | Stack | Daily Temperatures |
| Reverse/reorder/rearrange a linked list, or find a node via fast/slow pointers | Pointer rewiring | Linked List | Reverse Linked List |
| Merge multiple sorted linked structures | Pointer-per-list merge (+ heap for k) | Linked List | Merge Two Sorted Lists |
| Every element appears twice except one, O(1) space required | XOR self-cancellation | Bit Manipulation | Single Number |
| Per-number property computed for every number up to n | Bit trick + DP recurrence | Bit Manipulation | Counting Bits |
| Groups form incrementally from pairwise connections; need "same group?" queries | Union-Find | Union-Find | Number of Provinces |
| Find the edge that creates a cycle while building a graph incrementally | Union-Find cycle detection | Union-Find | Redundant Connection |

## Trees & Graphs

| Characteristic | Likely pattern | Technique | Representative problems |
|---|---|---|---|
| Level-by-level processing of a tree | BFS with level boundary | Trees / BFS | Binary Tree Level Order Traversal |
| BST + "find ancestor / path" | Exploit ordering | Trees / BST | Lowest Common Ancestor of a BST |
| Grid where cells connect to neighbors | Grid-as-graph | Graphs / BFS-DFS | Number of Islands |
| "Can all tasks/courses be completed given dependencies" | Cycle detection via topological sort | Graphs / Topological Sort | Course Schedule |
| Need shortest path, unweighted graph | BFS | Graphs | (planned: Word Ladder) |
| Need shortest path, weighted, non-negative | Dijkstra | Graphs / Shortest Path | Network Delay Time |
| Need shortest path, negative weights allowed | Bellman-Ford | Graphs / Shortest Path | (planned) |
| Need connectivity / cycle detection with union operations | Union-Find (DSU) | Union-Find | Redundant Connection, Number of Provinces |
| Need to connect all nodes with minimum total edge weight | MST (Kruskal/Prim) | Graphs / MST | (planned) |

## Dynamic Programming

| Characteristic | Likely pattern | Sub-pattern | Representative problems |
|---|---|---|---|
| "Count the number of ways" / "min/max cost to reach state i" with a small, linear state | 1D DP | 1D DP | Climbing Stairs |
| Adjacent-element constraint ("can't pick two neighbors") | Take/skip DP | 1D DP | House Robber |
| Choose a subset of items under a capacity constraint | Knapsack | 0/1 or unbounded knapsack | 0/1 Knapsack |
| Compare two sequences | LCS-family | String DP | Longest Common Subsequence |
| Find the longest ordered subsequence | LIS-family | Sequence DP | Longest Increasing Subsequence |
| Optimal way to combine/merge a contiguous range | Interval DP | Interval DP | (planned: Matrix Chain / Burst Balloons) |
| DP over tree structure | Tree DP | Tree DP | (planned) |
| Small N (≤ ~20) with subsets as state | Bitmask DP | Bitmask DP | (planned: TSP) |

See `dsa/dynamic-programming/dp-pattern-recognition.md` for the full decision
framework.

## Rule of thumb when nothing above matches

1. Can you brute force it? What's the complexity?
2. Is there redundant work being repeated? → Hashing, memoization/DP.
3. Is the input sorted or sortable? → Two pointers, binary search.
4. Is the answer monotonic in some parameter? → Binary search on the answer.
5. Do you need "all ways"? → Backtracking / DFS.
6. Is there an underlying graph (explicit or implicit, e.g. a grid)? → BFS/DFS/graph algorithms.
7. Is there a greedy local choice that's provably safe? → Greedy (prove exchange argument before trusting it).
