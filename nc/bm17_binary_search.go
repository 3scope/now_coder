package main

func search(nums []int, target int) int {
	// 区间是左闭右闭的，因此左指针会一直指向可能区间的第一个值，右指针会一直指向可能区间的最后一个值。
	left := 0
	right := len(nums) - 1
	// 当“left”和“right”指针相遇时，可能区间的大小为1。
	for left <= right {
		middle := (left + right) / 2
		// 中间值比目标值大， 证明目标值在左半区间。
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			// 匹配成功。
			return middle
		}
	}
	return -1
}

func searchBinary(nums []int, target int) int {
	// 区间是左闭右开的，因此左指针会一直指向可能区间的第一个值，右指针会一直指向可能区间最后一个值的下一个值。
	left := 0
	right := len(nums)
	// 当“left”和“right”指针相遇时，证明可能区间的长度为零，返回-1。
	for left < right {
		middle := (left + right) / 2
		if nums[middle] > target {
			// “middle”处的值以及不可能是匹配的了，因此右指针指向“middle”。
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
