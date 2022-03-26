package main

func Find(target int, array [][]int) bool {
	// write code here
	row := len(array)
	if row == 0 {
		return false
	}
	column := len(array[0])
	// The element of the right top is the middle, because the left of it smaller than it, and buttom of it bigger than it.
	// Middle.
	pos := []int{0, column - 1}
	for pos[0] < row && pos[1] >= 0 {
		if array[pos[0]][pos[1]] == target {
			return true
		} else if array[pos[0]][pos[1]] > target {
			pos[1]--
		} else {
			pos[0]++
		}
	}

	return false
}
