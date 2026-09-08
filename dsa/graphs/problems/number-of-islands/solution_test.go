package numberofislands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "classic example",
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			want: 3,
		},
		{
			name: "all water",
			grid: [][]byte{
				[]byte("000"),
				[]byte("000"),
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				[]byte("11"),
				[]byte("11"),
			},
			want: 1,
		},
		{
			name: "diagonal cells are separate islands",
			grid: [][]byte{
				[]byte("10"),
				[]byte("01"),
			},
			want: 2,
		},
		{
			name: "empty grid",
			grid: [][]byte{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIslands(tt.grid); got != tt.want {
				t.Errorf("NumIslands(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}
