# Linked List

## Concept

A sequence of nodes where each node points to the next (singly) and
optionally the previous (doubly). Unlike arrays, elements aren't contiguous
in memory, so there's no O(1) random access — but insertion/deletion at a
known position is O(1) (no shifting required).

## Operations & Complexity

| Operation | Complexity | Notes |
|---|---|---|
| Access by index | O(n) | Must walk from the head |
| Insert/delete at head | O(1) | |
| Insert/delete at a known node | O(1) | Given a pointer to it (O(1) for singly-list delete requires the *previous* node, or a trick — see below) |
| Search | O(n) | |

## When to Use

- Frequent insertions/deletions at arbitrary positions where you already
  hold a reference to the neighboring node.
- Implementing other structures (e.g., an LRU cache's usage-order list, a
  queue).
- Rarely the default choice otherwise — slices win for cache locality and
  random access in most practical Go code.

## Go Reference

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

// Reverse a singly linked list iteratively.
func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

// Floyd's cycle detection (slow/fast pointers).
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

Go's standard library also provides `container/list`, a doubly linked list,
for cases where you don't want to hand-roll node management.

## Common Interview Questions

- Reverse a linked list (iteratively and recursively).
- Detect a cycle, and find where it begins (Floyd's algorithm).
- Merge two sorted linked lists.
- Find the middle node in one pass (slow/fast pointers).
- Delete a node given only a pointer to it (no access to the head/previous
  node) — copy the next node's value into the current node, then skip it;
  doesn't work for the tail node, a good edge case to raise.
