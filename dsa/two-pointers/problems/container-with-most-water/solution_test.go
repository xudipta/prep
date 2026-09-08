package containerwater

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"classic example", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"two equal lines", []int{1, 1}, 1},
		{"increasing heights", []int{1, 2, 3, 4, 5}, 6},
		{"single line", []int{5}, 0},
		{"empty", []int{}, 0},
		{"all same height", []int{4, 4, 4, 4}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
