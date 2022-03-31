package main

func minNumberDisappeared(nums []int) int {
	// write code here
	// Use extra space.
	isExisted := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			isExisted[nums[i]] = true
		}
	}

	for i := 1; i <= len(nums); i++ {
		if !isExisted[i] {
			return i
		}
	}
	return len(nums) + 1
}

func minNumberNotExisting(nums []int) int {
	// Set the number to its index.
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] <= len(nums) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
