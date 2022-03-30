package main

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}

	pre := head
	cur := head.Next

	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return head
}
