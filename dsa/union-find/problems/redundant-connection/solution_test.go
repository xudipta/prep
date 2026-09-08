package redundantconnection

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][2]int
		want  [2]int
	}{
		{
			name:  "triangle",
			edges: [][2]int{{1, 2}, {1, 3}, {2, 3}},
			want:  [2]int{2, 3},
		},
		{
			name:  "extra edge appears later",
			edges: [][2]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			want:  [2]int{1, 4},
		},
		{
			name:  "star with one redundant spoke",
			edges: [][2]int{{1, 2}, {1, 3}, {1, 4}, {2, 4}},
			want:  [2]int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRedundantConnection(tt.edges); got != tt.want {
				t.Errorf("FindRedundantConnection(%v) = %v, want %v", tt.edges, got, tt.want)
			}
		})
	}
}
