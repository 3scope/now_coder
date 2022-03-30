package main

func maxLength(arr []int) int {
	// write code here
	if len(arr) == 0 {
		return 0
	}

	dedup := make(map[int]bool)
	left, right := 0, 0
	result := 0
	for right < len(arr) {
		if dedup[arr[right]] {
			dedup[arr[left]] = false
			left++
		} else {
			dedup[arr[right]] = true
			right++
		}
		// Store the max size.
		if result < right-left {
			result = right - left
		}
	}

	return result
}
