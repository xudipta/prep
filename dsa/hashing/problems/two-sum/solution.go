// Package twosum solves: find the indices of two numbers in an array that
// add up to a target value.
package twosum

// TwoSum returns the indices of two elements in nums that sum to target,
// and false if no such pair exists. Runs in O(n) time and O(n) space by
// recording each value's index in a map and checking for its complement
// before inserting, so an element never pairs with itself.
func TwoSum(nums []int, target int) ([]int, bool) {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}, true
		}
		seen[v] = i
	}
	return nil, false
}
