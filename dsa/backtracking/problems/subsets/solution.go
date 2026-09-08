// Package subsets solves: return every subset (the power set) of a slice
// of distinct integers.
package subsets

// Subsets returns all possible subsets of nums (the power set). Runs in
// O(n * 2^n) time and space by backtracking through an include/exclude
// decision at every index: each element is either added to the current
// path or skipped, and every complete path is copied into the result.
func Subsets(nums []int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(nums) {
			combo := make([]int, len(path))
			copy(combo, path)
			result = append(result, combo)
			return
		}

		// Include nums[index].
		path = append(path, nums[index])
		backtrack(index + 1)
		path = path[:len(path)-1] // undo

		// Exclude nums[index].
		backtrack(index + 1)
	}

	backtrack(0)
	return result
}
