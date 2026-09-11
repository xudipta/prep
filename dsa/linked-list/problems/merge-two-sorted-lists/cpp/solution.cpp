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

// Merges two sorted linked lists l1 and l2 into one sorted list and
// returns its head. Runs in O(n+m) time and O(1) extra space by splicing
// existing nodes: the smaller of the two current fronts is always the
// next node in the merged result, and once one list is exhausted the
// other's remainder (already sorted) is spliced on directly.
ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    ListNode dummy(0);
    ListNode* tail = &dummy;

    while (l1 != nullptr && l2 != nullptr) {
        if (l1->val <= l2->val) {
            tail->next = l1;
            l1 = l1->next;
        } else {
            tail->next = l2;
            l2 = l2->next;
        }
        tail = tail->next;
    }
    tail->next = (l1 != nullptr) ? l1 : l2;
    return dummy.next;
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
    assert(toVector(mergeTwoLists(buildList({1, 2, 4}), buildList({1, 3, 4}))) ==
           vector<int>({1, 1, 2, 3, 4, 4}));
    assert(toVector(mergeTwoLists(buildList({}), buildList({0}))) == vector<int>({0}));
    assert(toVector(mergeTwoLists(buildList({}), buildList({}))).empty());
    assert(toVector(mergeTwoLists(buildList({1}), buildList({2, 3, 4}))) ==
           vector<int>({1, 2, 3, 4}));
    assert(toVector(mergeTwoLists(buildList({5, 6, 7}), buildList({1}))) ==
           vector<int>({1, 5, 6, 7}));
    cout << "All tests passed\n";
    return 0;
}
