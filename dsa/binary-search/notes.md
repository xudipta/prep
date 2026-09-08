# Binary Search — Quick Revision

- **Idea**: halve a search space each step using a monotonic predicate;
  O(log n).
- **Classic template**: `lo <= hi`; `mid := lo + (hi-lo)/2`; compare, move
  `lo = mid+1` or `hi = mid-1`; `-1` if not found.
- **Bound-finding template**: `lo < hi`; `if feasible(mid) { hi = mid } else
  { lo = mid + 1 }`; converges with `lo == hi`.
- **Rotated array**: determine which half (`lo..mid` or `mid..hi`) is
  properly sorted, then check if target falls in that half's range.
- **Binary search on the answer**: search a value space (speed, capacity,
  time), not array indices — requires proving `feasible(x)` is monotonic
  first.
- **Complexity**: O(log n) time, O(1) space iteratively.
- **Bugs to avoid**: infinite loops (wrong mid update for the loop style
  used), off-by-one on `<=` vs `<`, assuming monotonicity without checking.
- **Representative problems**: Binary Search (Basic) (classic template),
  Search in Rotated Sorted Array (structural variant), Koko Eating Bananas
  (binary search on the answer).
