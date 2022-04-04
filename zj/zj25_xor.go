package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count, target := 0, 0
	fmt.Scanf("%d %d\n", &count, &target)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	numsStr := strings.Split(input.Text(), " ")
	nums := make([]int, len(numsStr))
	for i := 0; i < len(nums); i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}
	targetSlice := make([]int, 14)
	for i := 13; i >= 0; i-- {
		bit := target >> i & 1
		targetSlice[i] = bit
	}

	result := 0
	root := &Trie{
		Next: make(map[int]*Trie),
	}
	for i := 1; i < len(nums); i++ {
		root.Insert(nums[i-1])
		result += root.IsGreaterThan(nums[i], targetSlice)
	}
	fmt.Println(result)
}

type Trie struct {
	IsEnd bool
	Next  map[int]*Trie
}

func (root *Trie) Insert(value int) {
	// The max length of value is 14 bit.
	maxLength := 13
	cur := root
	for i := maxLength; i >= 0; i-- {
		bit := value >> i & 1
		if cur.Next[bit] == nil {
			temp := &Trie{
				Next: make(map[int]*Trie),
			}
			cur.Next[bit] = temp
		}
		cur = cur.Next[bit]
	}
	cur.IsEnd = true
}

func (root *Trie) IsGreaterThan(num int, targetSlice []int) int {
	maxLength := 13
	cur := root
	subResult := 0
	for i := maxLength; i >= 0; i-- {
		bit := num >> i & 1
		if targetSlice[i] == 1 && cur.Next[1^bit] == nil {
			return subResult
		} else if targetSlice[i] == 0 && cur.Next[1^bit] != nil {
			subResult += getLeaf(cur.Next[1^bit])
			if cur.Next[bit] != nil {
				cur = cur.Next[bit]
			} else {
				return subResult
			}
		} else if cur.Next[1^bit] == nil {
			cur = cur.Next[bit]
		} else {
			cur = cur.Next[1^bit]
		}
	}
	return subResult
}

func getLeaf(root *Trie) int {
	if root.IsEnd {
		return 1
	}

	count := 0
	for _, child := range root.Next {
		count += getLeaf(child)
	}
	return count
}
