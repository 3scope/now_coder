package main

func deleteDuplicatesTwo(head *ListNode) *ListNode {
	// write code here
	newHead := new(ListNode)
	newHead.Next = head
	pre := newHead
	cur := head

	dupValue := 0
	for cur != nil && cur.Next != nil {
		// The "cur" pointer to the last element of duplicates.
		if cur.Val == cur.Next.Val {
			// Delete all the nodes.
			dupValue = cur.Val
			for cur.Next != nil && cur.Next.Val == dupValue {
				cur = cur.Next
			}
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return newHead.Next
}
