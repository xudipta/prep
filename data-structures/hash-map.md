# Hash Map / Hash Set

## Concept

A hash map stores key-value pairs with average O(1) lookup, insert, and
delete via a hash function. A hash set is a hash map with no meaningful
value (just membership). Go's built-in `map[K]V` implements this directly.

## Visual Overview

```mermaid
flowchart LR
    K1["key: \"a\""] -->|hash| H1["bucket 2"]
    K2["key: \"b\""] -->|hash| H2["bucket 0"]
    K3["key: \"c\""] -->|hash| H1
    subgraph buckets ["buckets (array)"]
    direction TB
    Bk0["bucket 0: [(\"b\", 2)]"]
    Bk1["bucket 1: [ ]"]
    Bk2["bucket 2: [(\"a\", 1), (\"c\", 3)]"]
    end
    H1 -.-> Bk2
    H2 -.-> Bk0
```

A hash function maps each key to a bucket in O(1). Different keys can
collide into the same bucket (`"a"` and `"c"` above) — the bucket then
holds a short list checked in O(1) *average* time, since a good hash
function keeps collisions rare.

## Operations & Complexity

| Operation | Average | Worst Case |
|---|---|---|
| Lookup / insert / delete | O(1) | O(n) (pathological hash collisions — not a practical interview concern) |

## When to Use

See [`dsa/hashing`](../dsa/hashing/README.md) for the full pattern writeup.
In short: whenever a brute force's inner loop is "search for something
already seen," a map turns that into O(1).

## Go Reference

```go
m := make(map[string]int)
m["a"] = 1
if v, ok := m["a"]; ok {
    _ = v
}
delete(m, "a")

// Set via map[T]struct{} (zero memory overhead per value):
set := make(map[int]struct{})
set[5] = struct{}{}
_, exists := set[5]
```

**Gotchas**:
- Iteration order over a Go map is randomized by design — never depend on
  it.
- Slices, maps, and functions are not comparable and can't be used as map
  keys directly. Use an array (`[N]T`), a string, or a struct of
  comparable fields instead.
- The zero value trick: `m[k]++` works even if `k` isn't present yet,
  since Go initializes missing entries to the zero value (`0` for `int`).

## Common Interview Questions

- Two Sum (see `dsa/hashing/problems/two-sum`).
- Group Anagrams (canonical-key grouping).
- Design a data structure supporting O(1) insert/delete/getRandom (needs a
  map + array combo, not a map alone).
- LRU Cache (map + doubly linked list for O(1) get/put with eviction
  order).
