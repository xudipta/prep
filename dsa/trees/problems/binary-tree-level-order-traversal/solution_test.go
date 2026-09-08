package levelorder

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	t.Run("nil tree", func(t *testing.T) {
		if got := LevelOrder(nil); got != nil {
			t.Errorf("LevelOrder(nil) = %v, want nil", got)
		}
	})

	t.Run("single node", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		want := [][]int{{1}}
		if got := LevelOrder(root); !reflect.DeepEqual(got, want) {
			t.Errorf("LevelOrder(single) = %v, want %v", got, want)
		}
	})

	t.Run("classic example", func(t *testing.T) {
		//     3
		//    / \
		//   9  20
		//     /  \
		//    15   7
		root := &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}
		want := [][]int{{3}, {9, 20}, {15, 7}}
		if got := LevelOrder(root); !reflect.DeepEqual(got, want) {
			t.Errorf("LevelOrder(classic) = %v, want %v", got, want)
		}
	})

	t.Run("skewed left", func(t *testing.T) {
		root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
		want := [][]int{{1}, {2}, {3}}
		if got := LevelOrder(root); !reflect.DeepEqual(got, want) {
			t.Errorf("LevelOrder(skewed) = %v, want %v", got, want)
		}
	})
}
