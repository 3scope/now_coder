package main

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}

	oddHead := new(ListNode)
	oddHead.Next = head
	evenHead := new(ListNode)
	evenHead.Next = head.Next

	odd := head
	even := head.Next

	// The odd and even pointer is the tail of each list.
	// The even is behind of odd, so in each iteration it only judges whether even is nil.
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = even.Next
		even.Next = odd.Next
		even = odd.Next
	}
	odd.Next = evenHead.Next

	return oddHead.Next
}
