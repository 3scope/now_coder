package main

func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 1 {
		return rotateArray[0]
	}

	left := 0
	right := len(rotateArray) - 1
	for left <= right {
		middle := (left + right) / 2
		if rotateArray[middle] > rotateArray[right] {
			left = middle + 1
		} else if rotateArray[middle] < rotateArray[right] {
			right = middle
		} else {
			right--
		}
	}
	return rotateArray[right]
}
