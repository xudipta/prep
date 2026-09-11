#include <algorithm>
#include <array>
#include <cassert>
#include <climits>
#include <iostream>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

struct Edge {
    int to, weight;
};

// Returns the minimum time for a signal sent from node k to reach all n
// nodes, or -1 if some node is unreachable. Runs in O((V+E) log V) time
// using Dijkstra's algorithm with a min-heap: since all weights are
// non-negative, the first time a node is popped from the heap its
// distance is guaranteed final.
int networkDelayTime(const vector<array<int, 3>>& times, int n, int k) {
    unordered_map<int, vector<Edge>> adj;
    for (const auto& t : times) {
        adj[t[0]].push_back({t[1], t[2]});
    }

    const int infinity = INT_MAX;
    unordered_map<int, int> dist;
    for (int node = 1; node <= n; node++) dist[node] = infinity;
    dist[k] = 0;

    using Item = pair<int, int>; // (dist, node)
    priority_queue<Item, vector<Item>, greater<Item>> pq;
    pq.push({0, k});

    while (!pq.empty()) {
        auto [d, node] = pq.top();
        pq.pop();
        if (d > dist[node]) continue; // stale entry; a better distance was already finalized

        for (const Edge& e : adj[node]) {
            int next = d + e.weight;
            if (next < dist[e.to]) {
                dist[e.to] = next;
                pq.push({next, e.to});
            }
        }
    }

    int maxDist = 0;
    for (int node = 1; node <= n; node++) {
        if (dist[node] == infinity) return -1;
        maxDist = max(maxDist, dist[node]);
    }
    return maxDist;
}

int main() {
    assert(networkDelayTime({{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2) == 2);
    assert(networkDelayTime({}, 1, 1) == 0);
    assert(networkDelayTime({{1, 2, 1}}, 2, 2) == -1);
    assert(networkDelayTime({{1, 2, 5}, {1, 2, 1}}, 2, 1) == 1);
    cout << "All tests passed\n";
    return 0;
}
