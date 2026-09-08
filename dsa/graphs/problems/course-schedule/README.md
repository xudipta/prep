# Problem: Course Schedule

## Problem Statement

There are `numCourses` courses labeled `0` to `numCourses-1`. Given a list
of prerequisite pairs `[a, b]` meaning "to take course `a`, you must first
take course `b`," determine whether it's possible to finish all courses.

## Difficulty

Medium

## Technique

Topological Sort / Cycle Detection

## Problem Type

Graph

## Key Insight

Model courses as nodes and prerequisites as directed edges (`b -> a`). All
courses can be finished if and only if this directed graph has **no
cycle** — a cycle means a set of courses that all (indirectly) require each
other, which can never be satisfied. Kahn's algorithm detects this
naturally: if it can't produce a full ordering (some nodes never reach
in-degree 0), a cycle exists.

## How to Recognize This Pattern

- "Can all tasks be completed given dependency constraints" is the
  canonical topological-sort/cycle-detection signal.
- "Is a valid order possible" (rather than "give me the order") is
  specifically the cycle-detection framing of topological sort.

## Approach 1 — Brute Force

### Idea
DFS from each node, tracking the current recursion path; if DFS revisits a
node already in the current path, there's a cycle.

### Algorithm
Standard directed-cycle-detection DFS with three states per node: unvisited,
in-progress (on the current recursion stack), done.

### Complexity
Time: O(V + E)
Space: O(V) for recursion stack and state tracking

## Approach 2 — Optimized (Kahn's Algorithm, BFS-based)

### Idea
Compute in-degree for every node. Repeatedly remove nodes with in-degree 0
(they have no unmet prerequisites), decrementing the in-degree of their
dependents. If every node is eventually removed, no cycle exists and all
courses can be finished.

### Algorithm
1. Build adjacency list `adj[b] = [a, ...]` for every prerequisite pair
   `[a, b]`, and `indegree[a]++` for each such pair.
2. Initialize a queue with all nodes where `indegree == 0`.
3. While the queue is non-empty: pop a node, increment a `processed`
   counter, and for each neighbor, decrement its in-degree; if it reaches
   0, enqueue it.
4. Return `processed == numCourses`.

### Complexity
Time: O(V + E)
Space: O(V + E)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`numCourses = 4`, `prerequisites = [[1,0],[2,0],[3,1],[3,2]]`

- Edges: `0->1, 0->2, 1->3, 2->3`. In-degrees: `[0:0, 1:1, 2:1, 3:2]`.
- Queue starts with `[0]`.
- Pop `0`: `processed=1`; decrement `1` (→0, enqueue) and `2` (→0, enqueue).
  Queue: `[1, 2]`.
- Pop `1`: `processed=2`; decrement `3` (→1, not yet 0).
- Pop `2`: `processed=3`; decrement `3` (→0, enqueue). Queue: `[3]`.
- Pop `3`: `processed=4`.
- `processed == numCourses (4)` → return `true`.

## Edge Cases

- No prerequisites at all → trivially finishable (`true`).
- A self-loop (`[0,0]`) → node `0` never reaches in-degree 0 → `false`.
- A prerequisite pair repeated multiple times — doesn't break correctness,
  just means the in-degree is incremented more than the "true" dependency
  count, which Kahn's algorithm still handles correctly (relative
  reachability is unaffected).
- Disconnected groups of courses with no cross-dependencies — each group is
  processed independently; if any group has a cycle, the total processed
  count falls short.

## Common Mistakes

- Building the edge direction backward (`a -> b` instead of `b -> a`) —
  reverses what "prerequisite" means and produces wrong in-degrees.
- Forgetting to check `processed == numCourses` and instead just checking
  "queue became empty," which is always true even when a cycle traps some
  nodes forever at in-degree > 0.
- Not initializing in-degree correctly for nodes with no incoming edges.

## Interview Follow-ups

- Course Schedule II: return a valid course order instead of just a
  boolean — Kahn's algorithm already produces the order as a side effect
  (the sequence of popped nodes).
- What if prerequisites can have weights (e.g., a course takes time and you
  need the minimum total time to finish all courses)? Becomes a longest-path
  problem in a DAG, computed via DP over the topological order.

## Related Problems

- Course Schedule II (return the ordering)
- Alien Dictionary (topological sort derived from character ordering
  constraints)
- Minimum Height Trees
