# Linked List

## Definition

A sequence of nodes connected by pointers rather than contiguous memory
(see `data-structures/linked-list.md` for the structural overview and
complexity table). This technique section focuses on the recurring
*algorithmic* patterns used to manipulate linked lists in interviews.

## Core Intuition

Almost every linked-list problem is one of a small number of pointer-
manipulation patterns: reversing links, running two pointers at different
speeds, or using a **dummy head** node to avoid special-casing operations
on the real head. Because there's no random access, the "trick" is almost
always about *how many pointers you track simultaneously* and *in what
order you rewire `.Next` fields* — get the order wrong and you lose access
to the rest of the list.

## Recognition Patterns

- "Reverse," "reorder," or "rearrange" a linked list in place.
- Cycle detection/finding a specific node from the end — a strong
  fast/slow (Floyd's) pointer signal.
- "Merge" two or more sorted lists — a signal for a pointer-per-list
  comparison (optionally combined with a heap for *k* lists).
- Any operation that might affect the head node itself (delete the head,
  insert before the head) — a strong **dummy head** signal, so the real
  head doesn't need special-case handling.

## Generic Templates

**Reverse (iterative)**:

```go
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
```

**Fast/slow pointers** (find the middle, or detect a cycle):

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}
// slow is now at the middle (or meeting point, for cycle detection)
```

**Dummy head** (simplifies operations that might change the head):

```go
dummy := &ListNode{Next: head}
curr := dummy
// ... manipulate curr.Next freely without special-casing the real head ...
return dummy.Next
```

**Merge two sorted lists**:

```go
dummy := &ListNode{}
tail := dummy
for l1 != nil && l2 != nil {
    if l1.Val <= l2.Val {
        tail.Next, l1 = l1, l1.Next
    } else {
        tail.Next, l2 = l2, l2.Next
    }
    tail = tail.Next
}
if l1 != nil {
    tail.Next = l1
} else {
    tail.Next = l2
}
return dummy.Next
```

## Variations

- **In-place reversal** (whole list, or a sub-range `[left, right]`).
- **Fast/slow pointers**: middle-finding, cycle detection (Floyd's), "nth
  node from the end" (offset the fast pointer by n steps first).
- **Dummy head**: any insert/delete that might target the head.
- **Merge**: two sorted lists (direct comparison) or k sorted lists (heap
  of current heads — see `data-structures/heap.md`).

## Complexity

O(n) time for virtually all single-pass linked-list algorithms; O(1) extra
space for pointer manipulation (as opposed to, say, copying values into an
array first, which would be O(n) space).

## Common Mistakes

- Losing the reference to the rest of the list by overwriting `.Next`
  before saving it to a temporary variable (always save `next := node.Next`
  before reassigning `node.Next`).
- Off-by-one when using a fast pointer offset for "nth from the end"
  problems.
- Forgetting a dummy head, leading to awkward special-casing of head
  mutations (and easy-to-miss bugs when the head itself needs to change).
- Not handling `nil`/single-node lists as edge cases.

## Interview Tips

- Draw the list and pointer movements on paper/whiteboard before coding —
  linked-list bugs are almost always about get the *order* of pointer
  reassignment wrong.
- State the dummy-head trick explicitly when it applies; it signals you
  know how to avoid special-casing the head.
- For fast/slow pointer problems, state the loop invariant (e.g., "when
  `fast` reaches the end, `slow` is at the midpoint") before coding.

## Example

Reverse Linked List: track `prev`, `curr`, and (via a local variable)
`next`, rewiring one link at a time while advancing all three pointers.
See [`problems/reverse-linked-list`](problems/reverse-linked-list/README.md).

## When to Use

- The problem is inherently about a linked list, or a linked-list-like
  structure (e.g., a "next pointer" graph).

## When NOT to Use

- If random access or frequent lookups by value are needed, a linked list
  (without an auxiliary index) is the wrong structure — see
  `data-structures/array.md` or `data-structures/hash-map.md` instead.

## Quick Revision

- **Reverse**: `prev, curr, next` — save `next` before overwriting
  `curr.Next`.
- **Fast/slow**: middle-finding, cycle detection; fast moves 2x speed of
  slow.
- **Dummy head**: eliminates head-mutation special cases.
- **Merge**: pointer-per-list comparison (two lists) or a heap (k lists).
- **Complexity**: O(n) time, O(1) extra space for pointer-only
  manipulation.
- **Representative problems**: Reverse Linked List, Merge Two Sorted
  Lists.
