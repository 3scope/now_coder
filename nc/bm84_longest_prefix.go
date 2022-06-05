package main

import "sort"

// 如果不排序的话，需要比较所有的字符串，ru g
func longestCommonPrefix(strs []string) string {
	// 特殊情况，没有字符串。
	if len(strs) == 0 {
		return ""
	}

	count := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// 如果有字符串长度不够，或者字符不匹配，那么最长前缀长度就为“count”。
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:count]
			}
		}
		count++
	}

	return strs[0][:count]
}

func longestCommonPrefixSort(strs []string) string {
	// 特殊情况，没有任何字符串。
	if len(strs) == 0 {
		return ""
	}

	// 按照字典序排序。
	sort.Strings(strs)
	// 排完序后，只需要比较字典序最小的字符串和字典序最大的字符串。
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
