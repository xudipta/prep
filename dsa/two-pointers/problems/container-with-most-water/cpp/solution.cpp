#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the maximum area formed by any two lines in height. Runs in
// O(n) time and O(1) space using converging two pointers: the shorter
// line always bounds the area, so it is always safe (and only useful) to
// move the shorter side inward.
int maxArea(const vector<int>& height) {
    int lo = 0, hi = (int)height.size() - 1;
    int best = 0;

    while (lo < hi) {
        int h = min(height[lo], height[hi]);
        best = max(best, (hi - lo) * h);
        if (height[lo] < height[hi]) lo++;
        else hi--;
    }
    return best;
}

int main() {
    assert(maxArea({1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49);
    assert(maxArea({1, 1}) == 1);
    assert(maxArea({1, 2, 3, 4, 5}) == 6);
    assert(maxArea({5}) == 0);
    assert(maxArea({}) == 0);
    assert(maxArea({4, 4, 4, 4}) == 12);
    cout << "All tests passed\n";
    return 0;
}
