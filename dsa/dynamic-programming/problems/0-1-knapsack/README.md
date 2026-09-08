# Problem: 0/1 Knapsack

## Problem Statement

Given `n` items, each with a weight `weights[i]` and a value `values[i]`,
and a knapsack with capacity `W`, find the maximum total value obtainable by
choosing a subset of items whose total weight doesn't exceed `W`. Each item
can be used at most once (hence "0/1": include it or don't).

## Difficulty

Medium

## Technique

Dynamic Programming — Knapsack

## Problem Type

Knapsack DP

## Key Insight

For each item, in order, decide "include it or not." Including it is only
valid if there's enough remaining capacity, and it changes the *remaining
capacity* available to later decisions. Since both "which items considered
so far" and "capacity used so far" affect future decisions, the state needs
both dimensions.

## How to Recognize This Pattern

- "Choose a subset of items under a capacity/weight/budget constraint to
  maximize value" is the canonical 0/1 knapsack signal.
- Each item usable **at most once** (if items were reusable, it would be
  *unbounded* knapsack instead, with a different iteration order).

## Deriving the Recurrence

1. **Decision**: for item `i`, include it in the knapsack or not.
2. **State**: `dp[i][w]` = the maximum value obtainable using only the
   first `i` items with capacity `w`. Both dimensions are necessary: which
   items have been considered (to avoid reusing one), and how much capacity
   remains (to know what still fits).
3. **Transition**:
   - If `weights[i-1] > w`: item `i-1` can't fit — `dp[i][w] = dp[i-1][w]`.
   - Otherwise: `dp[i][w] = max(dp[i-1][w], values[i-1] + dp[i-1][w -
     weights[i-1]])` (better of skipping vs. including item `i-1`).
4. **Base case**: `dp[0][w] = 0` for all `w` (no items, no value).
5. **Answer**: `dp[n][W]`.
6. **Iteration order**: increasing `i` (items), and for each `i`, any order
   over `w` works in the 2D formulation since `dp[i][*]` only reads from
   `dp[i-1][*]`. In the space-optimized 1D formulation, `w` must be
   iterated **descending** so that `dp[w - weight]` still refers to the
   previous item's row, not the current item's already-updated value
   (which would incorrectly allow reusing the item).

```
State:            dp[i][w] = max value using first i items, capacity w
Transition:       dp[i][w] = dp[i-1][w]                                  if weights[i-1] > w
                   dp[i][w] = max(dp[i-1][w], values[i-1] + dp[i-1][w-weights[i-1]])  otherwise
Base Case:        dp[0][w] = 0 for all w
Answer:           dp[n][W]
Iteration Order:  i ascending; w descending in the 1D space-optimized form
Time Complexity:  O(n * W)
Space Complexity: O(W) with the 1D rolling-row optimization (O(n*W) for the full table)
```

## Approach 1 — Brute Force (Try All Subsets)

### Idea
Try every subset of items, check the weight constraint, track the best
value.

### Algorithm
Recursive include/exclude over all `n` items, or enumerate all `2ⁿ` subset
masks.

### Complexity
Time: O(2ⁿ)
Space: O(n) recursion stack

## Approach 2 — Optimized (1D Tabulation, Space-Optimized)

### Idea
Use a single array `dp[w]` representing "best value achievable with
capacity `w`, using items considered so far." Process items one at a time,
updating `dp` **from high capacity to low** so each item is only used once
per pass.

### Algorithm
1. `dp := make([]int, W+1)` (all zero).
2. For each item `i`:
   - For `w` from `W` down to `weights[i]`:
     - `dp[w] = max(dp[w], values[i] + dp[w-weights[i]])`.
3. Return `dp[W]`.

### Complexity
Time: O(n * W)
Space: O(W)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`weights = [1, 3, 4, 5]`, `values = [1, 4, 5, 7]`, `W = 7`

- Start `dp = [0,0,0,0,0,0,0,0]` (indices 0..7).
- Item 0 (w=1, v=1): for `w=7..1`: `dp[w] = max(dp[w], 1+dp[w-1])` → after
  this pass, `dp = [0,1,1,1,1,1,1,1]`.
- Item 1 (w=3, v=4): for `w=7..3`: e.g. `dp[7]=max(1, 4+dp[4])=max(1,4+1)=5`;
  `dp[6]=max(1,4+dp[3])=max(1,4+1)=5`; `dp[5]=max(1,4+dp[2])=5`;
  `dp[4]=max(1,4+dp[1])=5`; `dp[3]=max(1,4+dp[0])=4`. →
  `dp = [0,1,1,4,5,5,5,5]`.
- Item 2 (w=4, v=5): `dp[7]=max(5,5+dp[3])=max(5,5+4)=9`;
  `dp[6]=max(5,5+dp[2])=max(5,5+1)=6`; `dp[5]=max(5,5+dp[1])=max(5,6)=6`;
  `dp[4]=max(5,5+dp[0])=5`. → `dp = [0,1,1,4,5,6,6,9]`.
- Item 3 (w=5, v=7): `dp[7]=max(9,7+dp[2])=max(9,8)=9`;
  `dp[6]=max(6,7+dp[1])=max(6,8)=8`; `dp[5]=max(6,7+dp[0])=7`. →
  `dp = [0,1,1,4,5,7,8,9]`.
- Answer: `dp[7] = 9` (items with weight 1 and weight 4: values 1+5=6... —
  actually the optimal is items 1 and 2, weights 3+4=7, values 4+5=9).

## Edge Cases

- `W = 0` → answer is `0` (nothing fits).
- An item heavier than `W` on its own → never selectable, correctly
  excluded by the `w >= weights[i]` loop bound.
- No items → answer is `0`.
- All items fit within `W` → answer is the sum of all values.

## Common Mistakes

- Iterating `w` **ascending** in the 1D optimization — this allows an item
  to be counted multiple times (turns 0/1 knapsack into unbounded
  knapsack), which is wrong here.
- Forgetting the `w >= weights[i]` bound and indexing `dp` with a negative
  capacity.
- Confusing 0/1 knapsack (this problem) with unbounded knapsack (coin
  change-style problems where items are reusable) — they look similar but
  need different iteration orders.

## Interview Follow-ups

- Unbounded Knapsack (Coin Change): same recurrence shape but items are
  reusable — iterate `w` ascending instead of descending.
- Print the actual chosen subset, not just the max value: retain the full
  2D table and backtrack from `dp[n][W]`.
- Subset Sum (decision version: does *any* subset sum to exactly `W`?) is
  0/1 knapsack with `values[i] = weights[i]` and a boolean/reachability
  table instead of a max-value table.

## Related Problems

- Subset Sum
- Coin Change (unbounded knapsack)
- Partition Equal Subset Sum
