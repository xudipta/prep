// Package mergetwolists solves: merge two sorted linked lists into one
// sorted list.
package mergetwolists

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two sorted linked lists l1 and l2 into one sorted
// list and returns its head. Runs in O(n+m) time and O(1) extra space by
// splicing existing nodes: the smaller of the two current fronts is always
// the next node in the merged result, and once one list is exhausted the
// other's remainder (already sorted) is spliced on directly.
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}
	return dummy.Next
}
