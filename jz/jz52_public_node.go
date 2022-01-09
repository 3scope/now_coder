package main

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	slice1 := make([]*ListNode, 0)
	slice2 := make([]*ListNode, 0)
	for ; pHead1 != nil; pHead1 = pHead1.Next {
		slice1 = append(slice1, pHead1)
	}
	for ; pHead2 != nil; pHead2 = pHead2.Next {
		slice2 = append(slice2, pHead2)
	}
	var public *ListNode = nil
	for i, j := len(slice1)-1, len(slice2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if slice1[i].Val != slice2[j].Val {
			return public
		}
		public = slice1[i]
	}
	return public
}

// Another solution.

func FindFirstCommonNodeDoublePoiner(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
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
