// Package longestconsecutive solves: find the length of the longest run of
// consecutive integers in an unsorted slice.
package longestconsecutive

// LongestConsecutive returns the length of the longest run of consecutive
// integers present in nums. Runs in O(n) time and O(n) space: every number
// is placed in a set, and only numbers that start a run (no predecessor in
// the set) trigger a forward scan, so the total work across all scans is
// bounded by n.
func LongestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}

	best := 0
	for v := range set {
		if _, hasPredecessor := set[v-1]; hasPredecessor {
			continue // not a sequence head; will be counted from its head
		}
		length := 1
		for {
			if _, ok := set[v+length]; !ok {
				break
			}
			length++
		}
		if length > best {
			best = length
		}
	}
	return best
}
