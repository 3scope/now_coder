package main

import "sort"

func FindNumsAppearOnce(array []int) []int {
	// write code here
	result := make([]int, 0, 2)
	count := make(map[int]bool)
	for i := 0; i < len(array); i++ {
		if count[array[i]] {
			delete(count, array[i])
		} else {
			count[array[i]] = true
		}
	}

	for key := range count {
		result = append(result, key)
	}
	sort.Ints(result)

	return result
}
