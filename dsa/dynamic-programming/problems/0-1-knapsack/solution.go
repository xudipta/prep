// Package knapsack solves the 0/1 knapsack problem: choose a subset of
// items, each usable at most once, to maximize value within a weight
// capacity.
package knapsack

// Knapsack01 returns the maximum total value obtainable by choosing a
// subset of items (weights[i], values[i]) with total weight at most
// capacity, using each item at most once. Runs in O(n * capacity) time and
// O(capacity) space using a single rolling row, iterated over capacity in
// descending order so each item is only applied once per pass.
func Knapsack01(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1)

	for i := range weights {
		w, v := weights[i], values[i]
		for c := capacity; c >= w; c-- {
			if candidate := v + dp[c-w]; candidate > dp[c] {
				dp[c] = candidate
			}
		}
	}
	return dp[capacity]
}
