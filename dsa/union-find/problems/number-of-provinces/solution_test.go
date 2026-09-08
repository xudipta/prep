package numberofprovinces

import "testing"

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name: "two provinces",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "all isolated",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
		{
			name: "all connected",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 1,
		},
		{
			name:        "single city",
			isConnected: [][]int{{1}},
			want:        1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCircleNum(tt.isConnected); got != tt.want {
				t.Errorf("FindCircleNum(%v) = %d, want %d", tt.isConnected, got, tt.want)
			}
		})
	}
}
