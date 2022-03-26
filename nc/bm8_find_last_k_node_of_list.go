package main

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	newHead := new(ListNode)
	newHead.Next = pHead

	slow, fast := newHead, newHead
	for i := 0; fast != nil && i < k; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
