package problems

// ExampleSolution returns the sum of a slice — simple placeholder used for local builds.
func ExampleSolution(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
