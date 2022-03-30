package main

import "sort"

func minmumNumberOfHost(n int, startEnd [][]int) int {
	// write code here
	count := make(map[int]int)
	// The count of host only increases at the start time, only decreases at the end.
	for i := 0; i < len(startEnd); i++ {
		count[startEnd[i][0]]++
		count[startEnd[i][1]]--
	}

	// Sort the count.
	keys := make([]int, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	sum := 0
	result := 0
	for i := 0; i < len(keys); i++ {
		sum += count[keys[i]]
		if sum > result {
			result = sum
		}
	}

	return result
}
