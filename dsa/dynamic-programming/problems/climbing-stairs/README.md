# Problem: Climbing Stairs

## Problem Statement

You're climbing a staircase with `n` steps. Each move, you can climb either
1 or 2 steps. In how many distinct ways can you reach the top?

## Difficulty

Easy

## Technique

Dynamic Programming — 1D DP

## Problem Type

1D DP

## Key Insight

The last move to reach step `n` was either a 1-step from `n-1` or a 2-step
from `n-2`. Every way to reach `n` is therefore "a way to reach `n-1`" plus
"a way to reach `n-2`" — giving a Fibonacci-shaped recurrence.

## How to Recognize This Pattern

- "Count the number of ways to reach step/state `n`" where each step has a
  small, fixed set of predecessor moves is the classic 1D counting-DP
  signal.
- Plain recursion on `n` would recompute `climb(n-2)` many times (once
  directly, once via `climb(n-1)` calling it again) — that overlap is the
  DP tell.

## Deriving the Recurrence

1. **Decision**: at the final step to reach `n`, did the last move cover 1
   step or 2 steps?
2. **State**: the only thing that matters is how many steps remain (or
   equivalently, which step you're currently trying to reach) — so
   `dp[i]` = number of distinct ways to reach step `i`.
3. **Transition**: `dp[i] = dp[i-1] + dp[i-2]` (ways ending in a final
   1-step, plus ways ending in a final 2-step).
4. **Base case**: `dp[0] = 1` (one way to "reach" the ground: do nothing),
   `dp[1] = 1` (only one way: a single 1-step).
5. **Answer**: `dp[n]`.
6. **Iteration order**: increasing `i`, since `dp[i]` depends only on
   smaller indices.

```
State:            dp[i] = number of distinct ways to reach step i
Transition:       dp[i] = dp[i-1] + dp[i-2]
Base Case:        dp[0] = 1, dp[1] = 1
Answer:           dp[n]
Iteration Order:  i = 2 .. n
Time Complexity:  O(n)
Space Complexity: O(1) with rolling variables (O(n) for the full table)
```

## Approach 1 — Brute Force (Plain Recursion)

### Idea
Directly recurse: `climb(n) = climb(n-1) + climb(n-2)`.

### Algorithm
Recursive function with base cases at `n <= 1`.

### Complexity
Time: O(2ⁿ) — exponential, due to recomputing the same subproblems
Space: O(n) recursion stack

## Approach 2 — Optimized (Bottom-Up, Space-Optimized)

### Idea
Since `dp[i]` only needs the previous two values, keep two rolling
variables instead of a full array.

### Algorithm
1. `prev2, prev1 := 1, 1` (representing `dp[0]`, `dp[1]`).
2. For `i` from `2` to `n`: `curr := prev1 + prev2`; shift `prev2, prev1 =
   prev1, curr`.
3. Return `prev1`.

### Complexity
Time: O(n)
Space: O(1)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`n = 5`

- `dp[0]=1, dp[1]=1`
- `dp[2]=dp[1]+dp[0]=2`
- `dp[3]=dp[2]+dp[1]=3`
- `dp[4]=dp[3]+dp[2]=5`
- `dp[5]=dp[4]+dp[3]=8`
- Answer: `8`.

## Edge Cases

- `n = 0` → conventionally `1` (one way: do nothing) — confirm this
  assumption with the interviewer, as some phrasings start counting from
  `n = 1`.
- `n = 1` → `1`.
- Large `n` — the space-optimized version avoids O(n) memory; watch for
  integer overflow if `n` is large enough that the Fibonacci-like value
  exceeds the integer type's range (a good follow-up discussion point).

## Common Mistakes

- Off-by-one in base cases (`dp[1] = 1` vs. accidentally starting from
  `dp[1] = 2`).
- Recomputing without memoization, leading to exponential blowup for
  moderately large `n`.
- Not recognizing this as literally the Fibonacci recurrence and
  reinventing it more awkwardly.

## Interview Follow-ups

- What if you can climb 1, 2, or 3 steps at a time? `dp[i] = dp[i-1] +
  dp[i-2] + dp[i-3]` — same derivation, one more term.
- What if certain steps are "broken" and can't be landed on? Add a
  condition: `dp[i] = 0` if step `i` is broken, else the usual sum (with
  care for how broken steps affect reachability).

## Related Problems

- House Robber (same "look back a fixed number of states" 1D DP shape, but
  with a max/skip decision instead of counting)
- Fibonacci Number
- Min Cost Climbing Stairs
