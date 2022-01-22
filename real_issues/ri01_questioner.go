package main

import (
	"fmt"
	"sort"
)

func Solution(total int, dif []int) int {
	result := 0
	sort.Ints(dif)
	combo := 1

	for i := 0; i < total-1; {
		if dif[i+1]-dif[i] <= 10 {
			combo++
			i++
		} else if dif[i+1]-dif[i] <= 20 {
			if combo == 1 {
				result++
				i += 2
			} else if combo == 2 {
				result++
				i++
				combo = 1
			}
		} else {
			if combo == 1 {
				result += 2
				i++
			} else if combo == 2 {
				result++
				i++
				combo = 1
			}
		}
	}

	result += 3 - (total+result)%3
	return result
}

func main() {
	total := 0
	fmt.Scan(&total)

	dif := make([]int, total)
	for i := 0; i < total; i++ {
		fmt.Scan(&dif[i])
	}

	fmt.Println(Solution(total, dif))
}
