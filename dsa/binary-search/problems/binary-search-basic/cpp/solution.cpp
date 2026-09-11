#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the index of target in nums, or -1 if not present. Runs in
// O(log n) time and O(1) space by repeatedly halving the search range
// [lo, hi], which is always guaranteed to contain target if it exists.
int search(const vector<int>& nums, int target) {
    int lo = 0, hi = (int)nums.size() - 1;

    while (lo <= hi) {
        int mid = lo + (hi - lo) / 2;
        if (nums[mid] == target) return mid;
        if (nums[mid] < target) lo = mid + 1;
        else hi = mid - 1;
    }
    return -1;
}

int main() {
    assert(search({-1, 0, 3, 5, 9, 12}, 9) == 4);
    assert(search({-1, 0, 3, 5, 9, 12}, 2) == -1);
    assert(search({}, 5) == -1);
    assert(search({5}, 5) == 0);
    assert(search({5}, 3) == -1);
    assert(search({2, 4, 6, 8}, 2) == 0);
    assert(search({2, 4, 6, 8}, 8) == 3);
    cout << "All tests passed\n";
    return 0;
}
