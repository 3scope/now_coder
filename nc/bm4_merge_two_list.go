package main

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// 存储最终结果。
	newHead := new(ListNode)
	// 指向结果的最后一个节点。
	cur := newHead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val > pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		}
		// 尾指针更新。
		cur = cur.Next
	}

	if pHead1 == nil {
		cur.Next = pHead2
	} else if pHead2 == nil {
		cur.Next = pHead1
	}

	return newHead.Next
}
