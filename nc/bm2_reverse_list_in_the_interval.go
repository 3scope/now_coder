package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 建立新的头节点。
	newHead := new(ListNode)
	newHead.Next = head

	// 定位对应的位置，“pre”指向第一个需要翻转节点的前驱，“cur”指向第一个需要翻转的节点。
	var pre, cur *ListNode
	pre = newHead
	cur = head
	for i := 0; i < m-1; i++ {
		pre = cur
		cur = cur.Next
	}

	// 使用头插法，“cur”节点是尾节点。
	for i := m; i < n; i++ {
		// “temp”存储需要翻转的节点。
		temp := cur.Next
		cur.Next = temp.Next
		// 头插法。
		temp.Next = pre.Next
		pre.Next = temp
	}

	return newHead.Next
}
