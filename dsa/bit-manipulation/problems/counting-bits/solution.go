// Package countingbits solves: compute the number of set bits for every
// integer from 0 to n.
package countingbits

// CountBits returns a slice answer of length n+1 where answer[i] is the
// number of set bits in i. Runs in O(n) time and O(n) space using the
// recurrence answer[i] = answer[i&(i-1)] + 1: i&(i-1) clears i's lowest
// set bit, producing a strictly smaller number whose popcount was already
// computed earlier in the same pass.
func CountBits(n int) []int {
	answer := make([]int, n+1)
	for i := 1; i <= n; i++ {
		answer[i] = answer[i&(i-1)] + 1
	}
	return answer
}
