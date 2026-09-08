package rotatedsearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"classic rotated", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"target not present", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"no rotation", []int{1, 2, 3, 4, 5}, 3, 2},
		{"single element found", []int{1}, 1, 0},
		{"single element not found", []int{1}, 0, -1},
		{"rotation at last index", []int{5, 1, 3}, 5, 0},
		{"target at pivot", []int{6, 7, 0, 1, 2, 4, 5}, 0, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.nums, tt.target); got != tt.want {
				t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
