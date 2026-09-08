# Sliding Window

## Definition

Sliding Window maintains a contiguous range `[left, right]` over an array or
string and slides it across the input, adding an element at `right` and
removing one at `left`, to avoid recomputing overlapping work that a naive
"recompute every subarray from scratch" approach would repeat.

## Core Intuition

If you need something about every contiguous subarray/substring (a sum, a
count of distinct characters, "does it satisfy some constraint"), and
extending or shrinking the range by one element lets you update your running
answer in O(1) rather than recomputing from scratch, you can slide instead
of restarting — turning an O(n²) or O(n·k) brute force into O(n).

## Recognition Patterns

- The problem talks about **contiguous** subarrays/substrings (not
  arbitrary subsequences).
- Keywords: "longest/shortest substring/subarray such that...", "at most K
  distinct...", "sum equals/at least/at most K".
- A brute force would recompute a per-subarray value that only changes
  incrementally when the window's edges move.

## Generic Template

**Fixed-size window** (size `k` given):

```go
windowSum := 0
for i := 0; i < k; i++ {
    windowSum += arr[i]
}
best := windowSum
for right := k; right < len(arr); right++ {
    windowSum += arr[right] - arr[right-k]
    if windowSum > best {
        best = windowSum
    }
}
```

**Variable-size window** (grow to explore, shrink to restore a constraint):

```go
left := 0
state := newWindowState()
for right := 0; right < len(arr); right++ {
    state.add(arr[right])
    for !state.isValid() { // constraint violated — shrink
        state.remove(arr[left])
        left++
    }
    updateAnswer(right - left + 1) // or track state while valid
}
```

Some variable-window problems shrink to the point of *minimality* instead of
just restoring validity (e.g., Minimum Window Substring) — the shrink loop
condition there is "while still valid, keep shrinking and recording", not
"while invalid, shrink to fix."

## Variations

- **Fixed window**: window length is given (Max Sum Subarray of Size K).
- **Variable window, maximize**: grow while valid, shrink only when
  invalid (Longest Substring Without Repeating Characters).
- **Variable window, minimize**: grow until valid, then shrink as much as
  possible while it stays valid, recording the minimum along the way
  (Minimum Window Substring).
- **At-most-K transform**: "exactly K" is often computed as
  `atMost(K) - atMost(K-1)`.

## Complexity

O(n) time — each element enters and leaves the window at most once, so the
`left` and `right` pointers each traverse the array once. O(1) to O(k) space
depending on what the window's auxiliary state needs to hold (a count map,
a frequency array, etc.).

## Common Mistakes

- Shrinking with the wrong condition (fixing "invalid" vs. "still valid but
  could be smaller") — mixing these up between max/min variants.
- Off-by-one when computing window length: `right - left + 1`, not
  `right - left`.
- Not correctly decrementing/removing counts when `left` advances, leading
  to stale state.
- Using a fixed-size window template on a variable-size problem (or vice
  versa).

## Interview Tips

- Say explicitly whether the window is fixed or variable size before coding.
- For variable windows, state the invariant: "the window `[left, right]`
  always satisfies/violates constraint X."
- Mention that each pointer moves at most `n` times total, which is why the
  technique is O(n) despite the nested-looking loop.

## Example

Longest Substring Without Repeating Characters: expand `right`, and whenever
a character repeats within the window, shrink `left` past the previous
occurrence. The answer is the max window size seen. See
[`problems/longest-substring-without-repeating`](problems/longest-substring-without-repeating/README.md).

## When to Use

- Contiguous subarray/substring problems where the "goodness" of a window
  can be updated incrementally as the window slides.

## When NOT to Use

- The problem is about non-contiguous subsequences (use DP instead).
- The window's validity can't be checked/updated in O(1) or O(log n) as
  elements enter/leave (then sliding buys you nothing).
- You need to consider *all* subarrays' values, not just find an
  optimum/count meeting a threshold (e.g., "sum of all subarray sums" is
  better solved with prefix sums / combinatorics).

## Quick Revision

- **Fixed window**: precompute the first window, then slide: `+= arr[right],
  -= arr[right-k]`.
- **Variable window (max)**: grow right; on violation, shrink left until
  valid again; update answer with current window size.
- **Variable window (min)**: grow right until valid; then shrink left while
  still valid, updating the minimum each step.
- **Complexity**: O(n) — each pointer moves forward only, total O(n) steps.
- **State**: usually a frequency map/array or a running count/sum.
- **Watch for**: `right-left+1` for length, correct grow/shrink direction.
