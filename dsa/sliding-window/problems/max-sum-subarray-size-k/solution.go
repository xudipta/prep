// Package maxsumwindow solves: find the maximum sum of any contiguous
// subarray of a fixed size k.
package maxsumwindow

// MaxSumSubarray returns the maximum sum of any contiguous subarray of
// length k, and false if nums has fewer than k elements. Runs in O(n) time
// and O(1) space by sliding a fixed-size window: each step adds the
// incoming element and removes the outgoing one instead of resumming.
func MaxSumSubarray(nums []int, k int) (int, bool) {
	n := len(nums)
	if k <= 0 || n < k {
		return 0, false
	}

	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	best := windowSum

	for right := k; right < n; right++ {
		windowSum += nums[right] - nums[right-k]
		if windowSum > best {
			best = windowSum
		}
	}
	return best, true
}
