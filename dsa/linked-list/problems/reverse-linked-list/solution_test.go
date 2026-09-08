package reverselist

import "testing"

// buildList creates a linked list from vals.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// toSlice collects a linked list's values into a slice for comparison.
func toSlice(head *ListNode) []int {
	var vals []int
	for node := head; node != nil; node = node.Next {
		vals = append(vals, node.Val)
	}
	return vals
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"single element", []int{1}, []int{1}},
		{"empty list", []int{}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(ReverseList(buildList(tt.in)))
			if len(got) != len(tt.want) {
				t.Fatalf("ReverseList(%v) = %v, want %v", tt.in, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ReverseList(%v) = %v, want %v", tt.in, got, tt.want)
				}
			}
		})
	}
}
