# Backtracking — Quick Revision

- **Template**: choose → recurse → un-choose. Reuses mutable state across
  branches instead of copying it.
- **Copy before storing**: never append a live, still-mutating path slice
  directly into the results.
- **Prune early**: reject invalid partial states as soon as detectable —
  this is what makes backtracking fast in practice despite being
  exponential in the worst case.
- **Common shapes**: include/exclude (Subsets), choose-from-remaining
  (Permutations), constraint tracking with O(1) checks (N-Queens, Sudoku).
- **Complexity**: branching factor ^ depth; state honestly, don't claim
  polynomial time.
- **Representative problems**: Subsets (include/exclude), N-Queens
  (row-by-row placement + column/diagonal set tracking).
