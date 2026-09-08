package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
		wantOK bool
	}{
		{"classic example", []int{2, 7, 11, 15}, 9, []int{0, 1}, true},
		{"duplicate values pairing", []int{3, 3}, 6, []int{0, 1}, true},
		{"negative numbers", []int{-3, 4, 3, 90}, 0, []int{0, 2}, true},
		{"no solution", []int{1, 2, 3}, 100, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := TwoSum(tt.nums, tt.target)
			if ok != tt.wantOK {
				t.Fatalf("TwoSum(%v, %d) ok = %v, want %v", tt.nums, tt.target, ok, tt.wantOK)
			}
			if ok && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
