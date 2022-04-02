package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	count := strings.Split(input.Text(), " ")
	// Get the money number and things number.
	money, _ := strconv.Atoi(count[0])
	itemCount, _ := strconv.Atoi(count[1])
	// The first row is nil.
	items := make([][3]int, 1)
	for i := 0; i < itemCount; i++ {
		input.Scan()
		temp := strings.Split(input.Text(), " ")
		item := [3]int{}
		item[0], _ = strconv.Atoi(temp[0])
		item[1], _ = strconv.Atoi(temp[1])
		item[2], _ = strconv.Atoi(temp[2])
		items = append(items, item)
	}

	// Set the main items.
	prices := make(map[int][]int)
	satisfaction := make(map[int][]int)
	for i := 1; i < len(items); i++ {
		if items[i][2] == 0 {
			prices[i] = []int{items[i][0]}
			satisfaction[i] = []int{items[i][0] * items[i][1]}
		}
	}
	// Set attachments.
	for i := 1; i < len(items); i++ {
		if items[i][2] > 0 {
			index := items[i][2]
			prices[index] = append(prices[index], prices[index][0]+items[i][0])
			satisfaction[index] = append(satisfaction[index], satisfaction[index][0]+items[i][0]*items[i][1])
			// When a main item has two attachments.
			if len(prices[index]) == 3 {
				prices[index] = append(prices[index], prices[index][1]+items[i][0])
				satisfaction[index] = append(satisfaction[index], satisfaction[index][1]+items[i][0]*items[i][1])
			}
		}
	}

	result := getMaxSatisfaction(prices, satisfaction, money)
	fmt.Println(result)
}

func getMaxSatisfaction(prices, satisfaction map[int][]int, money int) int {
	dp := make([]int, money+1)
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	// Use "j" money, the maximum satisfaction is "dp[j]".
	dp[0] = 0
	for key, value := range prices {
		for j := money; j >= value[0]; j-- {
			for k := 0; k < len(value); k++ {
				if j-value[k] >= 0 {
					dp[j] = max(dp[j], dp[j-value[k]]+satisfaction[key][k])
				}
			}
		}
	}

	return dp[money]
}
