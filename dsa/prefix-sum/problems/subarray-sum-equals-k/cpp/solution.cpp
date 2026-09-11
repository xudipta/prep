#include <cassert>
#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

// Returns the number of contiguous subarrays of nums summing to k. Runs
// in O(n) time and O(n) space using a running prefix sum plus a hash map
// of counts: a subarray nums[i..j-1] sums to k exactly when
// P[j] - P[i] == k, so for each new prefix sum we count how many earlier
// prefix sums equal (current - k).
int subarraySum(const vector<int>& nums, int k) {
    unordered_map<int, int> seen{{0, 1}}; // empty prefix sums to 0, seen once
    int sum = 0, count = 0;

    for (int v : nums) {
        sum += v;
        auto it = seen.find(sum - k);
        if (it != seen.end()) count += it->second;
        seen[sum]++;
    }
    return count;
}

int main() {
    assert(subarraySum({1, 2, 3}, 3) == 2);
    assert(subarraySum({1, 1, 1}, 2) == 2);
    assert(subarraySum({1, -1, 0}, 0) == 3);
    assert(subarraySum({}, 0) == 0);
    assert(subarraySum({1, 2, 3}, 100) == 0);
    assert(subarraySum({0, 0, 0}, 0) == 6);
    cout << "All tests passed\n";
    return 0;
}
