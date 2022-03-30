package main

func rotateMatrix(mat [][]int, n int) [][]int {
	// write code here
	if mat == nil {
		return nil
	}
	for i := 0; i < len(mat); i++ {
		// If "j" equals to "i", the element is on diagonal.
		for j := 0; j < i; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	// Reverse each row of matrix.
	for i := 0; i < len(mat); i++ {
		for j, k := 0, len(mat[i])-1; j < k; j, k = j+1, k-1 {
			mat[i][j], mat[i][k] = mat[i][k], mat[i][j]
		}
	}

	return mat
}
