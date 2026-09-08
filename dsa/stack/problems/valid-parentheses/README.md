# Problem: Valid Parentheses

## Problem Statement

Given a string containing only the characters `(`, `)`, `{`, `}`, `[`, `]`,
determine if the input is valid: every opening bracket must be closed by
the same type of bracket, and in the correct order.

## Difficulty

Easy

## Technique

Stack

## Problem Type

String

## Key Insight

Every closing bracket must match the **most recently opened, still-unclosed**
bracket — exactly the LIFO behavior a stack provides. Push openers; on a
closer, check that it matches the top of the stack.

## How to Recognize This Pattern

- "Matching" or "nesting" of paired tokens is the canonical stack signal —
  the most recent unmatched opener is always what a closer must match.

## Approach 1 — Brute Force

### Idea
Repeatedly find and remove any adjacent matching pair (`()`, `{}`, `[]`)
until no more removals are possible; valid if the string becomes empty.

### Algorithm
Scan-and-replace loop until a fixed point.

### Complexity
Time: O(n²) (each removal pass can be O(n), repeated up to O(n) times)
Space: O(n)

## Approach 2 — Optimized (Stack)

### Idea
Push each opening bracket. On a closing bracket, check the stack isn't
empty and its top matches the corresponding opener; pop if so, fail
otherwise. At the end, the string is valid iff the stack is empty (no
unmatched openers remain).

### Algorithm
1. `stack := []byte{}`; `pairs := {')':'(', ']':'[', '}':'{'}`.
2. For each character `c` in the string:
   - If `c` is a closer: fail if the stack is empty or its top isn't
     `pairs[c]`; otherwise pop.
   - Else (an opener): push `c`.
3. Return `len(stack) == 0`.

### Complexity
Time: O(n)
Space: O(n)

## Sample Solution

See [`solution.go`](./solution.go).

## Dry Run

`s = "{[()]}"`

- `{`: push → `stack=[{]`.
- `[`: push → `stack=[{, []`.
- `(`: push → `stack=[{, [, (]`.
- `)`: closer, top is `(` → match, pop → `stack=[{, []`.
- `]`: closer, top is `[` → match, pop → `stack=[{]`.
- `}`: closer, top is `{` → match, pop → `stack=[]`.
- End: stack empty → `true`.

`s = "([)]"`

- `(`: push → `stack=[(]`.
- `[`: push → `stack=[(, []`.
- `)`: closer, top is `[`, expected `(` → mismatch → `false`.

## Edge Cases

- Empty string → valid (`true`) — an empty stack trivially satisfies "no
  unmatched openers."
- A lone closer with nothing open → fail immediately (stack is empty when
  a closer arrives).
- A lone opener with no closer → stack is non-empty at the end → `false`.
- Odd-length string → can never be fully valid (though the algorithm
  doesn't need this check explicitly — it falls out naturally).

## Common Mistakes

- Forgetting to check the stack is non-empty *before* peeking its top on a
  closer (would panic/index-out-of-bounds otherwise).
- Forgetting the final "stack must be empty" check — a string with only
  unmatched openers (`"((("`) would otherwise incorrectly pass.
- Using a map from opener→closer and checking membership the wrong
  direction — pick one direction (closer→opener, as above) and stay
  consistent.

## Interview Follow-ups

- Generate all valid combinations of `n` pairs of parentheses (this becomes
  a backtracking problem, not a stack-validation one — see
  `dsa/backtracking`).
- Minimum number of insertions to make a string of parentheses valid
  (a counting variant, still stack-based or counter-based).

## Related Problems

- Generate Parentheses (backtracking)
- Longest Valid Parentheses (stack of indices, or DP)
- Minimum Remove to Make Valid Parentheses
