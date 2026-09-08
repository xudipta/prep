// Package climbingstairs solves: count the number of distinct ways to
// climb n stairs taking 1 or 2 steps at a time.
package climbingstairs

// ClimbStairs returns the number of distinct ways to reach step n, taking
// 1 or 2 steps at a time. Runs in O(n) time and O(1) space: dp[i] depends
// only on dp[i-1] and dp[i-2], so two rolling variables replace a full
// table.
func ClimbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	prev2, prev1 := 1, 1 // dp[0], dp[1]
	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2, prev1 = prev1, curr
	}
	return prev1
}
