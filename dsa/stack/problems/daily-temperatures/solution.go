// Package dailytemperatures solves: for each day, find how many days must
// pass before a warmer temperature occurs.
package dailytemperatures

// DailyTemperatures returns, for each index i, the number of days until a
// warmer temperature, or 0 if none exists. Runs in O(n) time and O(n)
// space using a monotonic stack of indices with strictly decreasing
// temperatures: each new day resolves every stacked day it's warmer than,
// so every index is pushed once and popped at most once overall.
func DailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	var stack []int // indices; temperatures[stack[...]] strictly decreasing

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = i - top
		}
		stack = append(stack, i)
	}
	return answer
}
