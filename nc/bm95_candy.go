package main

func candy(arr []int) int {
	// write code here
	result := make([]int, len(arr))
	result[0] = 1

	for i := 1; i < len(result); i++ {
		result[i] = 1
		// If right greater than left, than plus one candy.
		if arr[i] > arr[i-1] {
			result[i] = result[i-1] + 1
		}
	}

	for i := len(result) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			// Need to update.
			if result[i+1]+1 > result[i] {
				result[i] = result[i+1] + 1
			}
		}
	}

	sum := 0
	for i := 0; i < len(result); i++ {
		sum += result[i]
	}

	return sum
}
