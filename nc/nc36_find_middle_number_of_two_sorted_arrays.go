package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(findKthInTwoArray(nums1, nums2, length/2+1))
	} else {
		return (float64(findKthInTwoArray(nums1, nums2, length/2)) + float64(findKthInTwoArray(nums1, nums2, length/2+1))) / 2
	}
}

// 中位数的下标是“k”，也就是求第“k”小的数（“k”的值最小为1）。
func findKthInTwoArray(arr1 []int, arr2 []int, k int) int {
	// 每次排除“k/2”个，结果向下取整。
	// 如果连个数组的长度是偶数的话，需要求第“k”小和第“k+1”小的数，之后平均一下。
	// “Start”指向没被删除的第一个，“End”指向没被删除的最后一个的下一个。
	// 即待删除区间是左闭右开“[ )”
	arr1Start, arr1End := 0, 0
	arr2Start, arr2End := 0, 0
	// 初始化第一次的待删除区间。
	arr1End = k / 2
	arr2End = k / 2
	min := func(i, j int) int {
		if i > j {
			return j
		} else {
			return i
		}
	}
	// 当“k”的值等于1时，跳出循环，返回两个待删除区间中的最小值即可。
	// 当某个数组所有值都被删除后，跳出循环，返回剩下的一个数组的第“k-1”个值即可。
	for k > 1 && arr1Start < len(arr1) && arr2Start < len(arr2) {
		// “arrEnd”的最大值是对应数组的长度。
		if arr1End > len(arr1) {
			arr1End = len(arr1)
		}
		if arr2End > len(arr2) {
			arr2End = len(arr2)
		}

		// 比较待删除区间的最大值，小的一方中对应的数值不可能是第“k”小的值。
		if arr1[arr1End-1] >= arr2[arr2End-1] {
			// 删除的长度是“arr2End - arr2Start”，更新“k”的值。
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

	// 当某个数组所有值都被删除后，跳出循环，返回剩下的一个数组的第“k-1”个值即可。
	if arr1Start == len(arr1) {
		return arr2[arr2Start+k-1]
	} else if arr2Start == len(arr2) {
		return arr1[arr1Start+k-1]
	}
	// 否则返回两个头指针中的最小值即可。
	return min(arr1[arr1Start], arr2[arr2Start])
}
