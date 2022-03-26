package main

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
		// Logical splicing.
		if cur1 == nil {
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
