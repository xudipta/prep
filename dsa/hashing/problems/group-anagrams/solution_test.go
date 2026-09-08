package groupanagrams

import (
	"reflect"
	"sort"
	"testing"
)

// normalize sorts each group's contents and then sorts the groups by their
// first element, so map-order randomness doesn't affect comparison.
func normalize(groups [][]string) [][]string {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "classic example",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "single empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "no anagrams at all",
			strs: []string{"abc", "def"},
			want: [][]string{{"abc"}, {"def"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(GroupAnagrams(tt.strs))
			want := normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strs, got, want)
			}
		})
	}
}
