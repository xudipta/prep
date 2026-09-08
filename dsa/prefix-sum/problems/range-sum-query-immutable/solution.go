// Package rangesumimmutable solves: answer many sum(nums[i..j]) queries
// against a static array in O(1) each.
package rangesumimmutable

// NumArray precomputes prefix sums so range-sum queries run in O(1) after
// an O(n) one-time build.
type NumArray struct {
	prefix []int // prefix[k] = sum of nums[0..k-1]
}

// NewNumArray builds a NumArray over nums in O(n) time and O(n) space.
func NewNumArray(nums []int) *NumArray {
	prefix := make([]int, len(nums)+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	return &NumArray{prefix: prefix}
}

// SumRange returns the sum of nums[i..j] (inclusive) in O(1).
func (n *NumArray) SumRange(i, j int) int {
	return n.prefix[j+1] - n.prefix[i]
}
