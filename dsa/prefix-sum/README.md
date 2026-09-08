# Prefix Sum

## Definition

A prefix sum array `P` where `P[i]` is the sum of all elements `arr[0..i-1]`
lets you compute the sum of any contiguous range `arr[i..j]` in O(1) as
`P[j+1] - P[i]`, after an O(n) one-time preprocessing pass — instead of
O(n) per range-sum query.

## Core Intuition

If a problem repeatedly asks "what's the sum (or count, or XOR) of this
contiguous range?", and the underlying array doesn't change between
queries, precompute cumulative values once. The prefix-sum + hash-map
combination extends this to "find a subarray whose sum equals K" in O(n):
`sum(i..j) == K` is equivalent to `P[j+1] - P[i] == K`, i.e. `P[i] ==
P[j+1] - K` — a complement lookup, exactly like `dsa/hashing`.

## Recognition Patterns

- Multiple range-sum queries over a static (or rarely-updated) array.
- "Number of subarrays with sum/product/XOR equal to K."
- "Equilibrium index" / "find a split point where both sides are equal."

## Generic Template

**Static range-sum queries**:

```go
prefix := make([]int, len(arr)+1)
for i, v := range arr {
    prefix[i+1] = prefix[i] + v
}
rangeSum := func(i, j int) int { return prefix[j+1] - prefix[i] } // sum of arr[i..j]
```

**Subarray sum equals K (prefix sum + hash map)**:

```go
count := 0
sum := 0
seen := map[int]int{0: 1} // empty prefix sums to 0, seen once
for _, v := range arr {
    sum += v
    count += seen[sum-k]
    seen[sum]++
}
```

## Variations

- **1D range sum** (the basic template above).
- **2D prefix sum**: precompute cumulative sums over a matrix for O(1)
  rectangle-sum queries.
- **Prefix sum + hash map**: for "subarray sum equals K" style problems,
  turning an O(n²) brute force into O(n).
- **Prefix XOR**: same idea, substituting XOR for sum (XOR is its own
  inverse, so `xor(i..j) = P[j+1] ^ P[i]`).

## Complexity

O(n) preprocessing, O(1) per range query (O(n) with the hash-map variant
for subarray-count problems, since it's one pass with O(1) map operations).

## Common Mistakes

- Off-by-one in the prefix array's indexing (`prefix[i+1]` corresponds to
  the sum of the first `i` elements — a common source of confusion is
  whether `prefix[i]` includes `arr[i]` or not; pick one convention and
  stay consistent).
- Forgetting to seed the hash map with `{0: 1}` in the subarray-sum-equals-K
  pattern — without it, subarrays starting at index 0 are never counted.
- Using prefix sums for a *mutable* array without a strategy for updates —
  plain prefix sums require O(n) to update after a single element changes;
  a Fenwick Tree (Binary Indexed Tree) is the right structure when updates
  and range queries are both needed.

## Interview Tips

- Explicitly state the prefix array's indexing convention
  (`prefix[i] = sum(arr[0..i-1])`, so `prefix[0] = 0`) before coding — this
  is the single most common point of confusion.
- For subarray-count problems, connect the technique back to hashing
  explicitly: "this is a complement lookup on prefix sums."

## Example

Subarray Sum Equals K: rather than checking every `(i, j)` pair (O(n²)),
track the count of each prefix sum seen so far; for each new prefix sum,
check how many earlier prefix sums equal `current - k`. See
[`problems/subarray-sum-equals-k`](problems/subarray-sum-equals-k/README.md).

## When to Use

- Multiple range-sum (or range-XOR) queries on a static array.
- Subarray-sum/count problems where a brute force would recompute sums for
  overlapping ranges repeatedly.

## When NOT to Use

- The array is frequently updated between queries — use a Fenwick Tree or
  Segment Tree instead, which support O(log n) updates alongside O(log n)
  range queries.
- The aggregate isn't a simple associative/invertible operation (sum, XOR,
  product of non-zero numbers) — prefix sums rely on being able to
  "subtract out" a prefix, which doesn't work for operations like min/max
  (those need a Sparse Table or Segment Tree instead).

## Quick Revision

- **Static range sum**: `prefix[j+1] - prefix[i]` for `sum(arr[i..j])`.
- **Subarray sum equals K**: prefix sum + hash map of counts, seeded with
  `{0: 1}`.
- **Complexity**: O(n) preprocessing, O(1) per range query.
- **Mutable array?** Use Fenwick/Segment Tree instead — plain prefix sums
  don't support efficient updates.
- **Representative problems**: Subarray Sum Equals K (prefix sum +
  hashing), Range Sum Query - Immutable (basic prefix sum).
