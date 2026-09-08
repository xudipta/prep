# Sliding Window — Quick Revision

- **Idea**: maintain a contiguous `[left, right]` range, update its state
  incrementally instead of recomputing per subarray.
- **Fixed window**: `sum += arr[right]; sum -= arr[right-k]`.
- **Variable, maximize**: expand `right`; while invalid, shrink `left`;
  track `max(right-left+1)`.
- **Variable, minimize**: expand `right` until valid; then shrink `left`
  while still valid, tracking the min length seen.
- **Complexity**: O(n) time (amortized — each index enters/leaves the window
  once), O(1)–O(k) space for window state.
- **Common state**: frequency map (`map[byte]int` or `[26]int`), running
  sum/count, "have vs need" counters for multi-character targets.
- **Common bugs**: wrong shrink condition, off-by-one on window length,
  forgetting to decrement counts on shrink.
- **Representative problems**: Max Sum Subarray of Size K (fixed),
  Longest Substring Without Repeating Characters (variable/maximize),
  Minimum Window Substring (variable/minimize with have/need counters).
