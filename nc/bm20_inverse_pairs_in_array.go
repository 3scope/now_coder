package main

func InversePairs(data []int) int {
	// write code here
	return mergeSortArray(data, 0, len(data)-1)
}

func mergeSortArray(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	// Postorder traversal.
	middle := (left + right) / 2
	leftCount := mergeSortArray(nums, left, middle)
	rightCount := mergeSortArray(nums, middle+1, right)

	// Merge option.
	temp := make([]int, right-left+1)
	// Inverse pairs.
	count := leftCount + rightCount
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
			// The left and the right array are all sorted.
			count += middle - leftPointer + 1

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
	return count % 1000000007
}
