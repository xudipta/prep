// Package gasstation solves: find the starting gas station that allows
// completing a circular route, if one exists.
package gasstation

// CanCompleteCircuit returns the index of the gas station to start from to
// travel the whole circuit once, or -1 if no such station exists (the
// problem guarantees the answer is unique when one exists). Runs in O(n)
// time and O(1) space: a solution exists iff total gas >= total cost, and
// whenever the running tank from a candidate start goes negative, every
// station up to that point can be ruled out at once, so the next
// candidate is the very next station.
func CanCompleteCircuit(gas, cost []int) int {
	totalTank, currentTank, start := 0, 0, 0

	for i := range gas {
		diff := gas[i] - cost[i]
		totalTank += diff
		currentTank += diff

		if currentTank < 0 {
			start = i + 1
			currentTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}
	return start
}
