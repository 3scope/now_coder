package main

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	newHead := new(ListNode)
	cur := newHead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val > pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		}
		cur = cur.Next
	}

	if pHead1 == nil {
		cur.Next = pHead2
	} else if pHead2 == nil {
		cur.Next = pHead1
	}

	return newHead.Next
}
