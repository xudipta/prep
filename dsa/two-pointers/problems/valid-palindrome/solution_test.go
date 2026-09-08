package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"classic sentence", "A man, a plan, a canal: Panama", true},
		{"not a palindrome", "race a car", false},
		{"empty string", "", true},
		{"only punctuation", ".,!?", true},
		{"single character", "a", true},
		{"mixed case letters", "Aa", true},
		{"digits matter", "0P", false},
		{"palindromic digits", "12321", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.in); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
