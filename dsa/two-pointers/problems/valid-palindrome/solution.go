// Package validpalindrome solves: determine whether a string is a
// palindrome once case is ignored and non-alphanumeric characters are
// skipped.
package validpalindrome

import "unicode"

// IsPalindrome reports whether s is a palindrome, ignoring case and
// non-alphanumeric characters. Runs in O(n) time and O(1) extra space by
// scanning two pointers inward over the rune slice.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	lo, hi := 0, len(runes)-1

	for lo < hi {
		for lo < hi && !isAlphanumeric(runes[lo]) {
			lo++
		}
		for lo < hi && !isAlphanumeric(runes[hi]) {
			hi--
		}
		if lo < hi {
			if unicode.ToLower(runes[lo]) != unicode.ToLower(runes[hi]) {
				return false
			}
			lo++
			hi--
		}
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
