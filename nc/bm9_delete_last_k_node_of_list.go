package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	newHead := new(ListNode)
	newHead.Next = head

	slow, fast := newHead, newHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
