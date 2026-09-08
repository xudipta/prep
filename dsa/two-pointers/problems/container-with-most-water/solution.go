// Package containerwater solves the "Container With Most Water" problem:
// find two lines that, with the x-axis, form the container holding the
// most water.
package containerwater

// MaxArea returns the maximum area formed by any two lines in height.
// Runs in O(n) time and O(1) space using converging two pointers: the
// shorter line always bounds the area, so it is always safe (and only
// useful) to move the shorter side inward.
func MaxArea(height []int) int {
	lo, hi := 0, len(height)-1
	best := 0

	for lo < hi {
		h := min(height[lo], height[hi])
		if area := (hi - lo) * h; area > best {
			best = area
		}
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return best
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
