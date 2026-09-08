// Package reverselist solves: reverse a singly linked list in place.
package reverselist

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses a singly linked list and returns its new head.
// Runs in O(n) time and O(1) space by walking the list once, rewiring
// each node's Next pointer to the previously visited node instead of the
// next one, taking care to save the original next pointer before
// overwriting it.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
