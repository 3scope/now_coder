package main

func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	middle := (left + right) / 2
	leftList := mergeLists(lists, left, middle)
	rightList := mergeLists(lists, middle+1, right)

	newHead := new(ListNode)
	cur := newHead
	for leftList != nil && rightList != nil {
		if leftList.Val > rightList.Val {
			cur.Next = rightList
			cur = cur.Next
			rightList = rightList.Next
		} else {
			cur.Next = leftList
			cur = cur.Next
			leftList = leftList.Next
		}
	}

	if leftList != nil {
		cur.Next = leftList
	}
	if rightList != nil {
		cur.Next = rightList
	}

	return newHead.Next
}
