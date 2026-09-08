// Package minwindow solves: find the minimum-length substring of s that
// contains every character of t (counting duplicates).
package minwindow

// MinWindow returns the shortest substring of s containing every character
// of t (respecting duplicate counts), or "" if no such substring exists.
// Runs in O(len(s) + len(t)) time and O(alphabet size of t) space using a
// variable-size sliding window with need/have counters: the window is
// grown until valid, then shrunk from the left as far as it stays valid.
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	required := len(need)

	windowCounts := make(map[byte]int)
	have := 0
	left := 0
	bestLen := -1
	bestStart := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		windowCounts[c]++
		if n, ok := need[c]; ok && windowCounts[c] == n {
			have++
		}

		for have == required {
			if bestLen == -1 || right-left+1 < bestLen {
				bestLen = right - left + 1
				bestStart = left
			}
			d := s[left]
			if n, ok := need[d]; ok && windowCounts[d] == n {
				have--
			}
			windowCounts[d]--
			left++
		}
	}

	if bestLen == -1 {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}
