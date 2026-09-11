#include <algorithm>
#include <cassert>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

// Returns the length of the longest common subsequence of text1 and
// text2. Runs in O(n*m) time and O(m) space using a rolling two-row DP
// table: dp[i][j] depends only on the previous row and the current
// row's earlier entries, so a full 2D table isn't needed.
int longestCommonSubsequence(const string& text1, const string& text2) {
    int n = (int)text1.size(), m = (int)text2.size();
    vector<int> prev(m + 1, 0);

    for (int i = 1; i <= n; i++) {
        vector<int> curr(m + 1, 0);
        for (int j = 1; j <= m; j++) {
            if (text1[i - 1] == text2[j - 1]) curr[j] = 1 + prev[j - 1];
            else curr[j] = max(prev[j], curr[j - 1]);
        }
        prev = curr;
    }
    return prev[m];
}

int main() {
    assert(longestCommonSubsequence("abcde", "ace") == 3);
    assert(longestCommonSubsequence("abc", "def") == 0);
    assert(longestCommonSubsequence("abc", "abc") == 3);
    assert(longestCommonSubsequence("", "abc") == 0);
    assert(longestCommonSubsequence("ab", "aebc") == 2);
    cout << "All tests passed\n";
    return 0;
}
