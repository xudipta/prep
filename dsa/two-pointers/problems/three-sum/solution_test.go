package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(triplets [][]int) [][]int {
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
	return triplets
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "classic example",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "no triplet",
			nums: []int{1, 2, -2, -1},
			want: nil,
		},
		{
			name: "too short",
			nums: []int{0, 1},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(ThreeSum(tt.nums))
			want := normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}
