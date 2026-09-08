package gasstation

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{"classic example", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"impossible circuit", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"single station exact", []int{5}, []int{5}, 0},
		{"single station insufficient", []int{4}, []int{5}, -1},
		{"start at index zero already works", []int{3, 1, 1}, []int{1, 2, 2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanCompleteCircuit(tt.gas, tt.cost); got != tt.want {
				t.Errorf("CanCompleteCircuit(%v, %v) = %d, want %d", tt.gas, tt.cost, got, tt.want)
			}
		})
	}
}
