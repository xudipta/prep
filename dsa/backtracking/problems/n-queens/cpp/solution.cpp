#include <algorithm>
#include <cassert>
#include <functional>
#include <iostream>
#include <string>
#include <unordered_set>
#include <vector>

using namespace std;

vector<string> buildBoard(const vector<int>& queenCol, int n) {
    vector<string> board(n);
    for (int row = 0; row < n; row++) {
        string line(n, '.');
        line[queenCol[row]] = 'Q';
        board[row] = line;
    }
    return board;
}

// Returns every solution to the n-queens problem, each as n strings
// representing the board ('Q' for a queen, '.' otherwise). Runs in
// pruned-exponential time (no polynomial solution is known) and O(n)
// space for bookkeeping by placing one queen per row and tracking
// occupied columns and diagonals in sets, so each placement's safety is
// checked in O(1) instead of rescanning the board.
vector<vector<string>> solveNQueens(int n) {
    unordered_set<int> cols, diag1, diag2; // diag1: row-col, diag2: row+col
    vector<int> queenCol(n); // queenCol[row] = column of the queen in that row
    vector<vector<string>> result;

    function<void(int)> place = [&](int row) {
        if (row == n) {
            result.push_back(buildBoard(queenCol, n));
            return;
        }
        for (int col = 0; col < n; col++) {
            if (cols.count(col) || diag1.count(row - col) || diag2.count(row + col)) {
                continue; // conflict; prune
            }

            cols.insert(col);
            diag1.insert(row - col);
            diag2.insert(row + col);
            queenCol[row] = col;

            place(row + 1);

            cols.erase(col);
            diag1.erase(row - col);
            diag2.erase(row + col);
        }
    };

    place(0);
    return result;
}

int main() {
    assert((int)solveNQueens(1).size() == 1);
    assert((int)solveNQueens(2).size() == 0);
    assert((int)solveNQueens(3).size() == 0);
    assert((int)solveNQueens(4).size() == 2);
    assert((int)solveNQueens(8).size() == 92);

    vector<vector<string>> solutions = solveNQueens(4);
    vector<vector<string>> want = {
        {".Q..", "...Q", "Q...", "..Q."},
        {"..Q.", "Q...", "...Q", ".Q.."},
    };
    sort(solutions.begin(), solutions.end());
    sort(want.begin(), want.end());
    assert(solutions == want);

    cout << "All tests passed\n";
    return 0;
}
