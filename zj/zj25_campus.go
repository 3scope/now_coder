package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	input.Scan()
	numsStr := strings.Split(input.Text(), " ")
	nums := make([]int, len(numsStr))
	for i := 0; i < len(numsStr); i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}
	sort.Ints(nums)
	fmt.Println(getResult(nums))
}

func getResult(nums []int) int {
	result := 0
	index := 0
	for index+2 < len(nums) {
		if nums[index+1]-nums[index] <= 10 && nums[index+2]-nums[index+1] <= 10 {
			index += 3
		} else if nums[index+1]-nums[index] <= 20 {
			result += 1
			index += 2
		} else if nums[index+1]-nums[index] > 20 {
			result += 2
			index += 1
		} else {
			result += 1
			index += 2
		}
	}

	if index == len(nums)-1 {
		result += 2
	} else if index == len(nums)-2 {
		if nums[index+1]-nums[index] <= 20 {
			result += 1
		} else {
			result += 4
		}
	}
	return result
}
