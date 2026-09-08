// Package binarysearch solves: find the index of a target value in a
// sorted slice of distinct integers.
package binarysearch

// Search returns the index of target in nums, or -1 if not present. Runs
// in O(log n) time and O(1) space by repeatedly halving the search range
// [lo, hi], which is always guaranteed to contain target if it exists.
func Search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}
