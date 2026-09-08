package longestunique

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"classic example", "abcabcbb", 3},
		{"all repeats", "bbbbb", 1},
		{"mixed with a longer unique run", "pwwkew", 3},
		{"empty string", "", 0},
		{"all unique", "abcdef", 6},
		{"single character", "a", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
