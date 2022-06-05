package main

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	// 将左上角和右下角的点不断向中间逼近，之后按照顺时针的方向获得结果。
	row := len(matrix)
	column := len(matrix[0])
	leftTop := []int{0, 0}
	rightBottom := []int{row - 1, column - 1}

	result := make([]int, 0, row*column)
	for leftTop[0] < rightBottom[0] && leftTop[1] < rightBottom[1] {
		// 上面的边，保证遍历区间的左闭右开。
		for j := leftTop[1]; j < rightBottom[1]; j++ {
			result = append(result, matrix[leftTop[0]][j])
		}
		// 右边的边。
		for j := leftTop[0]; j < rightBottom[0]; j++ {
			result = append(result, matrix[j][rightBottom[1]])
		}
		// 下面的边。
		for j := rightBottom[1]; j > leftTop[1]; j-- {
			result = append(result, matrix[rightBottom[0]][j])
		}
		// 左边的边。
		for j := rightBottom[0]; j > leftTop[0]; j-- {
			result = append(result, matrix[j][leftTop[1]])
		}
		// 将左上角和右下角的点不断向中间逼近。
		leftTop[0]++
		leftTop[1]++
		rightBottom[0]--
		rightBottom[1]--
	}

	// 当矩阵短的那条边是奇数的时候，需要特殊处理。
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	if min(row, column)%2 == 1 {
		// 如果行数相等，证明矩阵的行数“rows”比“columns”要短，因此最后按照从左到右的顺序处理中间一行。
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
