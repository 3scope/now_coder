package main

import "sort"

func minmumNumberOfHost(n int, startEnd [][]int) int {
	count := make(map[int]int)
	// 首先求的差分数组。
	for i := 0; i < len(startEnd); i++ {
		// 活动举办时主持人数应该加一。
		count[startEnd[i][0]]++
		// 结束时主持人数减一。
		count[startEnd[i][1]]--
	}

	// 通过时间排序，之后求前缀和，那么前缀和里面最大的值就是最后的结果。
	keys := make([]int, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// “sum”代表每个时刻需要主持人的数量。
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
