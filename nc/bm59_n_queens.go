package main

func Nqueen(n int) int {
	// write code here
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	result := 0
	NqueenBacktracking(&board, 0, &result)

	return result
}

func NqueenBacktracking(board *[][]bool, row int, result *int) {
	// When the "0 ~ len(board) - 1" row are all set, it finds a result.
	if row == len(*board) {
		(*result)++
		return
	}

	if row > len(*board) {
		return
	}

	for i := 0; i < len((*board)[row]); i++ {
		if IsValidQueen(*board, row, i) {
			(*board)[row][i] = true
			NqueenBacktracking(board, row+1, result)
			(*board)[row][i] = false
		}
	}
}

func IsValidQueen(board [][]bool, row, column int) bool {
	for i := 0; i < row; i++ {
		if board[i][column] {
			return false
		}
	}
	// Cannot be diagonal.
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	for i, j := row-1, column+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}
