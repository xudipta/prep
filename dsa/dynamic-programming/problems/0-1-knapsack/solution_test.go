package knapsack

import "testing"

func TestKnapsack01(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{"classic example", []int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7, 9},
		{"zero capacity", []int{1, 2, 3}, []int{10, 20, 30}, 0, 0},
		{"single item fits", []int{5}, []int{10}, 5, 10},
		{"single item too heavy", []int{10}, []int{100}, 5, 0},
		{"take everything", []int{1, 2, 3}, []int{10, 20, 30}, 6, 60},
		{"no items", []int{}, []int{}, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Knapsack01(tt.weights, tt.values, tt.capacity); got != tt.want {
				t.Errorf("Knapsack01(%v, %v, %d) = %d, want %d",
					tt.weights, tt.values, tt.capacity, got, tt.want)
			}
		})
	}
}
