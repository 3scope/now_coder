package main

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
		// 当两个指针不同时，“cur1”和“cur2”各自前进一步。
		if cur1 == nil {
			// 当“cur1”遍历完第一个链表后，去遍历第二个链表。
			cur1 = pHead2
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = pHead1
		} else {
			cur2 = cur2.Next
		}

	}
	return cur1
}
