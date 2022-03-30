package main

func spiralOrder(matrix [][]int) []int {
	// write code here
	if matrix == nil {
		return nil
	}
	// The four corners.
	row := len(matrix)
	column := len(matrix[0])
	leftTop := []int{0, 0}
	rightBottom := []int{row - 1, column - 1}

	result := make([]int, 0, row*column)
	for leftTop[0] < rightBottom[0] && leftTop[1] < rightBottom[1] {
		// Top edge.
		// Begin with left and end with right.
		for j := leftTop[1]; j < rightBottom[1]; j++ {
			result = append(result, matrix[leftTop[0]][j])
		}
		// Right edge.
		for j := leftTop[0]; j < rightBottom[0]; j++ {
			result = append(result, matrix[j][rightBottom[1]])
		}
		// Bottom edge.
		for j := rightBottom[1]; j > leftTop[1]; j-- {
			result = append(result, matrix[rightBottom[0]][j])
		}
		// Left edge.
		for j := rightBottom[0]; j > leftTop[0]; j-- {
			result = append(result, matrix[j][leftTop[1]])
		}
		leftTop[0]++
		leftTop[1]++
		rightBottom[0]--
		rightBottom[1]--
	}

	// When the one of row and column is odd, special handling required.
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	if min(row, column)%2 == 1 {
		// If row number is equal.
		if leftTop[0] == rightBottom[0] {
			for j := leftTop[1]; j <= rightBottom[1]; j++ {
				result = append(result, matrix[leftTop[0]][j])
			}
		} else if leftTop[1] == rightBottom[1] {
			for j := leftTop[0]; j <= rightBottom[0]; j++ {
				result = append(result, matrix[j][leftTop[1]])
			}
		}
	}

	return result
}
