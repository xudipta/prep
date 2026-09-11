#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Returns, for each index i, the number of days until a warmer
// temperature, or 0 if none exists. Runs in O(n) time and O(n) space
// using a monotonic stack of indices with strictly decreasing
// temperatures: each new day resolves every stacked day it's warmer
// than, so every index is pushed once and popped at most once overall.
vector<int> dailyTemperatures(const vector<int>& temperatures) {
    vector<int> answer(temperatures.size(), 0);
    vector<int> stack; // indices; temperatures[stack[...]] strictly decreasing

    for (int i = 0; i < (int)temperatures.size(); i++) {
        while (!stack.empty() && temperatures[stack.back()] < temperatures[i]) {
            int top = stack.back();
            stack.pop_back();
            answer[top] = i - top;
        }
        stack.push_back(i);
    }
    return answer;
}

int main() {
    assert(dailyTemperatures({73, 74, 75, 71, 69, 72, 76, 73}) ==
           vector<int>({1, 1, 4, 2, 1, 1, 0, 0}));
    assert(dailyTemperatures({5, 4, 3, 2, 1}) == vector<int>({0, 0, 0, 0, 0}));
    assert(dailyTemperatures({1, 2, 3, 4}) == vector<int>({1, 1, 1, 0}));
    assert(dailyTemperatures({50}) == vector<int>({0}));
    assert(dailyTemperatures({60, 60, 60}) == vector<int>({0, 0, 0}));
    cout << "All tests passed\n";
    return 0;
}
