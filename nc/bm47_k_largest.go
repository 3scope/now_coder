package main

func findKth(a []int, n int, K int) int {
	// write code here
	quickSortFindKth(a, 0, len(a)-1, K)

	return a[K-1]
}

func quickSortFindKth(nums []int, left, right int, K int) {
	// Termination condition.
	if left >= right {
		return
	}

	i, j := left, right
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] <= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot

	// Choose one.
	if i == K-1 {
		return
	} else if i < K-1 {
		quickSortFindKth(nums, i+1, right, K)
	} else {
		quickSortFindKth(nums, left, i-1, K)
	}
}
