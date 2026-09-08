# Problem: Maximum Sum Subarray of Size K

## Problem Statement

Given an array of integers `nums` and an integer `k`, find the maximum sum
of any contiguous subarray of size exactly `k`.

## Difficulty

Easy

## Technique

Sliding Window (fixed size)

## Problem Type

Array

## Key Insight

Consecutive windows of size `k` overlap in `k-1` elements. Instead of
resumming each window from scratch, slide the window by adding the new
right-most element and subtracting the element that fell off the left.

## How to Recognize This Pattern

- "Subarray of size exactly K" is the canonical fixed-window signal.
- Any per-window aggregate (sum, average, max count of something) that
  updates in O(1) when one element enters and one leaves.

## Approach 1 — Brute Force

### Idea
For each starting index, sum the next `k` elements.

### Algorithm
1. For `i` from `0` to `n-k`: sum `nums[i..i+k-1]`; track the max.

### Complexity
Time: O(n·k)
Space: O(1)

## Approach 2 — Optimized (Sliding Window)

### Idea
Compute the sum of the first window. Then slide right by one each step:
add the new element, remove the one that just left the window.

### Algorithm
1. If `n < k`, no valid window exists.
2. `windowSum := sum(nums[0:k])`; `best := windowSum`.
3. For `right` from `k` to `n-1`:
   - `windowSum += nums[right] - nums[right-k]`
   - `best = max(best, windowSum)`
4. Return `best`.

### Complexity
Time: O(n)
Space: O(1)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`nums = [2, 1, 5, 1, 3, 2]`, `k = 3`

- Initial window `[2,1,5]`: sum = 8, `best = 8`.
- `right=3`: `windowSum += nums[3] - nums[0] = 1 - 2 = -1` → 7.
- `right=4`: `windowSum += nums[4] - nums[1] = 3 - 1 = 2` → 9, `best = 9`.
- `right=5`: `windowSum += nums[5] - nums[2] = 2 - 5 = -3` → 6.
- Answer: `9` (subarray `[5,1,3]`).

## Edge Cases

- `k > len(nums)` → no valid window; define a sentinel return (this
  implementation returns `0, false`).
- `k == len(nums)` → only one window, equal to the sum of the whole array.
- Negative numbers present — sliding window still works since it's just
  arithmetic, no assumption of positivity needed for a fixed-size window.

## Common Mistakes

- Recomputing the window sum from scratch each slide (defeats the purpose,
  degrades back to O(n·k)).
- Off-by-one on which element leaves the window (`right-k`, not `right-k+1`
  or `right-k-1`).
- Not handling `k > len(nums)`.

## Interview Follow-ups

- What if you need the max of a different aggregate, like the max element
  in the window (not sum)? That needs a monotonic deque (see
  Stack/Monotonic Stack notes), since max isn't O(1) to update on removal.
- Streaming version: what if `nums` arrives one element at a time and you
  need the current window sum at every step? Same sliding approach, O(1)
  update per new element.

## Related Problems

- Longest Substring Without Repeating Characters (variable window)
- Sliding Window Maximum (monotonic deque, not simple sliding-window sum)
