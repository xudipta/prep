package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero steps", 0, 1},
		{"one step", 1, 1},
		{"two steps", 2, 2},
		{"three steps", 3, 3},
		{"five steps", 5, 8},
		{"ten steps", 10, 89},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
