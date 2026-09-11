#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the maximum amount obtainable from nums without robbing two
// adjacent houses. Runs in O(n) time and O(1) space: dp[i], the best
// result using houses 0..i, depends only on dp[i-1] and dp[i-2], so two
// rolling variables replace a full table.
int rob(const vector<int>& nums) {
    int prev2 = 0, prev1 = 0; // best result considering no houses yet

    for (int v : nums) {
        int curr = max(prev1, v + prev2);
        prev2 = prev1;
        prev1 = curr;
    }
    return prev1;
}

int main() {
    assert(rob({2, 7, 9, 3, 1}) == 12);
    assert(rob({1, 2, 3, 1}) == 4);
    assert(rob({}) == 0);
    assert(rob({5}) == 5);
    assert(rob({2, 9}) == 9);
    assert(rob({0, 0, 0}) == 0);
    cout << "All tests passed\n";
    return 0;
}
