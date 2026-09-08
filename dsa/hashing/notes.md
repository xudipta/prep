# Hashing — Quick Revision

- **Idea**: hash map/set gives O(1) average lookup/insert/membership,
  turning O(n²) linear-scan brute forces into O(n).
- **Complement pattern**: iterate once, check `seen[target-v]` before
  inserting `v` (Two Sum).
- **Canonical-key grouping**: derive a key that's identical for equivalent
  items (sorted string for anagrams), bucket by that key.
- **Set for O(1) membership**: dedup, visited-state tracking, "only start
  counting from a sequence head" tricks.
- **Prefix-sum + map**: store running prefix sums to answer "subarray sums
  to K" in O(n).
- **Complexity**: O(n) time & space.
- **Go gotchas**: map iteration order is randomized; slices can't be map
  keys (use arrays/strings); `struct{}{}` for a zero-size set value.
- **Representative problems**: Two Sum (complement), Group Anagrams
  (canonical key), Longest Consecutive Sequence (set + sequence-head trick).
