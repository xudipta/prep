#include <cassert>
#include <iostream>
#include <queue>
#include <string>
#include <vector>

using namespace std;

void floodFill(const vector<string>& grid, vector<vector<bool>>& visited, int startR, int startC) {
    int rows = (int)grid.size(), cols = (int)grid[0].size();
    queue<pair<int, int>> q;
    q.push({startR, startC});
    visited[startR][startC] = true;

    const int dr[4] = {-1, 1, 0, 0};
    const int dc[4] = {0, 0, -1, 1};

    while (!q.empty()) {
        auto [r, c] = q.front();
        q.pop();
        for (int d = 0; d < 4; d++) {
            int nr = r + dr[d], nc = c + dc[d];
            if (nr < 0 || nr >= rows || nc < 0 || nc >= cols) continue;
            if (grid[nr][nc] != '1' || visited[nr][nc]) continue;
            visited[nr][nc] = true;
            q.push({nr, nc});
        }
    }
}

// Returns the number of islands (4-directionally connected groups of '1'
// cells) in grid. Runs in O(m*n) time and O(m*n) worst-case space (BFS
// queue) by treating the grid as an implicit graph and counting
// connected components via flood fill from every unvisited land cell.
int numIslands(const vector<string>& grid) {
    if (grid.empty() || grid[0].empty()) return 0;

    int rows = (int)grid.size(), cols = (int)grid[0].size();
    vector<vector<bool>> visited(rows, vector<bool>(cols, false));

    int count = 0;
    for (int r = 0; r < rows; r++) {
        for (int c = 0; c < cols; c++) {
            if (grid[r][c] == '1' && !visited[r][c]) {
                count++;
                floodFill(grid, visited, r, c);
            }
        }
    }
    return count;
}

int main() {
    assert(numIslands({"11000", "11000", "00100", "00011"}) == 3);
    assert(numIslands({"000", "000"}) == 0);
    assert(numIslands({"11", "11"}) == 1);
    assert(numIslands({"10", "01"}) == 2);
    assert(numIslands({}) == 0);
    cout << "All tests passed\n";
    return 0;
}
