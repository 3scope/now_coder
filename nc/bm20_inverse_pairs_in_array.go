package main

func InversePairs(data []int) int {
	// 使用归并排序，在归并的中途可以判断逆序对的数量。
	return mergeSortArray(data, 0, len(data)-1)
}

// 函数返回值是逆序对的数量。
func mergeSortArray(nums []int, left, right int) int {
	// 只有一个元素，没有任何逆序对。
	if left >= right {
		return 0
	}

	// 归并排序属于后序遍历，首先得到排好序的左区间和右区间，之后合并成一个有序区间。
	// 所有区间是左闭右闭的（“[ ]”）。
	middle := (left + right) / 2
	leftCount := mergeSortArray(nums, left, middle)
	rightCount := mergeSortArray(nums, middle+1, right)

	// 进行合并，首先申请一片新的空间，用于存放中间值。
	temp := make([]int, right-left+1)
	// 收集左右区间的逆序对，之后加上自身的逆序对。
	count := leftCount + rightCount
	// 定义两个归并指针，对左右区间进行归并
	leftPointer := left
	rightPointer := middle + 1
	// 变量“i”存放“temp”中下一个排序元素的位置。
	i := 0
	for ; i < len(temp); i++ {
		// 如果其中一个数组首先归并完成，则直接跳出循环。
		if leftPointer > middle || rightPointer > right {
			break
		}
		// 如果左边区间的元素小于或者等于右边区间，则没有逆序对。
		if nums[leftPointer] <= nums[rightPointer] {
			// 归并操作。
			temp[i] = nums[leftPointer]
			leftPointer++
		} else {
			// 如果左边区间元素大于右边区间，则逆序对数量由左边区间剩余长度决定。
			count += middle - leftPointer + 1
			// 归并操作。
			temp[i] = nums[rightPointer]
			rightPointer++
		}
	}
	// 收集剩余的区间元素。
	for leftPointer <= middle {
		temp[i] = nums[leftPointer]
		i++
		leftPointer++
	}
	for rightPointer <= right {
		temp[i] = nums[rightPointer]
		i++
		rightPointer++
	}

	// 将“temp”中的值写回原数组。
	leftPointer = left
	i = 0
	for leftPointer <= right {
		nums[leftPointer] = temp[i]
		i++
		leftPointer++
	}
	return count % 1000000007
}
