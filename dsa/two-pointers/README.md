# Two Pointers

## Definition

Two Pointers is a technique where two indices traverse a data structure
(usually an array or string) — either moving toward each other, moving in
the same direction at different speeds, or one anchored while the other
scans — to avoid the nested loops a brute-force solution would need.

## Core Intuition

Whenever a brute-force solution checks every pair `(i, j)`, ask: does sorting
or the existing order let me rule out pairs without checking them
individually? If moving one pointer can only make the answer better or worse
in a predictable direction, you can discard a whole set of pairs at once
instead of checking them one by one — that's the O(n²) → O(n) (or O(n log n)
with a sort) win.

## Recognition Patterns

- The array is sorted, or can be sorted without losing needed information
  (e.g., you need values, not original indices).
- You're looking for a pair/triple with a target sum, product, or relation.
- You're comparing values from both ends of a structure (palindrome checks).
- You need the max/min of some function over a "window" defined by two ends,
  and moving the worse end can never hurt.
- Merging two sorted structures.

## Generic Template

**Converging pointers** (opposite ends, move inward):

```go
lo, hi := 0, len(arr)-1
for lo < hi {
    switch {
    case condition(arr[lo], arr[hi]):
        // record/update answer
        lo++
        hi--
    case needSmaller(arr[lo], arr[hi]):
        hi--
    default:
        lo++
    }
}
```

**Fast/slow pointers** (same direction, different speeds — e.g., cycle
detection, removing duplicates in place):

```go
slow := 0
for fast := 0; fast < len(arr); fast++ {
    if keep(arr[fast]) {
        arr[slow] = arr[fast]
        slow++
    }
}
```

## Variations

- **Converging**: both pointers start at the ends and move inward (Container
  With Most Water, Valid Palindrome, Two Sum on a sorted array).
- **Fixed anchor + scanning**: fix one element, two-pointer the rest (3Sum,
  4Sum).
- **Fast/slow (same direction)**: in-place array compaction, cycle detection
  in linked lists (Floyd's algorithm).
- **Merge-style**: one pointer per input, advancing whichever is smaller
  (merging two sorted arrays, merge step of merge sort).

## Complexity

Typically O(n) time after an O(n log n) sort (or O(n) if the input is
already sorted/no sort needed), O(1) extra space — the main advantage over
a hash-map approach when memory matters.

## Common Mistakes

- Forgetting to skip duplicate values when the problem asks for *unique*
  results (e.g., 3Sum) — leads to duplicate triples in the output.
- Moving both pointers when only one should move, silently skipping valid
  answers.
- Assuming two pointers works without checking that the "move toward a
  direction" logic is actually monotonic/greedy-safe.
- Off-by-one errors in the loop condition (`<` vs `<=` when `lo == hi` is a
  valid state, e.g., single-element windows).

## Interview Tips

- State the invariant out loud: "everything between `lo` and `hi` is still a
  candidate" or "everything left of `slow` is already processed/kept."
- If you're not sure two pointers applies, try to argue why moving the
  "worse" pointer can never skip the optimal answer — if you can't, it's
  probably not a valid two-pointer problem.
- Mention the brute force (O(n²)) first, then explain what sorting/ordering
  buys you.

## Example

Container With Most Water: brute force checks all O(n²) pairs of lines. Two
pointers starts at the widest container and always moves the shorter line
inward, because keeping the shorter line can never produce a taller
container with less width. See
[`problems/container-with-most-water`](problems/container-with-most-water/README.md).

## When to Use

- Sorted (or sortable) arrays with pair/triple-sum problems.
- Palindrome/mirror checks.
- Merging sorted sequences.
- In-place array partitioning/compaction.

## When NOT to Use

- When the order/positions of original elements matter and sorting would
  destroy that information (unless you carry original indices alongside).
- When there's no monotonic relationship to exploit — if moving a pointer
  doesn't predictably improve/worsen the objective, you likely need a
  different technique (e.g., hashing, DP).
- When you need *all* pairs, not an optimal one — two pointers usually finds
  one direction of improvement, not an exhaustive enumeration.

## Quick Revision

- **What**: two indices scanning a linear structure instead of nested loops.
- **When**: sorted/sortable data, pair/triple sums, palindrome checks, merging.
- **Converging**: `lo, hi` move inward; move the pointer that "can't be part
  of a better answer."
- **Fast/slow**: same direction, different speeds; used for compaction/cycle
  detection.
- **Complexity win**: O(n²) brute force → O(n) or O(n log n) with sort.
- **Watch for**: duplicate skipping, correct loop bound (`<` vs `<=`).
