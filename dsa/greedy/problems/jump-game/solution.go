// Package jumpgame solves: determine whether the last index of an array is
// reachable, where each element is the maximum jump length from that
// position.
package jumpgame

// CanJump reports whether the last index of nums is reachable starting
// from index 0. Runs in O(n) time and O(1) space by greedily tracking the
// furthest index reachable so far: extending that frontier as far as
// possible at each step can never hurt a future decision, so no
// backtracking is needed.
func CanJump(nums []int) bool {
	furthest := 0
	for i, jump := range nums {
		if i > furthest {
			return false // index i is unreachable
		}
		if reach := i + jump; reach > furthest {
			furthest = reach
		}
	}
	return true
}
