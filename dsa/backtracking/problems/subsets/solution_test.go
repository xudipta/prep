package subsets

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(subsets [][]int) [][]int {
	for _, s := range subsets {
		sort.Ints(s)
	}
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := range subsets[i] {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
	return subsets
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{{}},
		},
		{
			name: "single element",
			nums: []int{5},
			want: [][]int{{}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(Subsets(tt.nums))
			want := normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Subsets(%v) = %v, want %v", tt.nums, got, want)
			}
			if len(got) != 1<<len(tt.nums) {
				t.Errorf("Subsets(%v) returned %d subsets, want %d", tt.nums, len(got), 1<<len(tt.nums))
			}
		})
	}
}
