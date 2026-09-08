# Linked List — Quick Revision

- **Reverse (iterative)**: `prev, curr` walk forward; save `next :=
  curr.Next` before `curr.Next = prev`.
- **Fast/slow pointers**: middle-finding, Floyd's cycle detection (`fast`
  moves 2 steps per `slow`'s 1); meeting inside a loop implies a cycle.
- **Dummy head**: `dummy := &ListNode{Next: head}` eliminates special-
  casing when the real head might be deleted/replaced.
- **Merge two sorted lists**: pointer-per-list comparison + a dummy tail.
- **Merge k sorted lists**: same idea + a min-heap over current heads.
- **Complexity**: O(n) time, O(1) extra space (pointer manipulation only).
- **Common bug**: overwriting `.Next` before saving the old value — always
  `next := node.Next` first.
- **Representative problems**: Reverse Linked List, Merge Two Sorted
  Lists.
