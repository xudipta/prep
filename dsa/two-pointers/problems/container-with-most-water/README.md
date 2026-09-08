# Problem: Container With Most Water

## Problem Statement

Given an array `height` where `height[i]` is the height of a vertical line at
position `i`, find two lines that, together with the x-axis, form a
container holding the most water. Return the maximum area.

Area between lines `i` and `j` (`i < j`) is `(j - i) * min(height[i], height[j])`.

## Difficulty

Medium

## Technique

Two Pointers

## Problem Type

Array

## Key Insight

The area is limited by the **shorter** of the two lines. If you have a
candidate pair `(lo, hi)`, moving the taller line inward can only keep the
width smaller with the same or smaller limiting height — it can never help.
Moving the shorter line inward is the only move that could possibly find a
taller limiting line to compensate for the reduced width.

## How to Recognize This Pattern

- "Two ends of an array/structure" + "maximize/minimize a function of both
  ends" is a strong converging-two-pointers signal.
- The brute force is checking every pair, which is O(n²) — look for a reason
  one direction of movement is always safe to discard.

## Approach 1 — Brute Force

### Idea
Check every pair `(i, j)` and track the maximum area.

### Algorithm
1. For each `i` from `0` to `n-1`, for each `j` from `i+1` to `n-1`, compute
   the area and update the max.

### Complexity
Time: O(n²)
Space: O(1)

## Approach 2 — Optimized (Two Pointers)

### Idea
Start with the widest possible container (`lo = 0`, `hi = n-1`). Repeatedly
compute the area, update the max, then move the pointer at the **shorter**
line inward (moving the taller one can never improve the area, since width
only decreases and the limiting height can only stay the same or decrease).

### Algorithm
1. `lo := 0`, `hi := len(height) - 1`, `best := 0`.
2. While `lo < hi`:
   - `area := (hi - lo) * min(height[lo], height[hi])`; update `best`.
   - If `height[lo] < height[hi]`, `lo++`; else `hi--`.
3. Return `best`.

### Complexity
Time: O(n)
Space: O(1)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`height = [1, 8, 6, 2, 5, 4, 8, 3, 7]`

- `lo=0 (h=1), hi=8 (h=7)`: area = `8 * 1 = 8`; `height[lo] < height[hi]` → `lo++`.
- `lo=1 (h=8), hi=8 (h=7)`: area = `7 * 7 = 49`; `height[hi] < height[lo]` → `hi--`.
- `lo=1 (h=8), hi=7 (h=3)`: area = `6 * 3 = 18`; `hi--`.
- ... continues; the maximum found is `49`.

## Edge Cases

- Fewer than 2 lines → area is 0 (loop body never executes if `n < 2`).
- All lines the same height → area is determined purely by width, and the
  pointers still converge correctly to the widest pair.
- Two lines total → only one possible area, returned immediately.

## Common Mistakes

- Moving both pointers every iteration (loses potentially better answers).
- Moving the taller pointer instead of the shorter one (this discards the
  only direction that could improve the answer).
- Forgetting `min(height[lo], height[hi])` and using `height[lo]` or
  `height[hi]` unconditionally.

## Interview Follow-ups

- Prove why moving the shorter line is always safe (exchange argument: any
  container using the current shorter line's index at any future `lo`/`hi`
  position can't beat what's already been recorded, since width has already
  shrunk without the height limit improving).
- What changes if you need the *k* largest distinct areas rather than just
  the max? (No longer a pure two-pointer problem — would need to track more
  state.)

## Related Problems

- Trapping Rain Water (also two pointers, but tracks running max from both sides)
- 3Sum / Two Sum II (converging two pointers with a sum target instead of area)
