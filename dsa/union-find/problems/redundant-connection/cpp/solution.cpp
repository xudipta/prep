#include <array>
#include <cassert>
#include <iostream>
#include <utility>
#include <vector>

using namespace std;

// Minimal union-find with path compression and union by rank, scoped to
// this file. See data-structures/disjoint-set-union.md for the fully
// documented reference implementation.
class DSU {
public:
    explicit DSU(int n) : parent(n), rank_(n, 0) {
        for (int i = 0; i < n; i++) parent[i] = i;
    }

    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }

    // Merges the sets containing x and y, returning false if they were
    // already in the same set (meaning this edge would create a cycle).
    bool unite(int x, int y) {
        int rootX = find(x), rootY = find(y);
        if (rootX == rootY) return false;
        if (rank_[rootX] < rank_[rootY]) swap(rootX, rootY);
        parent[rootY] = rootX;
        if (rank_[rootX] == rank_[rootY]) rank_[rootX]++;
        return true;
    }

private:
    vector<int> parent;
    vector<int> rank_;
};

// Returns the edge (as {u, v}, 1-indexed) that can be removed to make
// edges describe a valid tree again. Runs in O(n * alpha(n)) time by
// union-ing edges in order; the first edge whose endpoints are already
// connected is the one that closes the graph's single cycle.
array<int, 2> findRedundantConnection(const vector<array<int, 2>>& edges) {
    int n = (int)edges.size(); // a tree with n nodes has n-1 edges; one extra is given
    DSU dsu(n + 1);

    for (const auto& edge : edges) {
        if (!dsu.unite(edge[0], edge[1])) return edge;
    }
    return {0, 0};
}

int main() {
    assert((findRedundantConnection({{1, 2}, {1, 3}, {2, 3}}) == array<int, 2>{2, 3}));
    assert((findRedundantConnection({{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}) ==
            array<int, 2>{1, 4}));
    assert((findRedundantConnection({{1, 2}, {1, 3}, {1, 4}, {2, 4}}) == array<int, 2>{2, 4}));
    cout << "All tests passed\n";
    return 0;
}
