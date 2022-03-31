package main

import "sort"

func permuteUnique(num []int) [][]int {
	// write code here
	sort.Ints(num)
	result := make([][]int, 0, 1)
	subResult := make([]int, 0)
	// The "num" has duplicates in it, so map use index to mark when it was used.
	usedIndex := make(map[int]bool)
	permuteUniqueBacktracking(num, &usedIndex, &result, &subResult)

	return result
}

func permuteUniqueBacktracking(nums []int, usedIndex *map[int]bool, result *[][]int, subResult *[]int) {
	if len(*subResult) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, *subResult)
		*result = append(*result, temp)
		return
	}

	levelUsed := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if levelUsed[nums[i]] || (*usedIndex)[i] {
			continue
		}
		levelUsed[nums[i]] = true
		(*usedIndex)[i] = true
		*subResult = append(*subResult, nums[i])
		permuteUniqueBacktracking(nums, usedIndex, result, subResult)
		(*usedIndex)[i] = false
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}
