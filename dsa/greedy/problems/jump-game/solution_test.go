package jumpgame

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"can reach the end", []int{2, 3, 1, 1, 4}, true},
		{"trapped by a zero", []int{3, 2, 1, 0, 4}, false},
		{"single element", []int{0}, true},
		{"zero at start blocks", []int{0, 1}, false},
		{"all ones", []int{1, 1, 1, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.nums); got != tt.want {
				t.Errorf("CanJump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
