package main

// When use greedy algorithm, it will not treat.

// func findPeakElement(nums []int) int {
// 	// write code here
// 	if len(nums) == 1 {
// 		return 0
// 	}

// 	preDifference := nums[1] - nums[0]
// 	curDifference := 0
// 	result := 0
// 	if nums[1] > nums[0] {
// 		result = 1
// 	}
// 	for i := 2; i < len(nums); i++ {
// 		curDifference = nums[i] - nums[i-1]
// 		if curDifference > 0 && preDifference <= 0 || curDifference < 0 && preDifference >= 0 {
// 			return i - 1
// 		}
// 	}

// 	return result
// }

func findPeakElementBinarySearch(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		middle := (left + right) / 2
		// Middle is not the peak.
		if nums[middle] < nums[middle+1] {
			left = middle + 1
		} else {
			// If middle equals to middle + 1 or more than middle + 1, than it may be the peak.
			right = middle
		}
	}

	return right
}
