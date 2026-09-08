// Package numberofprovinces solves: count the connected groups of cities
// given an adjacency matrix of direct connections.
package numberofprovinces

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

// FindCircleNum returns the number of provinces (connected components)
// described by the adjacency matrix isConnected. Runs in O(n^2) time,
// dominated by scanning the matrix, using union-find to merge components:
// starting from n separate components, every successful union decreases
// the count by one.
func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	d := newDSU(n)
	provinces := n

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && d.union(i, j) {
				provinces--
			}
		}
	}
	return provinces
}
