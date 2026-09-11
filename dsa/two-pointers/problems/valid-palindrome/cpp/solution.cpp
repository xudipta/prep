#include <cassert>
#include <cctype>
#include <iostream>
#include <string>

using namespace std;

bool isAlnum(char c) { return isalnum((unsigned char)c); }
char toLowerChar(char c) { return (char)tolower((unsigned char)c); }

// Reports whether s is a palindrome, ignoring case and non-alphanumeric
// characters. Runs in O(n) time and O(1) extra space by scanning two
// pointers inward.
bool isPalindrome(const string& s) {
    int lo = 0, hi = (int)s.size() - 1;

    while (lo < hi) {
        while (lo < hi && !isAlnum(s[lo])) lo++;
        while (lo < hi && !isAlnum(s[hi])) hi--;
        if (lo < hi) {
            if (toLowerChar(s[lo]) != toLowerChar(s[hi])) return false;
            lo++;
            hi--;
        }
    }
    return true;
}

int main() {
    assert(isPalindrome("A man, a plan, a canal: Panama") == true);
    assert(isPalindrome("race a car") == false);
    assert(isPalindrome("") == true);
    assert(isPalindrome(".,!?") == true);
    assert(isPalindrome("a") == true);
    assert(isPalindrome("Aa") == true);
    assert(isPalindrome("0P") == false);
    assert(isPalindrome("12321") == true);
    cout << "All tests passed\n";
    return 0;
}
