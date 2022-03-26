package main

func insertsionSort(nums []int) []int {
	// The first element is sorted at begining.
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		// The index of "i" is used to move the array backwards.
		// The "target" is to store the insert element.
		target := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > target {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func bubbleSort(nums []int) {
	// The bubble sort function only need to run "n" - 1 times.
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				flag = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if !flag {
			break
		}
	}
}

func selectionSort(nums []int) {
	// The last one operation is no need because "n" - 1 times before.
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
	if left >= right {
		return
	}

	// Preorder traversal.
	i, j := left, right
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
	if left >= right {
		return
	}

	// Postorder traversal.
	middle := (left + right) / 2
	mergeSort(nums, left, middle)
	mergeSort(nums, middle+1, right)

	// Merge option.
	temp := make([]int, right-left+1)
	// Reassign index.
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

	// Reassign index, then change the value of the original array.
	leftPointer = left
	i = 0
	for leftPointer <= right {
		nums[leftPointer] = temp[i]
		i++
		leftPointer++
	}
}
