package main

func sortInList(head *ListNode) *ListNode {
	// write code here
	return mergeSortList(head)
}

// 参考归并排序的思想。
func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 遍历得到链表的中点。
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	left := head
	right := slow.Next
	// 分割成两个链表。
	slow.Next = nil

	// 进行归并。
	left = mergeSortList(left)
	right = mergeSortList(right)
	// 合并两个链表。
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
