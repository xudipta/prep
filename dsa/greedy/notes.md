# Greedy — Quick Revision

- **Idea**: make the locally-best choice at each step, never revisit it.
- **Requires proof**: the greedy-choice property (an exchange argument
  showing the local choice is part of some global optimum) — an unproven
  greedy strategy is just a guess.
- **Common shape**: sort by a problem-specific key, then scan once applying
  an obviously-best rule.
- **Variants**: interval scheduling (sort by end time), interval covering
  (sort by start time), ratio-based (fractional knapsack), running-metric
  one-pass (Jump Game, Gas Station).
- **Complexity**: O(n log n), dominated by the sort.
- **Red flag**: can't explain why the choice is safe → try a
  counterexample; if one exists, it's not actually greedy-solvable (often
  DP instead).
- **Representative problems**: Jump Game (running reachability), Gas
  Station (running fuel deficit).
