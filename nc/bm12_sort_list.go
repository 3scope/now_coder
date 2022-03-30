package main

func sortInList(head *ListNode) *ListNode {
	// write code here
	return mergeSortList(head)
}

func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Prevent an infinite loop of two nodes.
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	left := head
	right := slow.Next
	// Divide it into two list.
	slow.Next = nil

	left = mergeSortList(left)
	right = mergeSortList(right)

	// Merge two lists.
	newHead := new(ListNode)
	cur := newHead
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			cur = cur.Next
			left = left.Next
		} else {
			cur.Next = right
			cur = cur.Next
			right = right.Next
		}
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return newHead.Next
}
