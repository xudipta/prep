package maxsumwindow

import "testing"

func TestMaxSumSubarray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		want   int
		wantOK bool
	}{
		{"classic example", []int{2, 1, 5, 1, 3, 2}, 3, 9, true},
		{"k equals length", []int{1, 2, 3}, 3, 6, true},
		{"k larger than length", []int{1, 2}, 3, 0, false},
		{"negative numbers", []int{-1, -2, -3, -4}, 2, -3, true},
		{"single element window", []int{4, -1, 7, 2}, 1, 7, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := MaxSumSubarray(tt.nums, tt.k)
			if ok != tt.wantOK || (ok && got != tt.want) {
				t.Errorf("MaxSumSubarray(%v, %d) = (%d, %v), want (%d, %v)",
					tt.nums, tt.k, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
