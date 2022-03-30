package main

import "sort"

func longestCommonPrefix(strs []string) string {
	// write code here
	if len(strs) == 0 {
		return ""
	}

	count := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:count]
			}
		}
		count++
	}

	return strs[0][:count]
}

func longestCommonPrefixSort(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	// When sorted lexicographically, just need to compare the first and last.
	first := strs[0]
	last := strs[len(strs)-1]
	count := 0
	for i := 0; i < len(first) && i < len(last); i++ {
		if first[i] == last[i] {
			count++
		} else {
			break
		}
	}

	return first[:count]
}
