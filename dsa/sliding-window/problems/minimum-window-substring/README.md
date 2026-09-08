# Problem: Minimum Window Substring

## Problem Statement

Given strings `s` and `t`, find the minimum-length substring of `s` that
contains every character of `t` (including duplicates — if `t` has two
`'a'`s, the window must contain at least two `'a'`s). Return `""` if no such
substring exists.

## Difficulty

Hard

## Technique

Sliding Window (variable size, minimize)

## Problem Type

String

## Key Insight

Grow the window until it contains all of `t` (validity), then shrink it from
the left as much as possible while it *stays* valid, recording the minimum
window at each valid state. A `need` count map (character → required count)
and a running `have` counter (how many distinct required characters are
currently fully satisfied) let you check "is the window valid?" in O(1)
instead of comparing full frequency maps.

## How to Recognize This Pattern

- "Minimum/shortest substring/subarray containing/satisfying..." is the
  minimize-variant window signal (as opposed to "longest," which grows and
  shrinks only to fix violations).
- Multiple required elements with counts (not just "contains character X")
  points to needing a need/have frequency-map approach rather than a single
  boolean check.

## Approach 1 — Brute Force

### Idea
Check every substring of `s`, verify it contains all of `t`'s characters
with sufficient counts, and track the shortest valid one.

### Algorithm
1. For each `(i, j)` pair, build a frequency count of `s[i:j]` and compare
   against `t`'s frequency count.

### Complexity
Time: O(n³) (or O(n² · alphabet) with smarter counting)
Space: O(alphabet size)

## Approach 2 — Optimized (Sliding Window with Need/Have Counters)

### Idea
Build `need[c]` = required count of each character in `t`. Expand `right`,
decrementing `need[c]` for the incoming character and incrementing `have`
whenever a character's remaining need reaches exactly zero (fully
satisfied). Once `have == len(need)` (all distinct required characters
satisfied), shrink `left` while the window stays fully valid, updating the
best window each time before it becomes invalid again.

### Algorithm
1. Build `need` from `t`; `required := len(need)` (distinct chars needed).
2. `left := 0`, `have := 0`, `windowCounts := map[byte]int{}`.
3. For `right` from `0` to `len(s)-1`:
   - `c := s[right]`; `windowCounts[c]++`.
   - If `c` is in `need` and `windowCounts[c] == need[c]`, `have++`.
   - While `have == required`:
     - Update best window if `right-left+1` is smaller than the current best.
     - `d := s[left]`; if `d` is in `need` and `windowCounts[d] == need[d]`,
       `have--` (about to become invalid by removing `d`).
     - `windowCounts[d]--`; `left++`.
4. Return the recorded best window, or `""` if none was found.

### Complexity
Time: O(|s| + |t|)
Space: O(alphabet size in t)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`s = "ADOBECODEBANC"`, `t = "ABC"` → `need = {A:1, B:1, C:1}`, `required = 3`.

- Expand `right` until the window `"ADOBEC"` (indices 0-5) satisfies all of
  A, B, C (`have == 3`).
- Shrink `left`: removing `'A'` at index 0 would break the A requirement, so
  the current window `"ADOBEC"` (length 6) is recorded, then shrinking
  stops being possible without breaking validity... actually shrinking
  continues while still valid: `left` moves from 0; window `"ADOBEC"` is
  valid, shrink to `"DOBEC"` — no longer has 'A', `have` drops to 2, stop
  shrinking.
- Continue expanding `right`; eventually the window `"CODEBA"` fails to
  reach validity as compactly, but later `"BANC"` (indices 9-12) becomes
  valid and is shorter (length 4).
- Final answer: `"BANC"`.

## Edge Cases

- `t` longer than `s` → impossible, return `""`.
- `t` has duplicate characters → `need` counts must be respected exactly
  (not just "character present").
- Empty `t` → conventionally return `""` (or define as trivially satisfied
  by the empty string — state your assumption in an interview).
- No valid window exists → return `""`.

## Common Mistakes

- Comparing `windowCounts[c] >= need[c]` incorrectly when updating `have`
  (must trigger exactly when the count transitions to satisfied, not every
  time it's already satisfied — otherwise `have` over-counts).
- Forgetting to decrement `have` when shrinking removes a required
  character below its needed count.
- Not restricting the `have`/`need` bookkeeping to characters actually in
  `t` (extra characters in `s` should still be counted in the window but
  shouldn't affect `have`).

## Interview Follow-ups

- What if you need the minimum window covering a *set* of arbitrary
  substrings, not just characters? Substantially harder — usually not
  solvable with this exact template.
- Return all minimal windows of that length, not just one.

## Related Problems

- Longest Substring Without Repeating Characters (variable window, maximize)
- Permutation in String / Find All Anagrams in a String (fixed-size window
  variant of the same need/have idea)
