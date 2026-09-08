# Trie (Prefix Tree)

## Concept

A tree where each path from the root spells out a prefix, and each node has
up to `alphabet size` children. Words sharing a prefix share the path down
to where they diverge, making prefix-based queries fast and memory-sharing
efficient.

## Operations & Complexity

| Operation | Complexity |
|---|---|
| Insert a word of length L | O(L) |
| Search for a word of length L | O(L) |
| Check if any word has a given prefix of length L | O(L) |

(Compare to a hash set of strings, where prefix queries would require O(n)
checks against every stored string.)

## When to Use

- Autocomplete / prefix search.
- Word search / dictionary problems where many words share prefixes.
- IP routing (longest prefix match) — a specialized bit-trie.

## Go Reference

```go
type TrieNode struct {
    children [26]*TrieNode
    isWord   bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie { return &Trie{root: &TrieNode{}} }

func (t *Trie) Insert(word string) {
    node := t.root
    for _, c := range word {
        idx := c - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &TrieNode{}
        }
        node = node.children[idx]
    }
    node.isWord = true
}

func (t *Trie) Search(word string) bool {
    node := t.find(word)
    return node != nil && node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
    return t.find(prefix) != nil
}

func (t *Trie) find(s string) *TrieNode {
    node := t.root
    for _, c := range s {
        idx := c - 'a'
        if node.children[idx] == nil {
            return nil
        }
        node = node.children[idx]
    }
    return node
}
```

## Common Interview Questions

- Implement Trie (Prefix Tree) — the exercise above.
- Word Search II (Trie + backtracking DFS over a grid).
- Design Add and Search Words (Trie with wildcard support, needs DFS at
  search time instead of direct child lookup).
- Longest Common Prefix of a list of strings (can be done with a trie or
  more simply with direct comparison — good to mention both).
