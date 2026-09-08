# Data Structures Cheatsheet

| Structure | Access | Search | Insert | Delete | When to use |
|---|---|---|---|---|---|
| Array/Slice | O(1) | O(n) | O(1) amortized (end) / O(n) (middle) | O(n) (middle) | Random access, known/incremental size |
| Linked List | O(n) | O(n) | O(1) (given node) | O(1) (given node) | Frequent insert/delete at known positions |
| Stack | O(1) top | O(n) | O(1) | O(1) | LIFO, matching/nesting, DFS |
| Queue/Deque | O(1) ends | O(n) | O(1) | O(1) | FIFO, BFS, sliding window max/min |
| Hash Map/Set | — | O(1) avg | O(1) avg | O(1) avg | Fast lookup/membership/grouping |
| Heap | O(1) top | O(n) | O(log n) | O(log n) top | Priority queue, top-K, Dijkstra |
| BST (balanced) | O(log n) | O(log n) | O(log n) | O(log n) | Ordered data, range queries |
| Trie | — | O(L) | O(L) | O(L) | Prefix search (L = string length) |
| DSU (Union-Find) | — | O(α(n)) | O(α(n)) union | — | Dynamic connectivity, Kruskal's MST |
| Segment Tree / Fenwick Tree | O(log n) query | — | O(log n) update | — | Range queries + point/range updates |

**Go specifics**:
- Slices are the default array-equivalent; `container/list` for a doubly
  linked list; `container/heap` for a heap (requires implementing
  `sort.Interface` + `Push`/`Pop`).
- Map iteration order is randomized — never depend on it.
- `struct{}{}` (zero-size) is the idiomatic "set value" in a
  `map[T]struct{}`.

See `data-structures/` for full notes with Go reference implementations per
structure, and `dsa/graphs` / `dsa/trees` for algorithm-level coverage of
graphs and trees specifically.
