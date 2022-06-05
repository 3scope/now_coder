package main

import "strings"

func trans(s string, n int) string {
	// 需要反序单词，那么先获得所有单词，之后交换顺序。
	worldSlice := strings.Split(s, " ")
	// 先处理单词的大小写，之后交换顺序。
	i, j := 0, len(worldSlice)-1
	for ; i < j; i, j = i+1, j-1 {
		worldSlice[i] = changeString(worldSlice[i])
		worldSlice[j] = changeString(worldSlice[j])
		worldSlice[i], worldSlice[j] = worldSlice[j], worldSlice[i]
	}

	// 当“i”和“j”相等时，需要处理最中间的单词。
	if i == j {
		worldSlice[i] = changeString(worldSlice[i])
	}
	s = strings.Join(worldSlice, " ")

	return s
}

// 改变字符串所有字符的大小写。
func changeString(s string) string {
	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		sByte[i] = changeCase(sByte[i])
	}

	return string(sByte)
}

// 改变字符的大小写。
func changeCase(c byte) byte {
	if c >= 'a' && c <= 'z' {
		c = c - 'a' + 'A'
	} else if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}
	return c
}
