#include <cassert>
#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

// Returns the shortest substring of s containing every character of t
// (respecting duplicate counts), or "" if no such substring exists. Runs
// in O(len(s) + len(t)) time and O(alphabet size of t) space using a
// variable-size sliding window with need/have counters: the window is
// grown until valid, then shrunk from the left as far as it stays valid.
string minWindow(const string& s, const string& t) {
    if (s.empty() || t.empty() || s.size() < t.size()) return "";

    unordered_map<char, int> need;
    for (char c : t) need[c]++;
    int required = (int)need.size();

    unordered_map<char, int> windowCounts;
    int have = 0, left = 0;
    int bestLen = -1, bestStart = 0;

    for (int right = 0; right < (int)s.size(); right++) {
        char c = s[right];
        windowCounts[c]++;
        auto it = need.find(c);
        if (it != need.end() && windowCounts[c] == it->second) have++;

        while (have == required) {
            if (bestLen == -1 || right - left + 1 < bestLen) {
                bestLen = right - left + 1;
                bestStart = left;
            }
            char d = s[left];
            auto dit = need.find(d);
            if (dit != need.end() && windowCounts[d] == dit->second) have--;
            windowCounts[d]--;
            left++;
        }
    }

    if (bestLen == -1) return "";
    return s.substr(bestStart, bestLen);
}

int main() {
    assert(minWindow("ADOBECODEBANC", "ABC") == "BANC");
    assert(minWindow("a", "aa") == "");
    assert(minWindow("a", "a") == "a");
    assert(minWindow("a", "ab") == "");
    assert(minWindow("abc", "cba") == "abc");
    cout << "All tests passed\n";
    return 0;
}
