// Package kokobananas solves: find the minimum constant eating speed that
// lets Koko finish all banana piles within h hours.
package kokobananas

// MinEatingSpeed returns the minimum integer speed k such that Koko can eat
// every pile within h hours. Runs in O(n * log(max(piles))) time and O(1)
// space by binary searching on the answer k: feasibility ("can she finish
// within h hours at speed k?") is monotonic in k, so the smallest feasible
// k is found by converging binary search rather than trying every speed.
func MinEatingSpeed(piles []int, h int) int {
	lo, hi := 1, maxPile(piles)

	for lo < hi {
		mid := lo + (hi-lo)/2
		if hoursNeeded(piles, mid) <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func hoursNeeded(piles []int, speed int) int {
	hours := 0
	for _, p := range piles {
		hours += (p + speed - 1) / speed // ceil(p / speed)
	}
	return hours
}

func maxPile(piles []int) int {
	best := 0
	for _, p := range piles {
		if p > best {
			best = p
		}
	}
	return best
}
