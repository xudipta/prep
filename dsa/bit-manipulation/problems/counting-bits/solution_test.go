package countingbits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"n=0", 0, []int{0}},
		{"n=2", 2, []int{0, 1, 1}},
		{"n=5", 5, []int{0, 1, 1, 2, 1, 2}},
		{"n=8", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBits(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
