package main

// 中位数的下标是K，也就是求第K小的数。
func findMedianinTwoSortedAray(arr1 []int, arr2 []int) int {
	// write code here
	if len(arr1) == 0 {
		return 0
	}
	k := len(arr1)
	// 每次排除 k / 2 个，k / 2 向下取整。
	// k的值不断变化，对算法没有影响。
	// 如果是求中位数的话，求第 k 小和第 k+1 小的数，之后平均一下。
	arr1Start, arr1End := 0, 0
	arr2Start, arr2End := 0, 0
	// Start 指向没被删除的第一个，End 指向没被删除的最后一个的下一个。
	// 左闭右开。
	del := k / 2
	arr1End += del
	arr2End += del
	min := func(i, j int) int {
		if i > j {
			return j
		} else {
			return i
		}
	}
	for k > 1 && arr1Start < len(arr1) && arr2Start < len(arr2) {
		if arr1End > len(arr1) {
			arr1End = len(arr1)
		}
		if arr2End > len(arr2) {
			arr2End = len(arr2)
		}

		if arr1[arr1End-1] >= arr2[arr2End-1] {
			k = k - (arr2End - arr2Start)
			// 删除操作。
			arr2Start = arr2End
			// 更新下次需要比较的区间。
			arr2End += k / 2
			arr1End = arr1Start + k/2
		} else {
			k = k - (arr1End - arr1Start)
			arr1Start = arr1End
			arr1End += k / 2
			arr2End = arr2Start + k/2
		}
	}

	return min(arr1[arr1Start], arr2[arr2Start])
}
