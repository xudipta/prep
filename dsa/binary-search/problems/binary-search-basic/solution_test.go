package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found in middle", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"not found", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"empty array", []int{}, 5, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 3, -1},
		{"first element", []int{2, 4, 6, 8}, 2, 0},
		{"last element", []int{2, 4, 6, 8}, 8, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.nums, tt.target); got != tt.want {
				t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
