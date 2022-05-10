package main

// Backtracking + dfs + memoization.
func solveGrid(grid [][]byte) int {
	// write code here
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfsSolve(grid, i, j)
			}
		}
	}

	return count
}

// Each time use bfs function, it turn one island gone.
func dfsSolve(grid [][]byte, row, column int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[row]) || grid[row][column] == '0' {
		return
	}

	grid[row][column] = '0'
	// Recursion.
	dfsSolve(grid, row-1, column)
	dfsSolve(grid, row+1, column)
	dfsSolve(grid, row, column+1)
	dfsSolve(grid, row, column-1)
}
