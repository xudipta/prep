#include <algorithm>
#include <array>
#include <cassert>
#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

string signature(const string& s) {
    array<int, 26> counts{};
    for (char c : s) counts[c - 'a']++;
    string sig;
    sig.reserve(counts.size() * 2);
    for (int c : counts) {
        sig += to_string(c);
        sig += ',';
    }
    return sig;
}

// Groups strs into vectors of mutual anagrams. Runs in O(n * k) time and
// O(n * k) space, where k is the max string length, by keying each string
// in a hash map with a 26-letter count signature (assumes lowercase
// English letters, the common constraint for this problem).
vector<vector<string>> groupAnagrams(const vector<string>& strs) {
    unordered_map<string, vector<string>> groups;
    for (const string& s : strs) {
        groups[signature(s)].push_back(s);
    }

    vector<vector<string>> result;
    result.reserve(groups.size());
    for (auto& [key, group] : groups) result.push_back(move(group));
    return result;
}

vector<vector<string>> normalize(vector<vector<string>> groups) {
    for (auto& g : groups) sort(g.begin(), g.end());
    sort(groups.begin(), groups.end(), [](const vector<string>& a, const vector<string>& b) {
        return a[0] < b[0];
    });
    return groups;
}

int main() {
    {
        auto got = normalize(groupAnagrams({"eat", "tea", "tan", "ate", "nat", "bat"}));
        auto want = normalize({{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}});
        assert(got == want);
    }
    {
        auto got = groupAnagrams({});
        assert(got.empty());
    }
    {
        auto got = normalize(groupAnagrams({""}));
        auto want = normalize({{""}});
        assert(got == want);
    }
    {
        auto got = normalize(groupAnagrams({"abc", "def"}));
        auto want = normalize({{"abc"}, {"def"}});
        assert(got == want);
    }
    cout << "All tests passed\n";
    return 0;
}
