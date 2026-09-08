# Tips, Tricks & Pattern Recognition

## General Interview Problem-Solving Framework

Before coding:

1. **Clarify input/output.** What are the exact types, ranges, and formats?
   Ask about anything ambiguous (e.g., can the array be empty? Are values
   unique?).
2. **Identify constraints.** Input size bounds tell you the expected time
   complexity (see the complexity-shortcuts table below).
3. **Think about brute force first.** Even if you won't code it, stating it
   gives you a correctness baseline and often reveals the key
   optimization.
4. **Estimate complexity** of the brute force, and of any faster approach
   you can already see.
5. **Look for a pattern.** Use `PATTERN-MAP.md` to match the problem's
   characteristics to a known technique.
6. **Choose the right data structure** for the chosen technique.
7. **Define invariants.** What's always true about your data structure's
   state at each step of the algorithm? Stating this out loud usually
   prevents off-by-one bugs before they happen.
8. **Test edge cases** mentally (or on paper) before writing code.
9. **Analyze final complexity** — time and space — and say it out loud.

## Complexity Shortcuts (Input Size → Expected Complexity)

| Input size (n) | Expected complexity | Typical technique |
|---|---|---|
| ≤ 10-12 | O(2ⁿ), O(n!) | Backtracking, brute force permutations |
| ≤ 20-25 | O(2ⁿ) | Bitmask DP |
| ≤ 500-1,000 | O(n²) or O(n³) | Nested loops, 2D DP |
| ≤ 10⁵-10⁶ | O(n log n) | Sorting, binary search, heap-based approaches |
| ≤ 10⁷-10⁸ | O(n) | Single pass, two pointers, sliding window, hashing |
| Very large / streaming | O(log n) or O(1) per operation | Binary search on the answer, heaps, hashing |

This is a heuristic, not a law — always confirm against the actual stated
constraints, but it's a fast sanity check on whether your planned approach
will pass within typical time limits (usually ~10⁸ simple operations per
second as a rough budget).

## Common Off-by-One Errors

- `<` vs `<=` in loop bounds — especially when the loop should include the
  last valid index (`i <= n-1` vs `i < n`).
- Window/subarray length: `right - left + 1`, not `right - left`.
- Binary search: `lo <= hi` (exact match) vs. `lo < hi` (converging bound
  search) — mixing these up is the most common binary-search bug.
- 1-indexed vs. 0-indexed confusion when a problem statement uses 1-indexed
  positions but your code uses 0-indexed arrays.

## Recursion Checklist

- [ ] Base case(s) defined and correct (what's the smallest input you can
      answer directly?).
- [ ] Each recursive call operates on a strictly smaller subproblem
      (otherwise: infinite recursion).
- [ ] All necessary state is passed as parameters (or captured correctly by
      closure) — don't rely on stale outer-scope variables across calls.
- [ ] Return value/side effect at each level is well-defined and combined
      correctly with the recursive result.
- [ ] Consider recursion depth vs. stack limits for large inputs — an
      iterative rewrite (explicit stack, or bottom-up DP) may be needed.

## DP Checklist

See `dsa/dynamic-programming/README.md` and `dp-pattern-recognition.md` for
the full framework. Quick checklist:

- [ ] Confirmed overlapping subproblems exist (wrote the brute-force
      recursion first).
- [ ] State defined explicitly — and is it *minimal* (no redundant
      dimensions) and *sufficient* (no missing information)?
- [ ] Transition derived from the problem's actual rules, not guessed.
- [ ] Base case(s) correct.
- [ ] Iteration order respects dependencies (every value read from the
      table was already written).
- [ ] Space optimization considered (can rolling variables replace a full
      table?).

## Graph Checklist

- [ ] Representation chosen (adjacency list vs. matrix) matches the
      graph's density and the operations needed.
- [ ] Visited tracking is correct — marked *before* recursing/enqueueing to
      avoid duplicate work or infinite loops on cyclic graphs.
- [ ] Directed vs. undirected cycle detection logic matches the actual
      graph type.
