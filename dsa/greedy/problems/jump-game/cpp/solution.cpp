#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Reports whether the last index of nums is reachable starting from
// index 0. Runs in O(n) time and O(1) space by greedily tracking the
// furthest index reachable so far: extending that frontier as far as
// possible at each step can never hurt a future decision, so no
// backtracking is needed.
bool canJump(const vector<int>& nums) {
    int furthest = 0;
    for (int i = 0; i < (int)nums.size(); i++) {
        if (i > furthest) return false; // index i is unreachable
        furthest = max(furthest, i + nums[i]);
    }
    return true;
}

int main() {
    assert(canJump({2, 3, 1, 1, 4}) == true);
    assert(canJump({3, 2, 1, 0, 4}) == false);
    assert(canJump({0}) == true);
    assert(canJump({0, 1}) == false);
    assert(canJump({1, 1, 1, 1}) == true);
    cout << "All tests passed\n";
    return 0;
}
