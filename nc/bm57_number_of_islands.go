package main

func solveGrid(grid [][]byte) int {
	count := 0

	// 方向数组。
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 每次碰到一个新的岛屿，就将它翻转为0。
			if grid[i][j] == '1' {
				count++
				dfsIslands(grid, i, j, directions)
			}
		}
	}

	return count
}

// “DFS”图的深度优先遍历和树的深度优先遍历很类似，都是先向着某个方向遍历，直到无法到下一步为止。
func dfsIslands(grid [][]byte, row, column int, directions [][2]int) {
	// 终止条件。
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[row]) {
		return
	}
	// 只有当前节点没有访问过，或者能够访问的时候，才可以继续递归，类似于记忆化递归中的“memo”矩阵。
	// “grid[row][column]”的值为“0”，代表之前访问过或者无法访问。
	if grid[row][column] == '0' {
		return
	}

	// 本层逻辑处理。
	grid[row][column] = '0'
	// 深度优先搜索。
	for _, direction := range directions {
		newRow := row + direction[0]
		newColumn := column + direction[1]
		dfsIslands(grid, newRow, newColumn, directions)
	}
}
