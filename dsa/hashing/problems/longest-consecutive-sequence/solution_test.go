package longestconsecutive

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"classic example", []int{100, 4, 200, 1, 3, 2}, 4},
		{"with duplicates", []int{1, 2, 0, 1}, 3},
		{"empty", []int{}, 0},
		{"all identical", []int{5, 5, 5}, 1},
		{"single element", []int{7}, 1},
		{"negative numbers", []int{-1, -2, -3, 0}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.nums); got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
