package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"classic example", []int{4, 1, 2, 1, 2}, 4},
		{"single element", []int{7}, 7},
		{"unique is zero", []int{0, 1, 1}, 0},
		{"negative numbers", []int{-1, -1, -2}, -2},
		{"unique at start", []int{9, 3, 3, 5, 5}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.nums); got != tt.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
