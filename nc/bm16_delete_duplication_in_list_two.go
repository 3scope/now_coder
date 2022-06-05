package main

func deleteDuplicatesTwo(head *ListNode) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	// "pre"指针指向需要保留节点中的最后一个节点。
	pre := newHead
	cur := head

	for cur != nil && cur.Next != nil {
		// 如果出现重复的情况，需要考虑删除。
		if cur.Val == cur.Next.Val {
			// 找到最后一个重复值。
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			// 没有出现重复，保留当前节点。
			pre = cur
			cur = cur.Next
		}
	}

	return newHead.Next
}
