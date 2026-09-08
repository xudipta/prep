# Union-Find — Quick Revision

- **Answers**: "are x and y in the same group?" efficiently as groups merge
  incrementally.
- **Complexity**: O(α(n)) amortized per Find/Union with path compression +
  union by rank — effectively O(1).
- **Cycle detection**: `Union(a, b)` returning `false` means `a` and `b`
  were already connected — that edge creates a cycle.
- **Used in**: Kruskal's MST, dynamic connectivity, equivalence-grouping
  problems (Accounts Merge).
- **Not for**: static one-time connectivity checks (plain BFS/DFS is
  simpler) or relationships that get *removed* (no efficient split).
- **Full implementation**: `data-structures/disjoint-set-union.md`.
- **Representative problems**: Redundant Connection (cycle detection),
  Number of Provinces (component counting).
