// Package groupanagrams solves: group a list of strings by anagram
// equivalence.
package groupanagrams

// GroupAnagrams groups strs into slices of mutual anagrams. Runs in
// O(n * k) time and O(n * k) space, where k is the max string length, by
// keying each string in a map with a 26-letter count signature (assumes
// lowercase English letters, the common constraint for this problem).
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var counts [26]int
		for i := 0; i < len(s); i++ {
			counts[s[i]-'a']++
		}
		groups[counts] = append(groups[counts], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
