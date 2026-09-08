# Problem: Subarray Sum Equals K

## Problem Statement

Given an integer array `nums` and an integer `k`, return the total number of
contiguous subarrays whose sum equals `k`. Elements can be negative.

## Difficulty

Medium

## Technique

Prefix Sum + Hashing

## Problem Type

Array

## Key Insight

Let `P[i]` be the sum of the first `i` elements. The sum of subarray
`nums[i..j-1]` equals `P[j] - P[i]`. We want `P[j] - P[i] == k`, i.e.
`P[i] == P[j] - k`. So for each running prefix sum `P[j]`, count how many
earlier prefix sums equal `P[j] - k` — a complement lookup identical in
spirit to Two Sum, tracked with a running hash map instead of the whole
array up front.

## How to Recognize This Pattern

- "Number of subarrays with sum equal to K" is the canonical prefix-sum +
  hashing signal — especially since negative numbers rule out a
  sliding-window approach (window sums aren't monotonic when values can be
  negative).
- Any "count pairs of prefix values satisfying a relation" problem reduces
  to this same complement-lookup shape.

## Approach 1 — Brute Force

### Idea
Check every subarray's sum directly.

### Algorithm
For each `i`, extend `j` and accumulate the running sum, checking against
`k` at each step.

### Complexity
Time: O(n²)
Space: O(1)

## Approach 2 — Optimized (Prefix Sum + Hash Map)

### Idea
Track a running prefix sum while scanning once. Maintain a map of
`{prefixSum: count of times seen}`, seeded with `{0: 1}` (representing the
empty prefix, needed so subarrays starting at index 0 are counted
correctly). At each step, add `map[currentSum - k]` to the answer, then
record the current sum in the map.

### Algorithm
1. `sum := 0`, `count := 0`, `seen := map[int]int{0: 1}`.
2. For each `v` in `nums`:
   - `sum += v`.
   - `count += seen[sum-k]`.
   - `seen[sum]++`.
3. Return `count`.

### Complexity
Time: O(n)
Space: O(n)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`nums = [1, 2, 3]`, `k = 3`

- `seen = {0:1}`, `sum=0`, `count=0`.
- `v=1`: `sum=1`; `count += seen[1-3=-2] = 0`; `seen = {0:1, 1:1}`.
- `v=2`: `sum=3`; `count += seen[3-3=0] = 1` → `count=1`; `seen = {0:1,
  1:1, 3:1}`.
- `v=3`: `sum=6`; `count += seen[6-3=3] = 1` → `count=2`; `seen = {0:1,
  1:1, 3:1, 6:1}`.
- Answer: `2` (subarrays `[1,2]` and `[3]`).

## Edge Cases

- Empty array → `0`.
- All zeros with `k=0` → every subarray qualifies; the `{0:1}` seed and
  running counts handle this correctly without special-casing.
- Negative numbers present — this is exactly why sliding window doesn't
  apply here (a window's sum isn't monotonic as it grows when negative
  values are possible).

## Common Mistakes

- Forgetting to seed `seen` with `{0: 1}`, which silently undercounts every
  subarray that starts at index 0.
- Updating `seen[sum]++` *before* checking `seen[sum-k]`, which would
  incorrectly let a subarray "match itself" when `k == 0`.
- Reaching for sliding window out of habit — it doesn't work here because
  negative numbers break the monotonic window-growth assumption.

## Interview Follow-ups

- Return the actual subarrays, not just the count (requires storing lists
  of indices per prefix sum instead of just counts).
- Subarray Sums Divisible by K — same template, key the map by
  `sum % k` (adjusted for negative remainders) instead of the raw sum.

## Related Problems

- Two Sum (same complement-lookup shape, no prefix-sum framing)
- Continuous Subarray Sum (prefix sum mod k)
- Range Sum Query - Immutable (basic prefix sum, no hashing needed)
