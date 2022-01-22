package main

import (
	"sort"
	"strconv"
)

type StringList []string

func Soluntion(n, m int) int {
	resultString := make([]string, n)
	for i := 1; i <= n; i++ {
		resultString[i-1] = strconv.Itoa(i)
	}
	sort.Sort(StringList(resultString))

	result, _ := strconv.Atoi(resultString[m-1])
	return result
}

func (s StringList) Len() int {
	return len(s)
}

func (s StringList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// func main() {
// 	n, m := 0, 0
// 	fmt.Scan(&n, &m)

// 	fmt.Println(Soluntion(n, m))
// }
