#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns a vector answer of length n+1 where answer[i] is the number of
// set bits in i. Runs in O(n) time and O(n) space using the recurrence
// answer[i] = answer[i & (i-1)] + 1: i & (i-1) clears i's lowest set
// bit, producing a strictly smaller number whose popcount was already
// computed earlier in the same pass.
vector<int> countBits(int n) {
    vector<int> answer(n + 1, 0);
    for (int i = 1; i <= n; i++) answer[i] = answer[i & (i - 1)] + 1;
    return answer;
}

int main() {
    assert(countBits(0) == vector<int>({0}));
    assert(countBits(2) == vector<int>({0, 1, 1}));
    assert(countBits(5) == vector<int>({0, 1, 1, 2, 1, 2}));
    assert(countBits(8) == vector<int>({0, 1, 1, 2, 1, 2, 2, 3, 1}));
    cout << "All tests passed\n";
    return 0;
}
