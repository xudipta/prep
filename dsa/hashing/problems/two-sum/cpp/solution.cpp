#include <cassert>
#include <iostream>
#include <unordered_map>
#include <utility>
#include <vector>

using namespace std;

// Returns the indices of two elements in nums that sum to target, and
// false if no such pair exists. Runs in O(n) time and O(n) space by
// recording each value's index in a map and checking for its complement
// before inserting, so an element never pairs with itself.
pair<vector<int>, bool> twoSum(const vector<int>& nums, int target) {
    unordered_map<int, int> seen;
    seen.reserve(nums.size());
    for (int i = 0; i < (int)nums.size(); i++) {
        auto it = seen.find(target - nums[i]);
        if (it != seen.end()) {
            return {{it->second, i}, true};
        }
        seen[nums[i]] = i;
    }
    return {{}, false};
}

int main() {
    {
        auto [idx, ok] = twoSum({2, 7, 11, 15}, 9);
        assert(ok && idx == vector<int>({0, 1}));
    }
    {
        auto [idx, ok] = twoSum({3, 3}, 6);
        assert(ok && idx == vector<int>({0, 1}));
    }
    {
        auto [idx, ok] = twoSum({-3, 4, 3, 90}, 0);
        assert(ok && idx == vector<int>({0, 2}));
    }
    {
        auto [idx, ok] = twoSum({1, 2, 3}, 100);
        assert(!ok);
    }
    cout << "All tests passed\n";
    return 0;
}
