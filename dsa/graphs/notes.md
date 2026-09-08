# Graphs — Quick Revision

- **Representations**: adjacency list (sparse, most common), adjacency
  matrix (dense/O(1) edge check), implicit (grids, state spaces).
- **BFS**: queue, level order, shortest path on unweighted graphs, O(V+E).
- **DFS**: recursion/stack, components, cycle detection, topological sort,
  O(V+E).
- **Cycle detection**: undirected = visited-but-not-parent; directed =
  node currently on the recursion stack.
- **Topological sort**: Kahn's algorithm (repeatedly remove in-degree-0
  nodes); incomplete order ⇒ cycle exists.
- **Shortest path ladder**: BFS (unweighted) → Dijkstra (weighted,
  non-negative, min-heap) → Bellman-Ford (negative edges OK, detects
  negative cycles) → Floyd-Warshall (all pairs, O(V³)).
- **MST**: Kruskal (sort edges + DSU) or Prim (grow from a node + heap).
- **Union-Find**: dynamic connectivity, cycle detection in Kruskal's.
- **Watch for**: marking visited before recursing, disconnected components,
  directed vs. undirected cycle-detection logic.
- **Representative problems**: Number of Islands (grid BFS/DFS), Course
  Schedule (topological sort / cycle detection).
