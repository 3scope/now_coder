package main

import "strings"

func trans(s string, n int) string {
	// write code here
	// Get all worlds.
	worldSlice := strings.Split(s, " ")
	// Reverse the index of worlds.
	i, j := 0, len(worldSlice)-1
	for ; i < j; i, j = i+1, j-1 {
		worldSlice[i] = changeString(worldSlice[i])
		worldSlice[j] = changeString(worldSlice[j])
		worldSlice[i], worldSlice[j] = worldSlice[j], worldSlice[i]
	}

	// Prevent the middle world do not convert.
	if i == j {
		worldSlice[i] = changeString(worldSlice[i])
	}
	s = strings.Join(worldSlice, " ")

	return s
}

// func reverseString(s string) string {
// 	sByte := []byte(s)
// 	for i, j := 0, len(sByte)-1; i < j; i, j = i+1, j-1 {
// 		sByte[i], sByte[j] = sByte[j], sByte[i]
// 	}
// 	return string(sByte)
// }

func changeString(s string) string {
	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		sByte[i] = changeCase(sByte[i])
	}

	return string(sByte)
}

func changeCase(c byte) byte {
	if c >= 'a' && c <= 'z' {
		c = c - 'a' + 'A'
	} else if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}

	return c
}
