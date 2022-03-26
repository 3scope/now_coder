package main

import (
	"strconv"
	"strings"
)

func compare(version1 string, version2 string) int {
	// write code here
	versionOneSlice := strings.Split(version1, ".")
	versionTwoSlice := strings.Split(version2, ".")
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	length := max(len(versionOneSlice), len(versionTwoSlice))

	var i, one, two int
	for ; i < length; i++ {
		if i < len(versionOneSlice) {
			one, _ = strconv.Atoi(versionOneSlice[i])
		} else {
			one = 0
		}
		if i < len(versionTwoSlice) {
			two, _ = strconv.Atoi(versionTwoSlice[i])
		} else {
			two = 0
		}
		if one > two {
			return 1
		} else if one < two {
			return -1
		}
	}
	return 0
}
