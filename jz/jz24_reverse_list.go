package main

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */

// 迭代
func ReverseList(pHead *ListNode) *ListNode {
	node := pHead
	var pre *ListNode = nil
	var next *ListNode = nil

	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

// Recursion.
func ReverseListRecursion(pHead *ListNode) *ListNode {
	var result *ListNode = nil
	// The recursion termination condition.
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	result = ReverseListRecursion(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil

	return result
}
