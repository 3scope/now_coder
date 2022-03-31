package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// Sort before dedupliates.
	sort.Ints(nums)
	result := make([][]int, 0)
	for index := 0; index < len(nums)-2; index++ {
		number1 := nums[index]
		if number1 > 0 {
			continue
		}
		if index > 0 && nums[index-1] == number1 {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		for left < right {
			if number1+nums[left]+nums[right] > 0 {
				right--
			} else if number1+nums[left]+nums[right] < 0 {
				left++
			} else {
				result = append(result, []int{number1, nums[left], nums[right]})
				left++
				right--
				// Left pointer deduplication.
				for left < right && nums[left-1] == nums[left] {
					left++
				}
				// Right pointer deduplication.
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			}
		}
	}

	return result
}
