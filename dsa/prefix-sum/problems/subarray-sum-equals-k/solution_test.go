package subarraysumk

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"classic example", []int{1, 2, 3}, 3, 2},
		{"single element equals k", []int{1, 1, 1}, 2, 2},
		{"with negatives", []int{1, -1, 0}, 0, 3},
		{"empty array", []int{}, 0, 0},
		{"no matching subarray", []int{1, 2, 3}, 100, 0},
		{"k is zero with zeros", []int{0, 0, 0}, 0, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
