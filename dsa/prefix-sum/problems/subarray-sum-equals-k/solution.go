// Package subarraysumk solves: count the number of contiguous subarrays
// whose sum equals a target k.
package subarraysumk

// SubarraySum returns the number of contiguous subarrays of nums summing
// to k. Runs in O(n) time and O(n) space using a running prefix sum plus a
// hash map of counts: a subarray nums[i..j-1] sums to k exactly when
// P[j] - P[i] == k, so for each new prefix sum we count how many earlier
// prefix sums equal (current - k).
func SubarraySum(nums []int, k int) int {
	seen := map[int]int{0: 1} // empty prefix sums to 0, seen once
	sum, count := 0, 0

	for _, v := range nums {
		sum += v
		count += seen[sum-k]
		seen[sum]++
	}
	return count
}
