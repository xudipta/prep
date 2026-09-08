// Package networkdelay solves: find the minimum time for a signal to
// reach every node from a source, over a directed weighted graph.
package networkdelay

import "container/heap"

type edge struct {
	to, weight int
}

// item is an entry in the priority queue: the current best known distance
// to reach node "node".
type item struct {
	node, dist int
}

// priorityQueue is a min-heap of items ordered by dist.
type priorityQueue []item

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(item)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	last := old[n-1]
	*pq = old[:n-1]
	return last
}

// NetworkDelayTime returns the minimum time for a signal sent from node k
// to reach all n nodes, or -1 if some node is unreachable. Runs in
// O((V+E) log V) time using Dijkstra's algorithm with a min-heap: since
// all weights are non-negative, the first time a node is popped from the
// heap its distance is guaranteed final.
func NetworkDelayTime(times [][3]int, n int, k int) int {
	adj := make(map[int][]edge, n)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		adj[u] = append(adj[u], edge{to: v, weight: w})
	}

	const infinity = int(^uint(0) >> 1)
	dist := make(map[int]int, n)
	for node := 1; node <= n; node++ {
		dist[node] = infinity
	}
	dist[k] = 0

	pq := &priorityQueue{{node: k, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(item)
		if curr.dist > dist[curr.node] {
			continue // stale entry; a better distance was already finalized
		}
		for _, e := range adj[curr.node] {
			if next := curr.dist + e.weight; next < dist[e.to] {
				dist[e.to] = next
				heap.Push(pq, item{node: e.to, dist: next})
			}
		}
	}

	maxDist := 0
	for node := 1; node <= n; node++ {
		if dist[node] == infinity {
			return -1
		}
		if dist[node] > maxDist {
			maxDist = dist[node]
		}
	}
	return maxDist
}
