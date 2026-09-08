package rangesumimmutable

import "testing"

func TestNumArraySumRange(t *testing.T) {
	arr := NewNumArray([]int{-2, 0, 3, -5, 2, -1})

	tests := []struct {
		name string
		i, j int
		want int
	}{
		{"classic range", 0, 2, 1},
		{"another range", 2, 5, -1},
		{"single element", 3, 3, -5},
		{"whole array", 0, 5, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arr.SumRange(tt.i, tt.j); got != tt.want {
				t.Errorf("SumRange(%d, %d) = %d, want %d", tt.i, tt.j, got, tt.want)
			}
		})
	}
}
