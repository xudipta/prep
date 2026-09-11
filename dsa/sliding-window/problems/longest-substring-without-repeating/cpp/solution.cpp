#include <algorithm>
#include <cassert>
#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

// Returns the length of the longest substring of s without repeating
// characters. Runs in O(n) time and O(min(n, alphabet size)) space using
// a variable-size sliding window: right expands the window, and left
// jumps directly past a duplicate's last-seen position using a map
// instead of stepping one index at a time.
int lengthOfLongestSubstring(const string& s) {
    unordered_map<char, int> lastSeen;
    int left = 0, best = 0;

    for (int right = 0; right < (int)s.size(); right++) {
        char c = s[right];
        auto it = lastSeen.find(c);
        if (it != lastSeen.end() && it->second >= left) {
            left = it->second + 1;
        }
        lastSeen[c] = right;
        best = max(best, right - left + 1);
    }
    return best;
}

int main() {
    assert(lengthOfLongestSubstring("abcabcbb") == 3);
    assert(lengthOfLongestSubstring("bbbbb") == 1);
    assert(lengthOfLongestSubstring("pwwkew") == 3);
    assert(lengthOfLongestSubstring("") == 0);
    assert(lengthOfLongestSubstring("abcdef") == 6);
    assert(lengthOfLongestSubstring("a") == 1);
    cout << "All tests passed\n";
    return 0;
}
