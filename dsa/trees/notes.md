# Trees — Quick Revision

- **Preorder** (node, left, right): top-down processing, serialization.
- **Inorder** (left, node, right): sorted order on a BST.
- **Postorder** (left, right, node): bottom-up, Tree DP, height/diameter.
- **Level order (BFS)**: queue; snapshot `len(queue)` before each level's
  inner loop.
- **BST property**: left < node < right at every node → O(height) search/
  insert/LCA instead of O(n).
- **LCA (BST)**: walk down; diverge when p, q split to different sides ⇒
  current node is the LCA.
- **LCA (generic tree)**: postorder DFS, node is LCA if p/q found in
  different subtrees.
- **Diameter/path-sum**: postorder DFS returning a per-node value while
  tracking a separate global best (the best path may not pass through the
  root).
- **Complexity**: O(n) time; O(h) space for DFS recursion, O(w) for BFS
  queue.
- **Representative problems**: Binary Tree Level Order Traversal (BFS),
  Lowest Common Ancestor of a BST (ordering exploitation).