- [ ] Disconnected components handled (loop over all nodes to find every
      component, not just one BFS/DFS from a fixed start).
- [ ] Correct algorithm chosen for the shortest-path variant needed (BFS
      unweighted / Dijkstra non-negative weights / Bellman-Ford negative
      weights allowed).

## Binary Search Checklist

- [ ] Confirmed the predicate/array is actually monotonic before applying
      binary search.
- [ ] Consistent template chosen: `lo <= hi` (exact match) or `lo < hi`
      (bound convergence) — and loop-body updates match the chosen
      template.
- [ ] `mid := lo + (hi-lo)/2` (avoids overflow, though rarely an issue in
      Go specifically — still good habit).
- [ ] For "binary search on the answer" problems: `feasible(x)` explicitly
      defined and its monotonicity justified before coding.

## Sliding-Window Checklist

- [ ] Fixed vs. variable window correctly identified.
- [ ] Window state (sum, frequency map, counters) updates in O(1) as
      elements enter/leave.
- [ ] Shrink condition matches the goal: "while invalid, shrink" (maximize)
      vs. "while still valid, shrink and record" (minimize).
- [ ] Window length computed as `right - left + 1`.

## Tree Checklist

- [ ] Correct traversal order chosen (preorder/inorder/postorder/level
      order) based on when you need a node's value relative to its
      children's results.
- [ ] `nil` base case present in every recursive traversal.
- [ ] BST property exploited when available (don't fall back to O(n)
      generic search unnecessarily).
- [ ] For "path" or "diameter"-style problems: confirmed the answer may
      not pass through the root, and a separate global-best variable is
      tracked alongside the per-node returned value.

## Debugging Checklist (When Your Solution Is Wrong)

1. Re-read the problem statement — is there a constraint or edge case you
   missed?
2. Trace through the smallest failing example by hand, step by step.
3. Check loop bounds and off-by-one errors first — they're the most common
   bug source.
4. Verify base cases (recursion/DP) are actually being hit correctly.
5. For a "runtime error" (not "wrong answer"): check for nil dereferences,
   out-of-bounds indices, and integer overflow.
6. For a "wrong answer" that passes some cases: check whether your
   assumption about sortedness/monotonicity/uniqueness actually holds for
   the failing case.
7. For a "time limit exceeded": re-derive the expected complexity from the
   constraints table above, and check whether your implementation actually
   matches the complexity you intended (e.g., an accidental O(n) operation
   inside an O(n) loop, making it O(n²)).

## How to Approach an Unfamiliar Problem

1. Restate the problem in your own words to confirm understanding.
2. Work a small example by hand before writing any code.
3. Identify what's being asked: a value (optimal count/sum), a
   yes/no decision, an enumeration of all solutions, or a construction
   (build a specific structure). This alone rules out entire technique
   families (e.g., "enumerate all" almost never wants DP; "optimal value
   with overlapping subproblems" almost never wants backtracking).
4. Map the problem's characteristics against `PATTERN-MAP.md`.
5. If nothing matches cleanly, fall back to brute force, then look for
   what's redundant/repeated in that brute force — that redundancy is
   usually exactly what the optimization exploits.

## How to Optimize a Brute Force

- Repeated lookups/membership checks → hashing.
- Repeated subproblem computation → memoization/DP.
- Sorted or sortable input with pairwise comparisons → two pointers or
  binary search.
- Contiguous range recomputation → sliding window or prefix sums.
- Need all valid configurations, with prunable partial states →
  backtracking.
- Need connectivity/shortest-path structure → graph algorithms.

## Quick Revision

- **Process**: clarify → constrain → brute force → complexity → pattern →
  data structure → invariant → edge cases → final complexity.
- **Size → complexity**: n≤12 exponential; n≤1000 quadratic; n≤10⁶ n log n;
  n≤10⁸ linear.
- **Off-by-ones**: loop bounds, window length, binary search template
  choice, 0- vs 1-indexing.
- **When stuck**: state the brute force, find what's redundant in it — that
  redundancy is the optimization target.
