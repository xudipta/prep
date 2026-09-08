// Package threesum solves: find all unique triplets in an array that sum
// to zero.
package threesum

import "sort"

// ThreeSum returns all unique triplets [a, b, c] from nums such that
// a + b + c == 0. Runs in O(n^2) time and O(1) extra space (beyond sorting
// and the output) by sorting nums and, for each anchor index, running
// converging two pointers over the remainder for a Two-Sum-style search.
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate anchors
		}
		if nums[i] > 0 {
			break // smallest element positive: no triplet can sum to zero
		}

		lo, hi := i+1, n-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			switch {
			case sum == 0:
				result = append(result, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			case sum < 0:
				lo++
			default:
				hi--
			}
		}
	}
	return result
}
