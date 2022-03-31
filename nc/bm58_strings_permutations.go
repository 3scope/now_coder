package main

import "sort"

func Permutation(str string) []string {
	// write code here
	byteSlice := []byte(str)
	sort.Slice(byteSlice, func(i, j int) bool {
		return byteSlice[i] < byteSlice[j]
	})
	result := make([]string, 0, 1)
	subResult := make([]byte, 0)
	usedIndex := make(map[int]bool)
	PermutationBacktracking(byteSlice, &usedIndex, &result, &subResult)

	return result
}

func PermutationBacktracking(str []byte, usedIndex *map[int]bool, result *[]string, subResult *[]byte) {
	if len(*subResult) == len(str) {
		temp := string(*subResult)
		*result = append(*result, temp)
		return
	}

	levelUsed := make(map[byte]bool)
	for i := 0; i < len(str); i++ {
		if levelUsed[str[i]] || (*usedIndex)[i] {
			continue
		}
		levelUsed[str[i]] = true
		(*usedIndex)[i] = true
		*subResult = append(*subResult, str[i])
		PermutationBacktracking(str, usedIndex, result, subResult)
		(*usedIndex)[i] = false
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}
