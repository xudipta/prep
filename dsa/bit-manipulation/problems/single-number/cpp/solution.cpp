#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the element that appears exactly once in nums, where every
// other element appears exactly twice. Runs in O(n) time and O(1) space
// by XOR-ing every element together: XOR is commutative, associative,
// and self-canceling (x^x=0, x^0=x), so every duplicated pair cancels
// out and only the unique value remains.
int singleNumber(const vector<int>& nums) {
    int result = 0;
    for (int v : nums) result ^= v;
    return result;
}

int main() {
    assert(singleNumber({4, 1, 2, 1, 2}) == 4);
    assert(singleNumber({7}) == 7);
    assert(singleNumber({0, 1, 1}) == 0);
    assert(singleNumber({-1, -1, -2}) == -2);
    assert(singleNumber({9, 3, 3, 5, 5}) == 9);
    cout << "All tests passed\n";
    return 0;
}
