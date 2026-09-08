// Package singlenumber solves: find the element that appears once in an
// array where every other element appears exactly twice.
package singlenumber

// SingleNumber returns the element that appears exactly once in nums,
// where every other element appears exactly twice. Runs in O(n) time and
// O(1) space by XOR-ing every element together: XOR is commutative,
// associative, and self-canceling (x^x=0, x^0=x), so every duplicated
// pair cancels out and only the unique value remains.
func SingleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
