// Package lcs solves: find the length of the longest common subsequence
// of two strings.
package lcs

// LongestCommonSubsequence returns the length of the longest common
// subsequence of text1 and text2. Runs in O(n*m) time and O(m) space using
// a rolling two-row DP table: dp[i][j] depends only on the previous row
// and the current row's earlier entries, so a full 2D table isn't needed.
func LongestCommonSubsequence(text1, text2 string) int {
	n, m := len(text1), len(text2)
	prev := make([]int, m+1)

	for i := 1; i <= n; i++ {
		curr := make([]int, m+1)
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				curr[j] = 1 + prev[j-1]
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		prev = curr
	}
	return prev[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
