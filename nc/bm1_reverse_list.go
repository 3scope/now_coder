package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var pre, cur *ListNode
	pre = nil
	cur = pHead
	for cur != nil {
		// Use temp to store the next pointer of current.
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}
