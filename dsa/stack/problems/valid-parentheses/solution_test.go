package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"simple pair", "()", true},
		{"mixed valid", "()[]{}", true},
		{"nested valid", "{[()]}", true},
		{"wrong order", "([)]", false},
		{"unmatched closer", "]", false},
		{"unmatched opener", "(((", false},
		{"empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
