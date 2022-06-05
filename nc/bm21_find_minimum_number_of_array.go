package main

func minNumberInRotateArray(rotateArray []int) int {
	// 将一个递增数字循环左移“<-”，得到旋转数组，因此比较最右边的值和“middle”的大小，就可以知道最小值所在位置（比较最左边的值和“middle”的大小，就可以知道最大值的位置）。
	left := 0
	right := len(rotateArray) - 1
	middle := (left + right) / 2
	// 当区间只有一个元素的时候，直接返回结果，不需要继续计算。
	for left < right {
		if rotateArray[middle] > rotateArray[right] {
			// 最小值在右边区间，并且自身不会是最小值。
			left = middle + 1
		} else if rotateArray[middle] < rotateArray[right] {
			// 最小值在左边区间，并且“middle”有可能是最小值。
			right = middle
		} else {
			// 如果相等，不断右移“right”指针，那么就会找到真正的最小值。
			right--
		}
		middle = (left + right) / 2
	}

	return rotateArray[middle]
}
