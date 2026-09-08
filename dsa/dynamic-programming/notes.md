# Dynamic Programming — Quick Revision

- **Definition**: optimal substructure + overlapping subproblems.
- **Process**: write the brute-force recursion → find the state → cache
  (top-down) or tabulate in dependency order (bottom-up).
- **Always answer**: State / Transition / Base Case / Answer / Iteration
  Order / Time / Space — for every DP problem, no exceptions.
- **Space optimization**: if `dp[i]` only depends on the last 1-2 rows/
  states, collapse the table to rolling variables.
- **1D DP**: `dp[i]` depends on a few smaller indices (Climbing Stairs,
  House Robber).
- **Knapsack**: `dp[i][w]` = items considered × capacity used; 0/1 iterates
  capacity descending in the 1D form, unbounded iterates ascending.
- **String DP / LCS / LIS / Interval / Tree / Bitmask / Digit DP**: see
  `dp-pattern-recognition.md` for the full shape-to-pattern map.
- **Common bugs**: wrong iteration order (reading an unwritten cell),
  off-by-one base cases, ascending vs. descending capacity loop in
  knapsack variants.
- **Representative problems**: Climbing Stairs (1D), House Robber
  (take/skip), 0/1 Knapsack (2D → 1D rolling).
