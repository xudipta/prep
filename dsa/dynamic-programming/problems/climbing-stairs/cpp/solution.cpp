#include <cassert>
#include <iostream>

using namespace std;

// Returns the number of distinct ways to reach step n, taking 1 or 2
// steps at a time. Runs in O(n) time and O(1) space: dp[i] depends only
// on dp[i-1] and dp[i-2], so two rolling variables replace a full table.
int climbStairs(int n) {
    if (n <= 1) return 1;

    int prev2 = 1, prev1 = 1; // dp[0], dp[1]
    for (int i = 2; i <= n; i++) {
        int curr = prev1 + prev2;
        prev2 = prev1;
        prev1 = curr;
    }
    return prev1;
}

int main() {
    assert(climbStairs(0) == 1);
    assert(climbStairs(1) == 1);
    assert(climbStairs(2) == 2);
    assert(climbStairs(3) == 3);
    assert(climbStairs(5) == 8);
    assert(climbStairs(10) == 89);
    cout << "All tests passed\n";
    return 0;
}
