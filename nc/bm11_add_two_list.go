package main

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)

	for node := head1; node != nil; node = node.Next {
		stack1 = append(stack1, node)
	}
	for node := head2; node != nil; node = node.Next {
		stack2 = append(stack2, node)
	}

	// 头插法生成一个新的链表。
	newHead := new(ListNode)
	carry := 0
	for len(stack1) != 0 && len(stack2) != 0 {
		cur1 := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		cur2 := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]

		sum := cur1.Val + cur2.Val + carry
		carry = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		node.Next = newHead.Next
		newHead.Next = node
	}

	for len(stack1) != 0 {
		cur := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		sum := carry + cur.Val
		carry = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		node.Next = newHead.Next
		newHead.Next = node
	}
	for len(stack2) != 0 {
		cur := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]

		sum := carry + cur.Val
		carry = sum / 10
		node := new(ListNode)
		node.Val = sum % 10
		node.Next = newHead.Next
		newHead.Next = node
	}

	if carry != 0 {
		node := new(ListNode)
		node.Val = carry
		node.Next = newHead.Next
		newHead.Next = node
	}
	return newHead.Next
}
