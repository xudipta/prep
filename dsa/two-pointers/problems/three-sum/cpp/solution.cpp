#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns all unique triplets [a, b, c] from nums such that a+b+c == 0.
// Runs in O(n^2) time and O(1) extra space (beyond sorting and the
// output) by sorting nums and, for each anchor index, running converging
// two pointers over the remainder for a Two-Sum-style search.
vector<vector<int>> threeSum(vector<int> nums) {
    sort(nums.begin(), nums.end());
    int n = (int)nums.size();
    vector<vector<int>> result;

    for (int i = 0; i < n - 2; i++) {
        if (i > 0 && nums[i] == nums[i - 1]) continue; // skip duplicate anchors
        if (nums[i] > 0) break; // smallest element positive: no triplet sums to zero

        int lo = i + 1, hi = n - 1;
        while (lo < hi) {
            int sum = nums[i] + nums[lo] + nums[hi];
            if (sum == 0) {
                result.push_back({nums[i], nums[lo], nums[hi]});
                lo++;
                hi--;
                while (lo < hi && nums[lo] == nums[lo - 1]) lo++;
                while (lo < hi && nums[hi] == nums[hi + 1]) hi--;
            } else if (sum < 0) {
                lo++;
            } else {
                hi--;
            }
        }
    }
    return result;
}

vector<vector<int>> normalize(vector<vector<int>> triplets) {
    sort(triplets.begin(), triplets.end());
    return triplets;
}

int main() {
    {
        auto got = normalize(threeSum({-1, 0, 1, 2, -1, -4}));
        auto want = normalize({{-1, -1, 2}, {-1, 0, 1}});
        assert(got == want);
    }
    {
        auto got = normalize(threeSum({0, 0, 0, 0}));
        auto want = normalize({{0, 0, 0}});
        assert(got == want);
    }
    assert(threeSum({1, 2, -2, -1}).empty());
    assert(threeSum({0, 1}).empty());
    cout << "All tests passed\n";
    return 0;
}
