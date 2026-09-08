package minwindow

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"classic example", "ADOBECODEBANC", "ABC", "BANC"},
		{"no valid window", "a", "aa", ""},
		{"single character match", "a", "a", "a"},
		{"t longer than s", "a", "ab", ""},
		{"whole string required", "abc", "cba", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("MinWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
