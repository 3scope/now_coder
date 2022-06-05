package main

func insertsionSort(nums []int) {
	// 第一个数默认有序。
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		// 存储需要插入的数值。
		target := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > target {
				// 当前数值比待排序数值大，将它往后移。
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				// 找到了待排序数值的位置。
				break
			}
		}
	}
}

func bubbleSort(nums []int) {
	// 冒泡排序最多执行“n-1”轮。
	// “flag”代表本轮是否出现交换，“false”证明没有交换，那么整个数组就是有序的。
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				flag = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		// 如果没有产生交换，直接跳出循环。
		if !flag {
			break
		}
	}
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

func quickSort(nums []int, left, right int) {
	// 当左指针越过右指针时，证明这个区间没有待排序列，直接返回。
	if left >= right {
		return
	}

	// 快速排序类似于先序遍历，先处理当前节点的值，之后递归。
	i, j := left, right
	// 选取一个枢纽，默认是“left”指针所在的位置。
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func mergeSort(nums []int, left, right int) {
	// 递归终止条件。
	if left >= right {
		return
	}

	// 类似于二叉树后序遍历，先处理左右孩子，之后处理根节点。
	middle := left + (right-left)/2
	mergeSort(nums, left, middle)
	mergeSort(nums, middle+1, right)

	// 进行归并。
	temp := make([]int, right-left+1)
	// “temp”数组用于临时存放归并结果，之后写到原数组中去。
	leftPointer := left
	rightPointer := middle + 1
	i := 0
	for ; i < len(temp); i++ {
		if leftPointer > middle || rightPointer > right {
			break
		}
		if nums[leftPointer] <= nums[rightPointer] {
			temp[i] = nums[leftPointer]
			leftPointer++
		} else {
			temp[i] = nums[rightPointer]
			rightPointer++
		}
	}
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

	// 写回到原数组。
	leftPointer = left
	i = 0
	for leftPointer <= right {
		nums[leftPointer] = temp[i]
		i++
		leftPointer++
	}
}

func heapSort(nums []int) {
	// 定义堆的上浮过程和下沉过程，默认使用的是小根堆。
	var shiftUp func(nums []int, index int)
	var shiftDown func(nums []int, length int) int
	// 堆元素上浮，每次插入时，都会将节点插入到数组的最后一个位置。
	shiftUp = func(nums []int, index int) {
		// 只要没到堆顶，就尝试上浮。
		for index > 0 {
			// 当前元素小于父亲节点，那么就交换位置，没有就跳出循环。
			father := (index - 1) / 2
			if nums[index] < nums[father] {
				nums[index], nums[father] = nums[father], nums[index]
				// 更新“index”的值。
				index = father
			} else {
				break
			}
		}
	}

	// 堆元素下沉，每次删除时，都会交换最后一个元素的堆顶元素的值，堆的长度减一。
	shiftDown = func(nums []int, length int) int {
		nums[0], nums[length-1] = nums[length-1], nums[0]
		length--
		// 判断当前节点有哪些孩子。
		index := 0
		for 2*index+1 < length {
			// 存放左右节点哪个最小。
			minIndex := index
			if nums[minIndex] > nums[2*index+1] {
				minIndex = 2*index + 1
			}
			if 2*index+2 < length {
				if nums[minIndex] > nums[2*index+2] {
					minIndex = 2*index + 2
				}
			}
			// 如果没有一个节点比它小，那么退出循环。
			if index == minIndex {
				break
			} else {
				nums[index], nums[minIndex] = nums[minIndex], nums[index]
				// 更新“index”的值。
				index = minIndex
			}
		}
		return length
	}

	// 构建堆。
	for i := 1; i < len(nums); i++ {
		// 从前往后，向上堆化。
		shiftUp(nums, i)
	}

	// for i := (len(nums) - 1/2); i >= 0; i-- {
	// 	shiftDown(nums, len(nums))
	// }

	// 排序。
	length := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		length = shiftDown(nums, length)
	}
}
