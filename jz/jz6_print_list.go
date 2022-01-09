package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Iteration.
func PrintListFromTailToHead(head *ListNode) []int {
	result := []int{}
	for node := head; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// Recursion.
func PrintListRecursion(head *ListNode) []int {
	var result []int
	if head == nil {
		return result
	}

	result = PrintListRecursion(head.Next)

	result = append(result, head.Val)

	return result
}
