#include <array>
#include <cassert>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

// Reports whether all numCourses courses can be completed given
// prerequisites, where each pair [a, b] means "b must be taken before a".
// Runs in O(V + E) time and space using Kahn's algorithm: courses can all
// be finished if and only if the prerequisite graph has no cycle, which
// Kahn's algorithm detects by checking whether every node can eventually
// reach in-degree zero.
bool canFinish(int numCourses, const vector<array<int, 2>>& prerequisites) {
    vector<vector<int>> adj(numCourses);
    vector<int> indegree(numCourses, 0);

    for (const auto& p : prerequisites) {
        int course = p[0], prereq = p[1];
        adj[prereq].push_back(course);
        indegree[course]++;
    }

    queue<int> q;
    for (int course = 0; course < numCourses; course++) {
        if (indegree[course] == 0) q.push(course);
    }

    int processed = 0;
    while (!q.empty()) {
        int course = q.front();
        q.pop();
        processed++;

        for (int next : adj[course]) {
            if (--indegree[next] == 0) q.push(next);
        }
    }

    return processed == numCourses;
}

int main() {
    assert(canFinish(2, {}) == true);
    assert(canFinish(2, {{1, 0}}) == true);
    assert(canFinish(2, {{1, 0}, {0, 1}}) == false);
    assert(canFinish(1, {{0, 0}}) == false);
    assert(canFinish(4, {{1, 0}, {2, 0}, {3, 1}, {3, 2}}) == true);
    assert(canFinish(4, {{1, 0}, {2, 1}, {0, 2}, {3, 2}}) == false);
    cout << "All tests passed\n";
    return 0;
}
