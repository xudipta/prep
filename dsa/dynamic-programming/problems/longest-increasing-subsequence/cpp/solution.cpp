#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the length of the longest strictly increasing subsequence of
// nums. Runs in O(n log n) time and O(n) space using patience sorting:
// tails[k] tracks the smallest possible tail value among all increasing
// subsequences of length k+1 found so far. Each new number either
// extends the longest subsequence found (appended) or improves a
// shorter subsequence's tail (replaces the first tail >= it, found via
// binary search) — the final size of tails is the answer, though tails
// itself need not be an actual subsequence of nums.
int lengthOfLIS(const vector<int>& nums) {
    vector<int> tails;

    for (int num : nums) {
        auto it = lower_bound(tails.begin(), tails.end(), num);
        if (it == tails.end()) tails.push_back(num);
        else *it = num;
    }
    return (int)tails.size();
}

int main() {
    assert(lengthOfLIS({10, 9, 2, 5, 3, 7, 101, 18}) == 4);
    assert(lengthOfLIS({5, 4, 3, 2, 1}) == 1);
    assert(lengthOfLIS({1, 2, 3, 4, 5}) == 5);
    assert(lengthOfLIS({}) == 0);
    assert(lengthOfLIS({7}) == 1);
    assert(lengthOfLIS({1, 3, 3, 5}) == 3);
    cout << "All tests passed\n";
    return 0;
}
