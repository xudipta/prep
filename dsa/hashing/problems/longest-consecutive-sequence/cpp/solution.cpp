#include <algorithm>
#include <cassert>
#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

// Returns the length of the longest run of consecutive integers present
// in nums. Runs in O(n) time and O(n) space: every number is placed in a
// set, and only numbers that start a run (no predecessor in the set)
// trigger a forward scan, so the total work across all scans is bounded
// by n.
int longestConsecutive(const vector<int>& nums) {
    unordered_set<int> set(nums.begin(), nums.end());
    int best = 0;
    for (int v : set) {
        if (set.count(v - 1)) continue; // not a sequence head
        int length = 1;
        while (set.count(v + length)) length++;
        best = max(best, length);
    }
    return best;
}

int main() {
    assert(longestConsecutive({100, 4, 200, 1, 3, 2}) == 4);
    assert(longestConsecutive({1, 2, 0, 1}) == 3);
    assert(longestConsecutive({}) == 0);
    assert(longestConsecutive({5, 5, 5}) == 1);
    assert(longestConsecutive({7}) == 1);
    assert(longestConsecutive({-1, -2, -3, 0}) == 4);
    cout << "All tests passed\n";
    return 0;
}
