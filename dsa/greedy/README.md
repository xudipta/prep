# Greedy

## Definition

A greedy algorithm builds a solution by always making the choice that looks
best *right now*, never reconsidering it later. Greedy only produces a
correct (optimal) answer when the problem has a specific structural
property — an unproven greedy strategy is a common source of wrong-answer
bugs, not a shortcut.

## Core Intuition

Greedy works when the problem has the **greedy-choice property**: a
locally optimal choice at each step is part of *some* globally optimal
solution, and doesn't need to be revisited once made. This is typically
proven with an **exchange argument**: assume an optimal solution doesn't
make the greedy choice at some step, then show you can modify it to make
that choice without making the solution worse — proving the greedy choice
is "safe."

## Recognition Patterns

- "Minimum number of X to cover/achieve Y" — often interval scheduling or
  covering problems.
- Sorting the input by some criterion makes an obviously-best choice
  visible at each step (earliest deadline, smallest end time, largest
  ratio).
- The problem explicitly allows/expects a "one-pass, never look back"
  solution — and (critically) you can articulate *why* an earlier
  suboptimal-seeming choice can never help later.

## Generic Template

Most greedy algorithms follow this shape:

```go
sort.Slice(items, func(i, j int) bool {
    return items[i].SortKey < items[j].SortKey // problem-specific ordering
})

result := initial()
for _, item := range items {
    if canTake(result, item) { // problem-specific feasibility check
        result = update(result, item)
    }
}
```

The hard part is never the loop — it's correctly identifying the sort key
and the feasibility/update rule, and proving they're safe.

## Variations

- **Interval scheduling** (maximize non-overlapping intervals kept): sort
  by end time, greedily keep an interval if it doesn't overlap the last
  kept one.
- **Interval covering** (minimum intervals to cover a range): sort by start
  time, greedily extend coverage as far as possible before adding a new
  interval.
- **Fractional/ratio greedy** (fractional knapsack, task scheduling by
  deadline): sort by a computed ratio (value/weight, profit/deadline).
- **Two-pointer/one-pass greedy** (Gas Station, Jump Game): a single
  forward pass tracking a running feasibility metric (fuel remaining,
  furthest reachable index).

## Complexity

Typically O(n log n), dominated by the initial sort; O(n) for the greedy
pass itself.

## Common Mistakes

- Applying greedy without proving the greedy-choice property — many
  "obviously greedy" strategies are actually wrong (e.g., naive coin change
  greedy fails for non-canonical coin systems).
- Sorting by the wrong key (e.g., sorting intervals by start time instead
  of end time for the maximum-non-overlapping-intervals problem gives a
  wrong answer on some inputs).
- Confusing "a greedy algorithm exists" with "the first greedy idea I had
  is correct" — always sanity-check against a small counterexample before
  committing.

## Interview Tips

- State the greedy choice explicitly, then give the exchange-argument
  intuition for why it's safe — this is what separates a correct greedy
  solution from a guess that happens to pass the given examples.
- If you can't articulate why the greedy choice is safe, say so — that's a
  signal the problem might actually need DP or backtracking instead.
- Mention the sort as a distinct, explicit step with its own complexity
  contribution (usually the dominant term).

## Example

Jump Game: track the furthest index reachable so far while scanning left to
right; if the current index ever exceeds that furthest-reachable bound,
it's unreachable and the answer is `false`. The greedy choice (always
extend "furthest reachable" as far as possible) is safe because reaching a
farther index can never hurt future reachability. See
[`problems/jump-game`](problems/jump-game/README.md).

## When to Use

- The greedy-choice property can be proven (via an exchange argument) for
  the specific problem.
- Interval scheduling/covering, or a problem cleanly reducible to sorting
  by one criterion and scanning once.

## When NOT to Use

- The "best choice now" can require revisiting later based on a decision
  that depends on future context — that's a signal for DP instead (DP
  considers all choices via subproblem reuse; greedy commits early).
- You can't construct a convincing exchange argument, or you can find a
  counterexample where the greedy choice fails — don't force it.

## Quick Revision

- **Greedy-choice property**: a locally optimal choice is part of some
  global optimum — prove this with an exchange argument before trusting a
  greedy strategy.
- **Common shape**: sort by a key, then scan once making the obviously-best
  choice at each step.
- **Complexity**: O(n log n) typical (dominated by the sort).
- **Red flag**: if you can't explain *why* the greedy choice is safe, it
  might be wrong — verify against a small counterexample.
- **Representative problems**: Jump Game (running reachability bound), Gas
  Station (running fuel deficit + a single pass).
