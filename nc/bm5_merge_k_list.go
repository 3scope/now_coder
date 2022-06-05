package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeSortedList(lists, 0, len(lists)-1)
}

// 使用归并排序，“left”指向需要排序的第一个链表，“right”指向需要排序的最后一个链表。
func mergeSortedList(lists []*ListNode, left, right int) *ListNode {
	// 当只有一个链表需要排序的时候，直接返回。
	// 递归终止条件。
	if left == right {
		return lists[left]
	}

	// 二分归并，得到中间值。
	middle := (left + right) / 2
	// “leftList”和“rightList”都是有序的。
	leftList := mergeSortedList(lists, left, middle)
	rightList := mergeSortedList(lists, middle+1, right)

	// 进行归并。
	newHead := new(ListNode)
	cur := newHead
	for leftList != nil && rightList != nil {
		if leftList.Val >= rightList.Val {
			cur.Next = rightList
			cur = rightList
			rightList = rightList.Next
		} else {
			cur.Next = leftList
			cur = leftList
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
