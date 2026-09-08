# Graph Algorithms

## Definition

A graph is a set of nodes (vertices) connected by edges. Most interview
graph problems are really "BFS/DFS on a structure that isn't obviously a
graph" — grids, dependency lists, word ladders — so the first skill is
recognizing the implicit graph, not memorizing algorithm names.

## Representation

- **Adjacency list** (`map[int][]int` or `[][]int`): the standard choice for
  sparse graphs — most interview problems. O(V+E) space.
- **Adjacency matrix** (`[][]bool` or `[][]int`): O(V²) space, O(1) edge
  lookup — useful for dense graphs or when you need fast "are u and v
  connected?" checks.
- **Implicit graph**: a grid where cell `(r, c)` connects to its 4/8
  neighbors, or a state space where "nodes" are configurations and "edges"
  are valid moves. Recognize these — they don't need an explicit adjacency
  structure, just a neighbor-generating function.

## BFS (Breadth-First Search)

**Intuition**: explore level by level using a queue. The first time you
reach a node is via a shortest path, in terms of number of edges, on an
unweighted graph.

**When to use**: shortest path on unweighted graphs, level-order processing,
"minimum number of steps/moves."

**Complexity**: O(V + E) time, O(V) space.

```go
func bfs(start int, adj map[int][]int) map[int]int {
    dist := map[int]int{start: 0}
    queue := []int{start}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        for _, next := range adj[node] {
            if _, seen := dist[next]; !seen {
                dist[next] = dist[node] + 1
                queue = append(queue, next)
            }
        }
    }
    return dist
}
```

## DFS (Depth-First Search)

**Intuition**: explore as far as possible along a branch before
backtracking, using recursion or an explicit stack.

**When to use**: connected components, cycle detection, topological sort,
exhaustive search (paired with backtracking), path existence.

**Complexity**: O(V + E) time, O(V) space (recursion stack in the worst
case).

```go
func dfs(node int, adj map[int][]int, visited map[int]bool) {
    visited[node] = true
    for _, next := range adj[node] {
        if !visited[next] {
            dfs(next, adj, visited)
        }
    }
}
```

## Connected Components

Run BFS or DFS from every unvisited node; each run discovers one component.
Count the runs, or record which component each node belongs to.

## Cycle Detection

- **Undirected graph**: during DFS, if you reach an already-visited node
  that isn't the immediate parent, there's a cycle.
- **Directed graph**: track nodes in the *current recursion stack*
  ("gray" nodes); if DFS reaches a gray node, there's a cycle. Equivalently,
  Kahn's algorithm (below) detects a cycle if not all nodes can be
  topologically ordered.

## Bipartite Check

2-color the graph via BFS/DFS: assign alternating colors as you traverse
each edge. If you ever need to assign a node a color that conflicts with an
already-assigned neighbor, the graph isn't bipartite.

## Topological Sort

**Intuition**: order nodes of a DAG so every edge `u -> v` has `u` before
`v`. Only defined for DAGs (a cycle makes topological order impossible).

**Kahn's algorithm** (BFS-based): repeatedly remove nodes with in-degree 0.

```go
func topoSort(n int, adj map[int][]int, indegree []int) ([]int, bool) {
    queue := []int{}
    for node := 0; node < n; node++ {
        if indegree[node] == 0 {
            queue = append(queue, node)
        }
    }
    order := []int{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        order = append(order, node)
        for _, next := range adj[node] {
            indegree[next]--
            if indegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return order, len(order) == n // false means a cycle exists
}
```

**When to use**: task scheduling with dependencies, build order, course
prerequisites — see [`problems/course-schedule`](problems/course-schedule/README.md).

**Complexity**: O(V + E).

## DAG DP

Once nodes are topologically ordered, you can compute `dp[node]` (e.g.,
longest path ending at `node`) by processing nodes in that order, since all
of a node's dependencies are guaranteed to be processed first.

## Union-Find / DSU

See [`data-structures/disjoint-set-union.md`](../../data-structures/disjoint-set-union.md)
for the full implementation. Used for dynamic connectivity queries and as
the core of Kruskal's MST algorithm.

## Shortest Path Algorithms

| Algorithm | Handles | Complexity | When to use |
|---|---|---|---|
| BFS | Unweighted | O(V+E) | Fewest edges/steps |
| Dijkstra | Weighted, non-negative | O((V+E) log V) with a heap | Weighted shortest path, no negative edges |
| Bellman-Ford | Weighted, negative allowed | O(V·E) | Negative edges, or need to detect negative cycles |
| Floyd-Warshall | All-pairs | O(V³) | Need shortest paths between *every* pair, V is small |

(Dijkstra/Bellman-Ford/Floyd-Warshall implementations and problems are
planned — see `PROGRESS.md`.)

## Minimum Spanning Tree

| Algorithm | Approach | Complexity |
|---|---|---|
| Kruskal | Sort edges, union-find to avoid cycles | O(E log E) |
| Prim | Grow a tree greedily from a start node, using a heap | O(E log V) |

(Planned — see `PROGRESS.md`.)

## Bridges & Articulation Points / SCCs

Advanced single-pass DFS algorithms (Tarjan's) using discovery times and
low-link values. (Planned — see `PROGRESS.md`.)

## Common Pitfalls

- Forgetting to mark a node visited *before* recursing (can cause infinite
  loops or exponential blowup on graphs with cycles).
- Using DFS recursion on a graph large enough to overflow the call stack —
  prefer an explicit stack for very deep/large graphs.
- Not handling disconnected graphs (need to iterate over all nodes to find
  every component, not just run BFS/DFS once from a fixed start).
- Confusing directed-graph cycle detection (needs recursion-stack tracking)
  with undirected-graph cycle detection (a simpler "visited but not parent"
  check).

## Quick Revision

- **BFS**: queue, level-by-level, shortest path (unweighted).
- **DFS**: stack/recursion, connected components, cycle detection,
  topological sort.
- **Topological sort**: Kahn's algorithm (BFS with in-degrees); a leftover
  non-empty queue-less node set means a cycle.
- **Cycle detection**: undirected → visited-not-parent; directed → track
  recursion-stack membership.
- **Shortest path**: BFS (unweighted) → Dijkstra (weighted, non-negative) →
  Bellman-Ford (negative edges) → Floyd-Warshall (all-pairs).
- **Complexity**: almost everything here is O(V + E).
