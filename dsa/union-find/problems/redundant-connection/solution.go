// Package redundantconnection solves: given the edges of a tree plus one
// extra edge, find the extra edge that creates the graph's single cycle.
package redundantconnection

// dsu is a minimal union-find with path compression and union by rank,
// scoped to this package. See data-structures/disjoint-set-union.md for
// the fully documented reference implementation.
type dsu struct {
	parent []int
	rank   []int
}

func newDSU(n int) *dsu {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &dsu{parent: parent, rank: make([]int, n)}
}

func (d *dsu) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

// union merges the sets containing x and y, returning false if they were
// already in the same set (meaning this edge would create a cycle).
func (d *dsu) union(x, y int) bool {
	rootX, rootY := d.find(x), d.find(y)
	if rootX == rootY {
		return false
	}
	if d.rank[rootX] < d.rank[rootY] {
		rootX, rootY = rootY, rootX
	}
	d.parent[rootY] = rootX
	if d.rank[rootX] == d.rank[rootY] {
		d.rank[rootX]++
	}
	return true
}

// FindRedundantConnection returns the edge (as [u, v], 1-indexed) that can
// be removed to make edges describe a valid tree again. Runs in O(n * α(n))
// time by union-ing edges in order; the first edge whose endpoints are
// already connected is the one that closes the graph's single cycle.
func FindRedundantConnection(edges [][2]int) [2]int {
	n := len(edges) // a tree with n nodes has n-1 edges; one extra edge is given
	d := newDSU(n + 1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !d.union(u, v) {
			return edge
		}
	}
	return [2]int{}
}
