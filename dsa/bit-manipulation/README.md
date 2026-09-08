# Bit Manipulation

## Definition

Bit manipulation solves problems by operating directly on a number's binary
representation using bitwise operators (`&`, `|`, `^`, `~`, `<<`, `>>`),
often achieving O(1) per-operation cost and O(1) extra space for problems
that would otherwise need an array or hash set.

## Core Intuition

Many "which elements are present/duplicated/missing" problems over a small,
bounded value range can be solved by encoding presence as individual bits
in an integer (a **bitmask**) instead of a hash set — trading a data
structure for arithmetic. Similarly, XOR's self-canceling property
(`x ^ x == 0`, `x ^ 0 == x`) turns "find the unique/odd-one-out element"
problems into a single pass with O(1) space, where a hashing approach would
need O(n).

## Recognition Patterns

- "Find the element that appears once/an odd number of times, given all
  others appear an even number of times" — XOR signal.
- Small, fixed universe of items (≤ ~20-30) where a subset needs to be
  tracked — bitmask signal (also see `dsa/dynamic-programming` bitmask DP).
- "Count set bits," "power of two," "single bit operations" — direct
  bit-trick territory.
- Need O(1) space where a hash set would otherwise be used, and the value
  range is small/bounded.

## Generic Toolkit

```go
n & (n-1)        // clears the lowest set bit — used to count set bits, check power of 2
n & (-n)         // isolates the lowest set bit
n | (1 << i)     // set bit i
n &^ (1 << i)    // clear bit i (Go's AND-NOT operator)
n ^ (1 << i)     // toggle bit i
(n >> i) & 1     // read bit i
n & (n-1) == 0   // true iff n is a power of two (and n != 0)
a ^ b ^ b == a   // XOR self-cancellation — the basis of "find the unique element"
```

## Variations

- **XOR for uniqueness**: find the single non-duplicated element when every
  other element appears an even number of times.
- **Bitmask as a set**: represent a subset of up to ~20-30 items as bits in
  an integer — enables O(1) union/intersection/membership and is the state
  representation for bitmask DP.
- **Counting set bits**: `n & (n-1)` clears the lowest set bit each
  iteration, giving an O(popcount) loop instead of checking every bit
  position; alternatively, build a DP table using the relation
  `bits(n) = bits(n >> 1) + (n & 1)`.
- **Bit tricks for arithmetic**: multiply/divide by powers of two via
  shifts, check sign bits, swap without a temp variable (`a ^= b; b ^= a;
  a ^= b` — clever but rarely clearer than a temp variable in production
  code).

## Complexity

O(1) per bitwise operation; O(n) to process n elements with O(1) work
each; O(32) or O(64) for per-bit loops over a fixed-width integer
(effectively O(1) since word size is constant).

## Common Mistakes

- Sign-extension surprises with right shift (`>>`) on signed integers in
  some languages — Go's `>>` on a signed `int` is an arithmetic shift
  (preserves sign), which can surprise you if you expect a logical shift;
  use an unsigned type (`uint`) if you need logical shift semantics.
- Off-by-one on bit indices (bit `0` is the least significant bit).
- Forgetting operator precedence — bitwise operators have lower precedence
  than comparison operators in some languages (not Go specifically, but
  worth double-checking with parentheses regardless, for readability).
- Reaching for bit tricks where a plain, readable boolean/arithmetic
  expression would be equally fast and much clearer — bit tricks should
  earn their keep via a real, checked performance or space requirement.

## Interview Tips

- State the invariant behind the trick you're using (e.g., "since every
  element but one appears an even number of times, XOR-ing everything
  cancels the duplicates and leaves the unique one").
- Mention the O(1) space win explicitly when replacing a hash set with a
  bitmask/XOR trick — that's usually the actual point of the technique.

## Example

Single Number: XOR every element together. Every value that appears twice
cancels itself out (`x ^ x = 0`); the only value left after processing the
whole array is the one that appeared once. See
[`problems/single-number`](problems/single-number/README.md).

## When to Use

- Small, bounded value ranges where set-like state fits in an integer's
  bits.
- "Find the unique/odd-count element" problems solvable via XOR's
  self-cancellation.
- When O(1) space is explicitly required and a hashing/array-based
  solution would use O(n).

## When NOT to Use

- The value range is large/unbounded and doesn't fit meaningfully into a
  fixed-width bitmask.
- A bit trick would only save a constant factor at a real cost to
  readability, with no stated space/performance constraint that requires
  it.

## Quick Revision

- **`n & (n-1)`**: clears the lowest set bit — used for popcount and
  power-of-two checks.
- **XOR self-cancellation**: `x ^ x = 0`, `x ^ 0 = x` — the basis for
  "find the unique element" in O(n) time, O(1) space.
- **Bitmask as a set**: represent small, bounded subsets as integer bits —
  O(1) membership/union/intersection; also the state shape for bitmask DP.
- **Go specifics**: `&^` is AND-NOT (bit clear); `>>` on signed ints is
  arithmetic (sign-preserving) — use `uint` for logical shifts.
- **Representative problems**: Single Number (XOR), Counting Bits (DP +
  bit trick).
