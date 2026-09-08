package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"classic example", []int{2, 7, 9, 3, 1}, 12},
		{"alternating best", []int{1, 2, 3, 1}, 4},
		{"empty", []int{}, 0},
		{"single house", []int{5}, 5},
		{"two houses", []int{2, 9}, 9},
		{"all zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rob(tt.nums); got != tt.want {
				t.Errorf("Rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
