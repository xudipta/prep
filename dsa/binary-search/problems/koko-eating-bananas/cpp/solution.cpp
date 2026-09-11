#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

int hoursNeeded(const vector<int>& piles, int speed) {
    int hours = 0;
    for (int p : piles) hours += (p + speed - 1) / speed; // ceil(p / speed)
    return hours;
}

// Returns the minimum integer speed k such that Koko can eat every pile
// within h hours. Runs in O(n * log(max(piles))) time and O(1) space by
// binary searching on the answer k: feasibility ("can she finish within
// h hours at speed k?") is monotonic in k, so the smallest feasible k is
// found by converging binary search rather than trying every speed.
int minEatingSpeed(const vector<int>& piles, int h) {
    int lo = 1, hi = *max_element(piles.begin(), piles.end());

    while (lo < hi) {
        int mid = lo + (hi - lo) / 2;
        if (hoursNeeded(piles, mid) <= h) hi = mid;
        else lo = mid + 1;
    }
    return lo;
}

int main() {
    assert(minEatingSpeed({3, 6, 7, 11}, 8) == 4);
    assert(minEatingSpeed({30, 11, 23, 4, 20}, 5) == 30);
    assert(minEatingSpeed({30, 11, 23, 4, 20}, 6) == 23);
    assert(minEatingSpeed({5}, 5) == 1);
    assert(minEatingSpeed({1, 2, 3}, 3) == 3);
    cout << "All tests passed\n";
    return 0;
}
