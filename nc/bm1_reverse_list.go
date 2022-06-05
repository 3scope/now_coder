package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// 特殊情况。
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	var pre, cur *ListNode
	pre = nil
	cur = pHead
	for cur != nil {
		// 使用临时变量存储“Next”指针。
		temp := cur.Next
		// 翻转链表。
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}
