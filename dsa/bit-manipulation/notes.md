# Bit Manipulation — Quick Revision

- **`n & (n-1)`**: clears the lowest set bit — popcount loops, power-of-two
  checks (`n != 0 && n&(n-1) == 0`).
- **`n & (-n)`**: isolates the lowest set bit.
- **XOR self-cancellation**: `x ^ x = 0`, `x ^ 0 = x` → find the single
  unique element among duplicates in O(n) time, O(1) space.
- **Bitmask as a set**: small bounded universe (≤ ~20-30 items) → represent
  a subset as bits in an int; O(1) union/intersection/membership; also the
  state for bitmask DP.
- **Set/clear/toggle/read bit i**: `n|(1<<i)`, `n&^(1<<i)`, `n^(1<<i)`,
  `(n>>i)&1`.
- **Go gotcha**: `>>` on a signed `int` is arithmetic (sign-preserving);
  use `uint` for logical shift semantics.
- **Representative problems**: Single Number (XOR), Counting Bits
  (`n & (n-1)` + DP).
