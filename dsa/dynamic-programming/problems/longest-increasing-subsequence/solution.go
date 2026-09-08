// Package lis solves: find the length of the longest strictly increasing
// subsequence of an integer slice.
package lis

import "sort"

// LengthOfLIS returns the length of the longest strictly increasing
// subsequence of nums. Runs in O(n log n) time and O(n) space using
// patience sorting: tails[k] tracks the smallest possible tail value among
// all increasing subsequences of length k+1 found so far. Each new number
// either extends the longest subsequence found (appended) or improves a
// shorter subsequence's tail (replaces the first tail >= it, found via
// binary search) — the final length of tails is the answer, though tails
// itself need not be an actual subsequence of nums.
func LengthOfLIS(nums []int) int {
	var tails []int

	for _, num := range nums {
		pos := sort.Search(len(tails), func(i int) bool { return tails[i] >= num })
		if pos == len(tails) {
			tails = append(tails, num)
		} else {
			tails[pos] = num
		}
	}
	return len(tails)
}
