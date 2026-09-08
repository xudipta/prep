# DP Pattern Recognition — Decision Framework

Use this before writing any code: read the problem's shape, follow it to a
candidate pattern, then confirm by deriving the recurrence (see the DP
`README.md` for the derivation steps).

## Step 1 — Confirm it's actually DP

Ask: if I wrote the brute-force recursion, would the same subproblem
(same arguments) be computed more than once? If yes → DP. If no (each
subproblem is only ever needed once) → plain recursion or Divide & Conquer,
not DP.

## Step 2 — Match the shape to a pattern

| Signal in the problem | Pattern | State shape |
|---|---|---|
| Answer for `n` depends on answers for `n-1`, `n-2`, ... (counting paths, min cost to reach step `n`) | **1D DP** | `dp[i]` = answer using the first `i` elements/steps |
| Like 1D DP, but you can't use two adjacent elements/steps | **Take/Skip DP** | `dp[i]` = best answer considering index `i`, decision to take or skip it |
| Choose a subset of items under a capacity/weight constraint, each item used at most once | **0/1 Knapsack** | `dp[i][w]` = best value using first `i` items with capacity `w` |
| Same as above but items can be reused unlimited times | **Unbounded Knapsack** | `dp[w]` = best value with capacity `w`, iterating items in the outer or inner loop depending on order needed |
| Compare/align two sequences (edit distance, subsequence matching) | **String DP / LCS-family** | `dp[i][j]` = answer using first `i` chars of A and first `j` chars of B |
| Find the longest strictly-ordered subsequence | **LIS-family** | `dp[i]` = length of the LIS ending at index `i` (O(n²)); or patience-sorting + binary search for O(n log n) |
| Optimal way to combine/split/merge a contiguous *range*, cost depends on how you split it | **Interval DP** | `dp[i][j]` = best answer for the range `[i, j]`, built from splitting at every `k` in between |
| Grid, movement constrained to right/down (or similar), min/max cost path or path count | **2D / Grid DP** | `dp[i][j]` = answer at cell `(i, j)`, built from `dp[i-1][j]` and `dp[i][j-1]` |
| DP over a tree, answer at a node depends on its children's answers | **Tree DP** | `dp[node]` = answer for the subtree rooted at `node`, often as a pair (e.g., "best including node" / "best excluding node") |
| A "state" beyond position also matters (holding stock, on cooldown, previous action) | **State Machine DP** | `dp[i][state]` = best answer at step `i` while in `state` |
| `N` is small (≤ ~20) and you need to track a subset of items/visited nodes | **Bitmask DP** | `dp[mask]` or `dp[i][mask]` = answer given the subset represented by `mask` |
| Counting numbers with digit-level constraints, bound by a numeric range | **Digit DP** | `dp[pos][tight][...]` = count of valid completions from digit position `pos` |
| LIS-style problem where transition itself is monotonic/searchable | **DP + Binary Search** | Same state as LIS-family, but transition uses binary search to find where an element belongs instead of scanning all previous states |
| DAG, answer depends on reachable nodes in topological order | **DP on DAGs** | `dp[node]` = best answer starting/ending at `node`, filled in topological order |

## Step 3 — Derive, don't guess

Once you've matched a candidate pattern, use the six-question framework from
`README.md` (State / Transition / Base Case / Answer / Iteration Order /
Complexity) to derive the actual recurrence for *this* problem — the table
above tells you the shape of the state, not the specific transition, which
still has to come from the problem's own rules.

## Quick sanity checks

- If your recurrence needs information you didn't include in the state,
  your state definition is incomplete — expand it (at the cost of a larger
  state space).
- If two different states in your table would always hold the same value,
  your state has redundant dimensions — shrink it.
- If you can't write a base case without referring to another
  not-yet-defined state, your iteration order is wrong (or you're missing a
  base case).
