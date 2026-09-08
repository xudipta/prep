// Package nqueens solves: place n queens on an n x n board so that no two
// attack each other, returning every distinct solution.
package nqueens

import "strings"

// SolveNQueens returns every solution to the n-queens problem, each as a
// slice of n strings representing the board ('Q' for a queen, '.'
// otherwise). Runs in pruned-exponential time (no polynomial solution is
// known) and O(n) space for bookkeeping by placing one queen per row and
// tracking occupied columns and diagonals in sets, so each placement's
// safety is checked in O(1) instead of rescanning the board.
func SolveNQueens(n int) [][]string {
	cols := make(map[int]bool)
	diag1 := make(map[int]bool) // row - col
	diag2 := make(map[int]bool) // row + col
	queenCol := make([]int, n)  // queenCol[row] = column of the queen in that row

	var result [][]string

	var place func(row int)
	place = func(row int) {
		if row == n {
			result = append(result, buildBoard(queenCol, n))
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue // conflict; prune
			}

			cols[col], diag1[row-col], diag2[row+col] = true, true, true
			queenCol[row] = col

			place(row + 1)

			cols[col], diag1[row-col], diag2[row+col] = false, false, false
		}
	}

	place(0)
	return result
}

func buildBoard(queenCol []int, n int) []string {
	board := make([]string, n)
	for row, col := range queenCol {
		var sb strings.Builder
		for c := 0; c < n; c++ {
			if c == col {
				sb.WriteByte('Q')
			} else {
				sb.WriteByte('.')
			}
		}
		board[row] = sb.String()
	}
	return board
}
