package kokobananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"classic example", []int{3, 6, 7, 11}, 8, 4},
		{"tight deadline", []int{30, 11, 23, 4, 20}, 5, 30},
		{"generous deadline", []int{30, 11, 23, 4, 20}, 6, 23},
		{"single pile", []int{5}, 5, 1},
		{"h equals number of piles", []int{1, 2, 3}, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEatingSpeed(tt.piles, tt.h); got != tt.want {
				t.Errorf("MinEatingSpeed(%v, %d) = %d, want %d", tt.piles, tt.h, got, tt.want)
			}
		})
	}
}
