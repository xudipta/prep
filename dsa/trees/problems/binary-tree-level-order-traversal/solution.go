// Package levelorder solves: return a binary tree's node values grouped by
// level, top to bottom.
package levelorder

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder returns the values of root's tree grouped by level, top to
// bottom, left to right. Runs in O(n) time and O(w) space (w = the tree's
// maximum width) using iterative BFS: the queue's length is snapshotted
// before processing each level, since every node enqueued during that
// processing belongs to the next level.
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
