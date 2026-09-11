#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

// Singly linked list node.
struct ListNode {
    int val;
    ListNode* next;
    explicit ListNode(int v, ListNode* n = nullptr) : val(v), next(n) {}
};

// Reverses a singly linked list and returns its new head. Runs in O(n)
// time and O(1) space by walking the list once, rewiring each node's
// next pointer to the previously visited node instead of the next one,
// taking care to save the original next pointer before overwriting it.
ListNode* reverseList(ListNode* head) {
    ListNode* prev = nullptr;
    ListNode* curr = head;

    while (curr != nullptr) {
        ListNode* next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
}

ListNode* buildList(const vector<int>& vals) {
    ListNode dummy(0);
    ListNode* tail = &dummy;
    for (int v : vals) {
        tail->next = new ListNode(v);
        tail = tail->next;
    }
    return dummy.next;
}

vector<int> toVector(ListNode* head) {
    vector<int> vals;
    for (ListNode* node = head; node != nullptr; node = node->next) vals.push_back(node->val);
    return vals;
}

int main() {
    assert(toVector(reverseList(buildList({1, 2, 3, 4, 5}))) == vector<int>({5, 4, 3, 2, 1}));
    assert(toVector(reverseList(buildList({1, 2}))) == vector<int>({2, 1}));
    assert(toVector(reverseList(buildList({1}))) == vector<int>({1}));
    assert(toVector(reverseList(buildList({}))).empty());
    cout << "All tests passed\n";
    return 0;
}
