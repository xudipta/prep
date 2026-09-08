// Package rotatedsearch solves: find a target's index in a sorted array
// that has been rotated at an unknown pivot.
package rotatedsearch

// Search returns the index of target in nums, which is a sorted array of
// distinct integers rotated at an unknown pivot, or -1 if not present.
// Runs in O(log n) time and O(1) space: at every step at least one half of
// the current range is properly sorted, so checking which half is sorted
// and whether target's value falls in its range lets binary search proceed
// as usual.
func Search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}

		if nums[lo] <= nums[mid] { // left half [lo, mid] is sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half [mid, hi] is sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
