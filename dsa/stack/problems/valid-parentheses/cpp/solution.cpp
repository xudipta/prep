#include <cassert>
#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

// Reports whether s consists of validly matched and nested brackets. Runs
// in O(n) time and O(n) space using a stack: openers are pushed, and each
// closer must match the most recently pushed, still-open bracket —
// exactly the LIFO order a stack provides.
bool isValid(const string& s) {
    unordered_map<char, char> closerToOpener{{')', '('}, {']', '['}, {'}', '{'}};
    vector<char> stack;

    for (char c : s) {
        auto it = closerToOpener.find(c);
        if (it != closerToOpener.end()) {
            if (stack.empty() || stack.back() != it->second) return false;
            stack.pop_back();
        } else {
            stack.push_back(c);
        }
    }
    return stack.empty();
}

int main() {
    assert(isValid("()") == true);
    assert(isValid("()[]{}") == true);
    assert(isValid("{[()]}") == true);
    assert(isValid("([)]") == false);
    assert(isValid("]") == false);
    assert(isValid("(((") == false);
    assert(isValid("") == true);
    cout << "All tests passed\n";
    return 0;
}
