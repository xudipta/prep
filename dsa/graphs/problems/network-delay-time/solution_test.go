package networkdelay

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][3]int
		n     int
		k     int
		want  int
	}{
		{
			name:  "classic example",
			times: [][3]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:     4,
			k:     2,
			want:  2,
		},
		{
			name:  "single node",
			times: [][3]int{},
			n:     1,
			k:     1,
			want:  0,
		},
		{
			name:  "unreachable node",
			times: [][3]int{{1, 2, 1}},
			n:     2,
			k:     2,
			want:  -1,
		},
		{
			name:  "picks the shorter of two parallel edges",
			times: [][3]int{{1, 2, 5}, {1, 2, 1}},
			n:     2,
			k:     1,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NetworkDelayTime(tt.times, tt.n, tt.k); got != tt.want {
				t.Errorf("NetworkDelayTime(%v, %d, %d) = %d, want %d",
					tt.times, tt.n, tt.k, got, tt.want)
			}
		})
	}
}
