#include <algorithm>
#include <cassert>
#include <iostream>
#include <utility>
#include <vector>

using namespace std;

// Returns the maximum sum of any contiguous subarray of length k, and
// false if nums has fewer than k elements. Runs in O(n) time and O(1)
// space by sliding a fixed-size window: each step adds the incoming
// element and removes the outgoing one instead of resumming.
pair<int, bool> maxSumSubarray(const vector<int>& nums, int k) {
    int n = (int)nums.size();
    if (k <= 0 || n < k) return {0, false};

    int windowSum = 0;
    for (int i = 0; i < k; i++) windowSum += nums[i];
    int best = windowSum;

    for (int right = k; right < n; right++) {
        windowSum += nums[right] - nums[right - k];
        best = max(best, windowSum);
    }
    return {best, true};
}

int main() {
    {
        auto [got, ok] = maxSumSubarray({2, 1, 5, 1, 3, 2}, 3);
        assert(ok && got == 9);
    }
    {
        auto [got, ok] = maxSumSubarray({1, 2, 3}, 3);
        assert(ok && got == 6);
    }
    {
        auto [got, ok] = maxSumSubarray({1, 2}, 3);
        assert(!ok);
    }
    {
        auto [got, ok] = maxSumSubarray({-1, -2, -3, -4}, 2);
        assert(ok && got == -3);
    }
    {
        auto [got, ok] = maxSumSubarray({4, -1, 7, 2}, 1);
        assert(ok && got == 7);
    }
    cout << "All tests passed\n";
    return 0;
}
