#include <cassert>
#include <iostream>

using namespace std;

// Binary search tree node.
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    explicit TreeNode(int v, TreeNode* l = nullptr, TreeNode* r = nullptr)
        : val(v), left(l), right(r) {}
};

// Returns the lowest common ancestor of p and q in the BST rooted at
// root. Runs in O(h) time and O(1) space (h = tree height) by exploiting
// BST ordering: while both p and q are on the same side of the current
// node, descend that side; the first node where they diverge (or one of
// them equals the node) is the LCA.
TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    TreeNode* node = root;
    while (node != nullptr) {
        if (p->val < node->val && q->val < node->val) node = node->left;
        else if (p->val > node->val && q->val > node->val) node = node->right;
        else return node;
    }
    return nullptr;
}

int main() {
    //         6
    //       /    \.
    //      2      8
    //     / \.   / \.
    //    0   4  7   9
    //       / \.
    //      3   5
    TreeNode n0(0), n3(3), n5(5);
    TreeNode n4(4, &n3, &n5);
    TreeNode n2(2, &n0, &n4);
    TreeNode n7(7), n9(9);
    TreeNode n8(8, &n7, &n9);
    TreeNode root(6, &n2, &n8);

    assert(lowestCommonAncestor(&root, &n2, &n8) == &root);
    assert(lowestCommonAncestor(&root, &n2, &n4) == &n2);
    assert(lowestCommonAncestor(&root, &n0, &n3) == &n2);
    assert(lowestCommonAncestor(&root, &n0, &n0) == &n0);

    cout << "All tests passed\n";
    return 0;
}
