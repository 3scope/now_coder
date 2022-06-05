package main

import (
	"strconv"
	"strings"
)

func compare(version1 string, version2 string) int {
	// 版本号用“.”隔开，所以先获得每个版本具体的值。
	versionOneSlice := strings.Split(version1, ".")
	versionTwoSlice := strings.Split(version2, ".")
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	// 获得其中最长版本号的长度，因为“1.1”和“1.1.0”是相同的，所以可以为短的一方补零。
	length := max(len(versionOneSlice), len(versionTwoSlice))

	var one, two int
	for i := 0; i < length; i++ {
		if i < len(versionOneSlice) {
			// 转换为整型的时候会去掉前导零。
			one, _ = strconv.Atoi(versionOneSlice[i])
		} else {
			// 短的一方直接补零。
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
