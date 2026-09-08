# Prefix Sum — Quick Revision

- **Idea**: precompute cumulative sums once; answer any range-sum query in
  O(1) via `prefix[j+1] - prefix[i]`.
- **Indexing convention**: `prefix[0] = 0`, `prefix[i] = sum(arr[0..i-1])` —
  state this explicitly to avoid off-by-one bugs.
- **Subarray sum equals K**: prefix sum + hash map of `{prefixSum: count}`,
  seeded with `{0: 1}`; for each new prefix sum, add
  `count[currentSum - k]`.
- **Variants**: 2D prefix sum (rectangle queries), prefix XOR.
- **When NOT to use**: array is mutated between queries (use Fenwick/
  Segment Tree instead) or the aggregate isn't invertible (min/max).
- **Representative problems**: Subarray Sum Equals K, Range Sum Query -
  Immutable.
