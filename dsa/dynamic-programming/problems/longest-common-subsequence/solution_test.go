package lcs

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{"classic example", "abcde", "ace", 3},
		{"no common characters", "abc", "def", 0},
		{"identical strings", "abc", "abc", 3},
		{"one string empty", "", "abc", 0},
		{"one is subsequence of other", "ab", "aebc", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("LongestCommonSubsequence(%q, %q) = %d, want %d",
					tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}
