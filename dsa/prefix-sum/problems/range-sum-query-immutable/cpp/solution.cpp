#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Precomputes prefix sums so range-sum queries run in O(1) after an O(n)
// one-time build.
class NumArray {
public:
    explicit NumArray(const vector<int>& nums) : prefix(nums.size() + 1, 0) {
        for (size_t i = 0; i < nums.size(); i++) prefix[i + 1] = prefix[i] + nums[i];
    }

    // Returns the sum of nums[i..j] (inclusive) in O(1).
    int sumRange(int i, int j) const { return prefix[j + 1] - prefix[i]; }

private:
    vector<int> prefix; // prefix[k] = sum of nums[0..k-1]
};

int main() {
    NumArray arr({-2, 0, 3, -5, 2, -1});
    assert(arr.sumRange(0, 2) == 1);
    assert(arr.sumRange(2, 5) == -1);
    assert(arr.sumRange(3, 3) == -5);
    assert(arr.sumRange(0, 5) == -3);
    cout << "All tests passed\n";
    return 0;
}
