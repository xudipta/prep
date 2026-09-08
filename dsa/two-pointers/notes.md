# Two Pointers — Quick Revision

- **Idea**: replace nested loops with two indices that move based on a
  monotonic decision, exploiting sorted/orderable structure.
- **Converging** (`lo=0, hi=n-1`, move inward): pair-sum problems, palindrome
  checks, container/area problems.
- **Fast/slow** (both start at 0, different speeds/conditions): in-place
  compaction, cycle detection.
- **Anchor + converge**: fix index `i`, run converging two pointers on
  `i+1..n-1` — turns 3Sum into "Two Sum on sorted array" for each `i`.
- **Complexity**: O(n) after O(n log n) sort, O(1) extra space.
- **Before you reach for it**: can you prove that discarding one pointer's
  current position never discards the optimal answer? If not, reconsider.
- **Common bugs**: forgetting duplicate-skipping, `<` vs `<=`, moving the
  wrong pointer.
- **Representative problems**: Valid Palindrome (converging + skip
  non-alphanumeric), Container With Most Water (greedy pointer move), 3Sum
  (anchor + converge + dedup).
