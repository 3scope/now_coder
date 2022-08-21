package main

func maxLength(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// 用于判断当前中是否已有相应的数字。
	dedup := make(map[int]bool)
	left, right := 0, 0
	result := 0
	for right < len(arr) {
		// 如果出现了重复，那么就需要缩短，此时窗口变小。
		if dedup[arr[right]] {
			dedup[arr[left]] = false
			left++
		} else {
			// 如果没有，也就是满足条件，那么就继续变长。
			dedup[arr[right]] = true
			right++
		}
		// 记录窗口的最大值。
		if result < right-left {
			result = right - left
		}
	}

	return result
}
