package mergetwolists

import "testing"

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	var vals []int
	for node := head; node != nil; node = node.Next {
		vals = append(vals, node.Val)
	}
	return vals
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		l1, l2 []int
		want   []int
	}{
		{"classic example", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"first list empty", []int{}, []int{0}, []int{0}},
		{"both empty", []int{}, []int{}, nil},
		{"second list longer", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"first list longer", []int{5, 6, 7}, []int{1}, []int{1, 5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(MergeTwoLists(buildList(tt.l1), buildList(tt.l2)))
			if len(got) != len(tt.want) {
				t.Fatalf("MergeTwoLists(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
				}
			}
		})
	}
}
