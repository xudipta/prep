// Package houserobber solves: find the maximum sum obtainable from an
// array without picking two adjacent elements.
package houserobber

// Rob returns the maximum amount obtainable from nums without robbing two
// adjacent houses. Runs in O(n) time and O(1) space: dp[i], the best result
// using houses 0..i, depends only on dp[i-1] and dp[i-2], so two rolling
// variables replace a full table.
func Rob(nums []int) int {
	prev2, prev1 := 0, 0 // best result considering no houses yet

	for _, v := range nums {
		curr := max(prev1, v+prev2)
		prev2, prev1 = prev1, curr
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
