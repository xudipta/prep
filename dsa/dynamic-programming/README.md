# Dynamic Programming

Dynamic Programming (DP) gets special treatment in this repository because
it's the broadest technique here — it isn't one pattern but a family of
patterns unified by one idea: **solve overlapping subproblems once, and
reuse the result.**

## Definition

DP applies to problems with two properties:

1. **Optimal substructure** — an optimal solution to the problem can be
   built from optimal solutions to its subproblems.
2. **Overlapping subproblems** — the same subproblem is needed multiple
   times if solved naively via plain recursion.

If a problem has optimal substructure but *not* overlapping subproblems
(e.g., classic Divide & Conquer like merge sort), memoization buys you
nothing — that's a Divide & Conquer problem, not a DP one.

## Core Intuition

Every DP problem starts life as a recursive brute force: "the answer for
input `n` depends on the answer for some smaller input(s)." DP's job is
purely mechanical from there:

1. Write the brute-force recursion first.
2. Identify what varies between recursive calls — that's your **state**.
3. Notice the same state gets recomputed — cache it (top-down memoization),
   or compute states in dependency order and store them in a table
   (bottom-up tabulation).

Don't skip step 1. A recurrence "presented" without deriving it from the
brute-force recursion is the single biggest reason DP feels like magic
instead of mechanical — see the pattern-recognition guide for a full
decision framework.

## How to Derive a Recurrence (not just present one)

For any DP problem, answer these in order:

1. **What decision am I making at each step?** (Take this item or not?
   Match these two characters or not? Cut here or not?)
2. **What information do I need to make that decision correctly?** That
   information — the minimum set of variables that fully determines the
   subproblem — is your **state**.
3. **Given a state, what are the possible decisions, and what smaller state
   does each decision lead to?** That's your **transition**.
4. **What's the smallest/simplest state I can answer directly, without
   recursing further?** That's your **base case**.
5. **Which state holds the final answer?**
6. **In what order must I fill states so that every transition's dependency
   is already computed?** That's your **iteration order** (only relevant
   for bottom-up tabulation).

Every DP problem write-up in this repo answers all six explicitly:

```
State:
Transition:
Base Case:
Answer:
Iteration Order:
Time Complexity:
Space Complexity:
```

## DP Patterns Covered

| Pattern | Status | Notes |
|---|---|---|
| 1D DP | ✅ | [Climbing Stairs](problems/climbing-stairs/README.md) |
| Take/Skip (adjacency-constrained 1D DP) | ✅ | [House Robber](problems/house-robber/README.md) |
| 0/1 Knapsack | ✅ | [0/1 Knapsack](problems/0-1-knapsack/README.md) |
| 2D / Grid DP | planned | e.g. Unique Paths, Minimum Path Sum |
| Subset Sum | planned | variant of 0/1 Knapsack with a boolean objective |
| Partition DP | planned | e.g. Palindrome Partitioning |
| String DP | planned | e.g. Edit Distance |
| Longest Common Subsequence | planned | classic 2D string DP |
| Longest Increasing Subsequence | planned | O(n²) and O(n log n) variants |
| Interval DP | planned | e.g. Matrix Chain Multiplication, Burst Balloons |
| Tree DP | planned | e.g. House Robber III, Diameter with DP |
| State Machine DP | planned | e.g. Best Time to Buy/Sell Stock with cooldown |
| Bitmask DP | planned | e.g. Traveling Salesman on small N |
| Digit DP | planned | e.g. counting numbers with a digit property |
| DP + Binary Search | planned | e.g. LIS in O(n log n) |
| DP on DAGs | planned | e.g. Longest Path in a DAG |

See `dp-pattern-recognition.md` for the full decision framework used to pick
a pattern for a new problem.

## Complexity Habits

- State space size × transition cost = total time complexity. Always state
  both factors, not just the final product.
- Space can often be reduced from O(n·m) to O(m) (or O(1)) when the
  transition only depends on the previous row/state — call this out
  explicitly in every write-up ("space-optimized" section).

## Common Mistakes

- Presenting a recurrence without deriving it from a brute-force recursion
  — makes DP feel unlearnable. Always derive.
- Wrong iteration order in tabulation (reading a table cell before it's
  been written).
- Off-by-one in base cases (`dp[0]` vs `dp[1]` as the starting index).
- Forgetting to consider *all* valid transitions from a state (missing a
  case in the recurrence).
- Reaching for DP when the problem doesn't have overlapping subproblems
  (wasted complexity) or when a simpler technique — greedy, two pointers —
  already solves it optimally.

## Interview Tips

- Always state the recursive brute force and its complexity first, then
  show what memoization/tabulation buys you.
- Explicitly name the state variables and justify why they're sufficient
  (why you don't need *more* state than that).
- Mention the space optimization even if you code the full table first —
  it signals you understand the dependency structure, not just the
  recurrence.

## When to Use

- Optimal substructure + overlapping subproblems, confirmed by having
  written out the brute-force recursion first.

## When NOT to Use

- No overlapping subproblems (plain divide & conquer, or a problem solvable
  greedily with a provable exchange argument).
- The state space is too large to store or compute in the time/memory
  budget — look for a way to shrink the state, or consider whether the
  problem actually needs an approximation/heuristic instead.

## Quick Revision

- **Definition**: optimal substructure + overlapping subproblems.
- **Process**: brute-force recursion → identify state → memoize (top-down)
  or tabulate (bottom-up).
- **Always state**: State / Transition / Base Case / Answer / Iteration
  Order / Time / Space.
- **Space optimization**: if transition only needs the previous
  row/state(s), collapse the table to O(1) or O(current dimension) space.
- **See also**: `dp-pattern-recognition.md` for choosing the right pattern.
