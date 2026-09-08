package nqueens

import (
	"reflect"
	"sort"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int // expected solution count
	}{
		{"n=1 trivial solution", 1, 1},
		{"n=2 no solutions", 2, 0},
		{"n=3 no solutions", 3, 0},
		{"n=4 two solutions", 4, 2},
		{"n=8 classic count", 8, 92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveNQueens(tt.n)
			if len(got) != tt.want {
				t.Errorf("SolveNQueens(%d) returned %d solutions, want %d", tt.n, len(got), tt.want)
			}
		})
	}
}

func TestSolveNQueensBoardsAreValid(t *testing.T) {
	solutions := SolveNQueens(4)
	want := [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}

	sort.Slice(solutions, func(i, j int) bool { return solutions[i][0] < solutions[j][0] })
	sort.Slice(want, func(i, j int) bool { return want[i][0] < want[j][0] })

	if !reflect.DeepEqual(solutions, want) {
		t.Errorf("SolveNQueens(4) = %v, want %v", solutions, want)
	}
}
