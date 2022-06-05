package main

// 二分查找的本质是找一个中间值，保证每次可以去除多个错误答案。
func Find(target int, array [][]int) bool {
	// 特殊情况。
	if len(array) == 0 {
		return false
	}
	rows := len(array)
	columns := len(array[0])
	// 从数组的右上角出发，左边的数比它小，右边的数比它大。
	middle := [2]int{0, columns - 1}
	for middle[0] < rows && middle[1] >= 0 {
		element := array[middle[0]][middle[1]]
		if element == target {
			return true
		} else if element > target {
			middle[1]--
		} else if element < target {
		}
		middle[0]++
	}
	// 没有找到则表明不存在。
	return false
}
