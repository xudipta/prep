package lcabst

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	// Fixture BST:
	//
	//         6
	//       /   \
	//      2     8
	//     / \   / \
	//    0   4 7   9
	//       / \
	//      3   5
	n0 := &TreeNode{Val: 0}
	n3 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4, Left: n3, Right: n5}
	n2 := &TreeNode{Val: 2, Left: n0, Right: n4}
	n7 := &TreeNode{Val: 7}
	n9 := &TreeNode{Val: 9}
	n8 := &TreeNode{Val: 8, Left: n7, Right: n9}
	root := &TreeNode{Val: 6, Left: n2, Right: n8}

	tests := []struct {
		name string
		p, q *TreeNode
		want *TreeNode
	}{
		{"split at root", n2, n8, root},
		{"one is ancestor of the other", n2, n4, n2},
		{"deep nodes sharing grandparent", n0, n3, n2},
		{"same node", n0, n0, n0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowestCommonAncestor(root, tt.p, tt.q); got != tt.want {
				t.Errorf("LowestCommonAncestor(p=%d, q=%d) = %v, want %v",
					tt.p.Val, tt.q.Val, got.Val, tt.want.Val)
			}
		})
	}
}
