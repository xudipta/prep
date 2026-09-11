#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the maximum total value obtainable by choosing a subset of
// items (weights[i], values[i]) with total weight at most capacity,
// using each item at most once. Runs in O(n * capacity) time and
// O(capacity) space using a single rolling row, iterated over capacity
// in descending order so each item is only applied once per pass.
int knapsack01(const vector<int>& weights, const vector<int>& values, int capacity) {
    vector<int> dp(capacity + 1, 0);

    for (size_t i = 0; i < weights.size(); i++) {
        int w = weights[i], v = values[i];
        for (int c = capacity; c >= w; c--) {
            dp[c] = max(dp[c], v + dp[c - w]);
        }
    }
    return dp[capacity];
}

int main() {
    assert(knapsack01({1, 3, 4, 5}, {1, 4, 5, 7}, 7) == 9);
    assert(knapsack01({1, 2, 3}, {10, 20, 30}, 0) == 0);
    assert(knapsack01({5}, {10}, 5) == 10);
    assert(knapsack01({10}, {100}, 5) == 0);
    assert(knapsack01({1, 2, 3}, {10, 20, 30}, 6) == 60);
    assert(knapsack01({}, {}, 10) == 0);
    cout << "All tests passed\n";
    return 0;
}
