#include <cassert>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

// Binary tree node.
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    explicit TreeNode(int v, TreeNode* l = nullptr, TreeNode* r = nullptr)
        : val(v), left(l), right(r) {}
};

// Returns the values of root's tree grouped by level, top to bottom, left
// to right. Runs in O(n) time and O(w) space (w = the tree's maximum
// width) using iterative BFS: the queue's length is snapshotted before
// processing each level, since every node enqueued during that
// processing belongs to the next level.
vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> result;
    if (root == nullptr) return result;

    queue<TreeNode*> q;
    q.push(root);

    while (!q.empty()) {
        int levelSize = (int)q.size();
        vector<int> level;
        level.reserve(levelSize);

        for (int i = 0; i < levelSize; i++) {
            TreeNode* node = q.front();
            q.pop();
            level.push_back(node->val);
            if (node->left != nullptr) q.push(node->left);
            if (node->right != nullptr) q.push(node->right);
        }
        result.push_back(level);
    }
    return result;
}

int main() {
    assert(levelOrder(nullptr).empty());

    {
        TreeNode root(1);
        assert(levelOrder(&root) == vector<vector<int>>({{1}}));
    }
    {
        //     3
        //    /  \.
        //   9   20
        //      /  \.
        //    15    7
        TreeNode n15(15), n7(7);
        TreeNode n20(20, &n15, &n7);
        TreeNode n9(9);
        TreeNode root(3, &n9, &n20);
        vector<vector<int>> want = {{3}, {9, 20}, {15, 7}};
        assert(levelOrder(&root) == want);
    }
    {
        TreeNode n3(3);
        TreeNode n2(2, &n3, nullptr);
        TreeNode n1(1, &n2, nullptr);
        vector<vector<int>> want = {{1}, {2}, {3}};
        assert(levelOrder(&n1) == want);
    }

    cout << "All tests passed\n";
    return 0;
}
