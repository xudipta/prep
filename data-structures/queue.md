# Queue / Deque

## Concept

**Queue**: FIFO (first-in, first-out). **Deque** (double-ended queue):
supports push/pop at both ends. A deque generalizes both a stack and a
queue.

## Operations & Complexity

| Operation | Complexity |
|---|---|
| Enqueue / push back | O(1) amortized |
| Dequeue / pop front | O(1) amortized with the right structure (see below) |
| Push/pop front (deque) | O(1) amortized |

**Caveat in Go**: `queue = queue[1:]` on a slice is O(1) *per call*
(just moves the slice header) but leaks the popped element's backing
array space until the whole slice is garbage collected, and repeated
reslicing from the front eventually needs a new backing array once
capacity from the front is exhausted — acceptable for interview-sized
inputs, but for production code prefer `container/list` or a ring buffer.

## When to Use

- **Queue**: BFS (graphs, trees), task scheduling, producer/consumer
  buffering.
- **Deque**: sliding window maximum/minimum (monotonic deque), palindrome
  checks from both ends, implementing both stack and queue behavior with
  one structure.

## Go Reference

```go
// Simple slice-backed queue (fine for BFS on interview-sized inputs):
queue := []int{start}
for len(queue) > 0 {
    node := queue[0]
    queue = queue[1:]
    // enqueue neighbors...
}

// container/list as a deque:
dq := list.New()
dq.PushBack(1)
dq.PushFront(2)
front := dq.Front().Value
dq.Remove(dq.Front())
```

**Monotonic deque** (Sliding Window Maximum): keep indices in the deque
with strictly decreasing values; the front is always the current window's
maximum.

```go
func maxSlidingWindow(nums []int, k int) []int {
    var deque []int // indices, values decreasing front-to-back
    var result []int
    for i, v := range nums {
        for len(deque) > 0 && nums[deque[len(deque)-1]] < v {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        if deque[0] <= i-k {
            deque = deque[1:]
        }
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    return result
}
```

## Common Interview Questions

- Implement a queue using two stacks (and vice versa).
- Sliding Window Maximum (monotonic deque).
- BFS-based problems generally (level order traversal, shortest path).
- Design a circular buffer / circular queue.
