// Package validparentheses solves: determine whether a string of bracket
// characters is validly matched and nested.
package validparentheses

// IsValid reports whether s consists of validly matched and nested
// brackets. Runs in O(n) time and O(n) space using a stack: openers are
// pushed, and each closer must match the most recently pushed, still-open
// bracket — exactly the LIFO order a stack provides.
func IsValid(s string) bool {
	closerToOpener := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		if opener, isCloser := closerToOpener[c]; isCloser {
			if len(stack) == 0 || stack[len(stack)-1] != opener {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
