# Backtracking

## Definition

Backtracking systematically explores a decision tree — try a choice, recurse
on the consequences, then undo the choice ("backtrack") and try the next
one. It's exhaustive search with early pruning: abandon a branch as soon as
it's known to be invalid, rather than continuing to build on top of it.

## Core Intuition

Every backtracking problem can be framed as: "build a solution one decision
at a time; at each step, try every valid choice for the current decision,
recurse, then undo." The word "undo" is what separates it from plain
recursion — you're reusing the same mutable state (an array, a set, a
board) across branches instead of copying it, so you must carefully revert
any mutation before trying the next choice.

## Recognition Patterns

- "Generate all subsets/permutations/combinations."
- "Find all valid ways to place/arrange items under constraints" (N-Queens,
  Sudoku).
- The search space is a decision tree with a **branching factor** at each
  step and a **depth** equal to the number of decisions.
- Constraints can be checked *incrementally* (partial solutions can be
  pruned before being fully built) — this is what makes backtracking faster
  than brute-force enumeration.

## Generic Template

```go
func backtrack(path []int, remainingChoices []int, results *[][]int) {
    if isCompleteSolution(path) {
        // copy path before appending — path is mutated by the caller
        combo := make([]int, len(path))
        copy(combo, path)
        *results = append(*results, combo)
        return
    }
    for i, choice := range remainingChoices {
        if !isValid(path, choice) {
            continue // prune
        }
        path = append(path, choice)                 // choose
        backtrack(path, nextChoices(remainingChoices, i), results)
        path = path[:len(path)-1]                    // un-choose
    }
}
```

## Variations

- **Include/exclude** decision at each element (Subsets): binary decision
  tree, `2ⁿ` total leaves.
- **Permutation-style**: choose one of the *remaining* unused elements at
  each step (uses a "used" marker instead of a simple index cursor).
- **Constraint-satisfaction with heavy pruning** (N-Queens, Sudoku): track
  auxiliary state (occupied columns/diagonals) to reject invalid partial
  placements in O(1) instead of re-validating the whole board.
- **Combination-with-repetition** (Combination Sum): allow reusing the same
  element, controlled by whether the recursive call advances the index.

## Complexity

Backtracking is inherently exponential in the worst case (it's exhaustive
search) — the goal isn't to make it polynomial, but to **prune early** so
the exponent's base/constant factor is as small as possible. State time
complexity as "branching factor ^ depth," and explicitly describe what
pruning removes from that bound.

## Common Mistakes

- Forgetting to undo a mutation (the "un-choose" step), which corrupts
  state for sibling branches.
- Appending a mutable slice reference to the results directly instead of
  copying it — since the underlying array keeps changing, every stored
  "answer" ends up pointing to the same (wrong) final state.
- Not pruning early enough — checking validity only once a solution is
  fully built, instead of rejecting invalid partial states as soon as
  possible.
- Off-by-one in the "which choices remain" logic, causing duplicate or
  missing results.

## Interview Tips

- State the decision being made at each recursive call explicitly ("at each
  step, I either include or exclude the current element").
- Call out what's being pruned and why it's safe to prune (an invariant
  that, once violated, can never be fixed by future choices).
- Mention the exponential worst case honestly — don't claim backtracking is
  polynomial; the value is in constant-factor pruning, not asymptotic
  improvement.

## Example

N-Queens: place queens one row at a time; track occupied columns and both
diagonals as sets, so checking "is this column/diagonal safe?" is O(1)
instead of re-scanning the board. See
[`problems/n-queens`](problems/n-queens/README.md).

## When to Use

- Need to enumerate *all* valid solutions/combinations, not just one
  optimal value (if you only need an optimal value, DP or greedy is often
  a better fit).
- Constraints can be checked incrementally, enabling meaningful pruning.

## When NOT to Use

- You only need one answer's *value* (max/min/count), and the problem has
  overlapping subproblems — that's DP, which reuses subproblem results
  instead of re-exploring them.
- The search space is too large to prune effectively, and no polynomial
  algorithm exists — that's a sign to look for approximations or a
  different problem framing (rare in interviews, but worth naming as a
  limitation).

## Quick Revision

- **Template**: choose → recurse → un-choose (undo the mutation before
  trying the next option).
- **Copy before storing**: never append a mutable path slice directly to
  the results list.
- **Prune early**: reject invalid partial states as soon as detectable, not
  only at the leaves.
- **Complexity**: branching factor ^ depth, reduced by pruning — state this
  honestly rather than claiming polynomial time.
- **Representative problems**: Subsets (include/exclude), N-Queens
  (constraint tracking + heavy pruning).
