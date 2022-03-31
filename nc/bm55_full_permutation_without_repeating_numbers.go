package main

import "sort"

func permute(num []int) [][]int {
	// write code here
	sort.Ints(num)
	result := make([][]int, 0)
	subResult := make([]int, 0)
	usedIndex := make(map[int]bool)

	permutationBacktracking(num, &usedIndex, &result, &subResult)

	return result
}

func permutationBacktracking(nums []int, usedIndex *map[int]bool, result *[][]int, subResult *[]int) {
	if len(*subResult) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, *subResult)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*usedIndex)[i] {
			continue
		}
		(*usedIndex)[i] = true
		*subResult = append(*subResult, nums[i])
		permutationBacktracking(nums, usedIndex, result, subResult)
		(*usedIndex)[i] = false
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}
