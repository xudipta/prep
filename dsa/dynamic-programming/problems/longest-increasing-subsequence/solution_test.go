package lis

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"classic example", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"strictly decreasing", []int{5, 4, 3, 2, 1}, 1},
		{"strictly increasing", []int{1, 2, 3, 4, 5}, 5},
		{"empty", []int{}, 0},
		{"single element", []int{7}, 1},
		{"with duplicates", []int{1, 3, 3, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
