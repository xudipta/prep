package utils

// Sum returns the sum of a slice of ints.
func Sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

// Reverse returns a new slice with elements in reverse order.
func Reverse(nums []int) []int {
	out := make([]int, len(nums))
	for i, v := range nums {
		out[len(nums)-1-i] = v
	}
	return out
}
