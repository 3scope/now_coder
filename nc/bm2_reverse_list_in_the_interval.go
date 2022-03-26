package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	newHead := new(ListNode)
	newHead.Next = head

	var pre, cur *ListNode
	pre = newHead
	cur = head
	for i := 0; i < m-1; i++ {
		pre = cur
		cur = cur.Next
	}

	// Insert to the head.
	// The pre pointer is the head pointer of the list which should be reverse.
	// The cur pointer is the tail pointer of the list which is reversed.
	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return newHead.Next
}
