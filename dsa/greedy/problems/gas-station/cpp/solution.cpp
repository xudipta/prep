#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns the index of the gas station to start from to travel the whole
// circuit once, or -1 if no such station exists (the problem guarantees
// the answer is unique when one exists). Runs in O(n) time and O(1)
// space: a solution exists iff total gas >= total cost, and whenever the
// running tank from a candidate start goes negative, every station up to
// that point can be ruled out at once, so the next candidate is the very
// next station.
int canCompleteCircuit(const vector<int>& gas, const vector<int>& cost) {
    int totalTank = 0, currentTank = 0, start = 0;

    for (int i = 0; i < (int)gas.size(); i++) {
        int diff = gas[i] - cost[i];
        totalTank += diff;
        currentTank += diff;

        if (currentTank < 0) {
            start = i + 1;
            currentTank = 0;
        }
    }

    return totalTank < 0 ? -1 : start;
}

int main() {
    assert(canCompleteCircuit({1, 2, 3, 4, 5}, {3, 4, 5, 1, 2}) == 3);
    assert(canCompleteCircuit({2, 3, 4}, {3, 4, 3}) == -1);
    assert(canCompleteCircuit({5}, {5}) == 0);
    assert(canCompleteCircuit({4}, {5}) == -1);
    assert(canCompleteCircuit({3, 1, 1}, {1, 2, 2}) == 0);
    cout << "All tests passed\n";
    return 0;
}
