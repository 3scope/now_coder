package main

func isPail(head *ListNode) bool {
	// write code here
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next
	stack := make([]*ListNode, 1)
	stack[0] = slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		stack = append(stack, slow)
	}

	// When the node count is odd, the middle node should be deleted.
	if fast == nil {
		stack = stack[:len(stack)-1]
	}
	slow = slow.Next
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i].Val != slow.Val {
			return false
		} else {
			slow = slow.Next
		}
	}
	return true
}
