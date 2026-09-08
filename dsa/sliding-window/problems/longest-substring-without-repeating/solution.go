// Package longestunique solves: find the length of the longest substring
// of s that contains no repeating characters.
package longestunique

// LengthOfLongestSubstring returns the length of the longest substring of s
// without repeating characters. Runs in O(n) time and O(min(n, alphabet
// size)) space using a variable-size sliding window: right expands the
// window, and left jumps directly past a duplicate's last-seen position
// using a map instead of stepping one index at a time.
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left, best := 0, 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		if idx, ok := lastSeen[c]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[c] = right
		if windowLen := right - left + 1; windowLen > best {
			best = windowLen
		}
	}
	return best
}
