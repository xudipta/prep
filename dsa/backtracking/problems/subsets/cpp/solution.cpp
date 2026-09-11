#include <algorithm>
#include <cassert>
#include <functional>
#include <iostream>
#include <vector>

using namespace std;

// Returns all possible subsets of nums (the power set). Runs in
// O(n * 2^n) time and space by backtracking through an include/exclude
// decision at every index: each element is either added to the current
// path or skipped, and every complete path is copied into the result.
vector<vector<int>> subsets(const vector<int>& nums) {
    vector<vector<int>> result;
    vector<int> path;

    function<void(int)> backtrack = [&](int index) {
        if (index == (int)nums.size()) {
            result.push_back(path);
            return;
        }

        path.push_back(nums[index]); // include nums[index]
        backtrack(index + 1);
        path.pop_back(); // undo

        backtrack(index + 1); // exclude nums[index]
    };

    backtrack(0);
    return result;
}

vector<vector<int>> normalize(vector<vector<int>> subs) {
    for (auto& s : subs) sort(s.begin(), s.end());
    sort(subs.begin(), subs.end(), [](const vector<int>& a, const vector<int>& b) {
        if (a.size() != b.size()) return a.size() < b.size();
        return a < b;
    });
    return subs;
}

int main() {
    {
        auto got = normalize(subsets({1, 2}));
        auto want = normalize({{}, {1}, {2}, {1, 2}});
        assert(got == want);
        assert(got.size() == (size_t)(1 << 2));
    }
    {
        auto got = normalize(subsets({}));
        auto want = normalize({{}});
        assert(got == want);
        assert(got.size() == 1);
    }
    {
        auto got = normalize(subsets({5}));
        auto want = normalize({{}, {5}});
        assert(got == want);
        assert(got.size() == 2);
    }
    cout << "All tests passed\n";
    return 0;
}
