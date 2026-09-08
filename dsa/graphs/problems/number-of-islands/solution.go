// Package numberofislands solves: count the connected land regions in a
// grid of '1' (land) and '0' (water) cells.
package numberofislands

// NumIslands returns the number of islands (4-directionally connected
// groups of '1' cells) in grid. Runs in O(m*n) time and O(m*n) worst-case
// space (BFS queue) by treating the grid as an implicit graph and counting
// connected components via flood fill from every unvisited land cell.
// grid is read but not mutated; a separate visited matrix tracks state.
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' && !visited[r][c] {
				count++
				floodFill(grid, visited, r, c)
			}
		}
	}
	return count
}

func floodFill(grid [][]byte, visited [][]bool, startR, startC int) {
	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{{startR, startC}}
	visited[startR][startC] = true

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, d := range dirs {
			nr, nc := cell[0]+d[0], cell[1]+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if grid[nr][nc] != '1' || visited[nr][nc] {
				continue
			}
			visited[nr][nc] = true
			queue = append(queue, [2]int{nr, nc})
		}
	}
}
