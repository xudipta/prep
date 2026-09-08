// Package lcabst solves: find the lowest common ancestor of two nodes in a
// binary search tree.
package lcabst

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LowestCommonAncestor returns the lowest common ancestor of p and q in
// the BST rooted at root. Runs in O(h) time and O(1) space (h = tree
// height) by exploiting BST ordering: while both p and q are on the same
// side of the current node, descend that side; the first node where they
// diverge (or one of them equals the node) is the LCA.
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root
	for node != nil {
		switch {
		case p.Val < node.Val && q.Val < node.Val:
			node = node.Left
		case p.Val > node.Val && q.Val > node.Val:
			node = node.Right
		default:
			return node
		}
	}
	return nil
}
