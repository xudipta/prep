# Stack

## Concept

LIFO (last-in, first-out). In Go, a slice used with `append`/truncate at
the end *is* a stack — no special type is needed.

## Operations & Complexity

| Operation | Complexity |
|---|---|
| Push | O(1) amortized |
| Pop | O(1) |
| Peek | O(1) |

## When to Use

- Matching/nesting problems (balanced parentheses, nested structures).
- DFS with an explicit stack instead of recursion (avoids stack-overflow
  risk on very deep graphs/trees).
- **Monotonic stack**: maintaining a stack that's always increasing or
  decreasing lets you answer "next greater/smaller element" queries in
  O(n) total, since each element is pushed and popped at most once.

## Go Reference

```go
stack := []int{}
stack = append(stack, x) // push
top := stack[len(stack)-1] // peek
stack = stack[:len(stack)-1] // pop

// Monotonic decreasing stack — e.g., "next greater element":
func nextGreater(nums []int) []int {
    result := make([]int, len(nums))
    for i := range result {
        result[i] = -1
    }
    var stack []int // holds indices; nums[stack[i]] decreasing bottom-to-top
    for i, v := range nums {
        for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[top] = v
        }
        stack = append(stack, i)
    }
    return result
}
```

## Common Interview Questions

- Valid Parentheses (matching via a stack of expected closers).
- Evaluate Reverse Polish Notation.
- Next Greater Element / Daily Temperatures (monotonic stack).
- Largest Rectangle in Histogram (monotonic stack tracking bar indices).
- Implement a queue using two stacks (and vice versa).
