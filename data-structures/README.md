# Data Structures — Revision Notes

Concise reference notes for the data structures interviews expect you to
know cold: what they are, their operation complexities, when to reach for
them, and a minimal Go reference. For structures with dedicated problem
sets, see the linked `dsa/` technique instead (e.g., binary trees and BSTs
are covered in depth in [`dsa/trees`](../dsa/trees/README.md)).

| Structure | Notes |
|---|---|
| [Array](array.md) | Fixed-size contiguous storage; Go's dynamic equivalent is a slice |
| [String](string.md) | Immutable byte/rune sequence in Go |
| [Linked List](linked-list.md) | Singly/doubly linked nodes |
| [Stack](stack.md) | LIFO; slice-backed in Go |
| [Queue / Deque](queue.md) | FIFO / double-ended; `container/list` or slice-backed |
| [Hash Map / Hash Set](hash-map.md) | Go's built-in `map` |
| [Heap / Priority Queue](heap.md) | `container/heap` interface |
| [Trie](trie.md) | Prefix tree for string sets |
| [Graph](graph.md) | Adjacency list/matrix — see [`dsa/graphs`](../dsa/graphs/README.md) for algorithms |
| [Disjoint Set Union](disjoint-set-union.md) | Union-Find with path compression + union by rank |

Binary trees, BSTs, segment trees, and Fenwick trees are covered under
`dsa/trees` and (planned) `dsa/segment-tree` since they're more
algorithm-heavy than these reference-style notes.
