#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the index of target in nums, which is a sorted array of
// distinct integers rotated at an unknown pivot, or -1 if not present.
// Runs in O(log n) time and O(1) space: at every step at least one half
// of the current range is properly sorted, so checking which half is
// sorted and whether target's value falls in its range lets binary
// search proceed as usual.
int search(const vector<int>& nums, int target) {
    int lo = 0, hi = (int)nums.size() - 1;

    while (lo <= hi) {
        int mid = lo + (hi - lo) / 2;
        if (nums[mid] == target) return mid;

        if (nums[lo] <= nums[mid]) { // left half [lo, mid] is sorted
            if (nums[lo] <= target && target < nums[mid]) hi = mid - 1;
            else lo = mid + 1;
        } else { // right half [mid, hi] is sorted
            if (nums[mid] < target && target <= nums[hi]) lo = mid + 1;
            else hi = mid - 1;
        }
    }
    return -1;
}

int main() {
    assert(search({4, 5, 6, 7, 0, 1, 2}, 0) == 4);
    assert(search({4, 5, 6, 7, 0, 1, 2}, 3) == -1);
    assert(search({1, 2, 3, 4, 5}, 3) == 2);
    assert(search({1}, 1) == 0);
    assert(search({1}, 0) == -1);
    assert(search({5, 1, 3}, 5) == 0);
    assert(search({6, 7, 0, 1, 2, 4, 5}, 0) == 2);
    cout << "All tests passed\n";
    return 0;
}
