#include <cassert>
#include <iostream>
#include <utility>
#include <vector>

using namespace std;

// Minimal union-find with path compression and union by rank.
class DSU {
public:
    explicit DSU(int n) : parent(n), rank_(n, 0) {
        for (int i = 0; i < n; i++) parent[i] = i;
    }

    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }

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

// Returns the number of provinces (connected components) described by
// the adjacency matrix isConnected. Runs in O(n^2) time, dominated by
// scanning the matrix, using union-find to merge components: starting
// from n separate components, every successful union decreases the
// count by one.
int findCircleNum(const vector<vector<int>>& isConnected) {
    int n = (int)isConnected.size();
    DSU dsu(n);
    int provinces = n;

    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            if (isConnected[i][j] == 1 && dsu.unite(i, j)) provinces--;
        }
    }
    return provinces;
}

int main() {
    assert(findCircleNum({{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}) == 2);
    assert(findCircleNum({{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}) == 3);
    assert(findCircleNum({{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}) == 1);
    assert(findCircleNum({{1}}) == 1);
    cout << "All tests passed\n";
    return 0;
}
