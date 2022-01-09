package main

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	head := new(ListNode)
	temp := head

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			temp.Next = pHead1
			temp = pHead1
			pHead1 = pHead1.Next
		} else {
			temp.Next = pHead2
			temp = pHead2
			pHead2 = pHead2.Next
		}
	}
	if pHead1 == nil {
		temp.Next = pHead2
	} else {
		temp.Next = pHead1
	}
	// If the list had no head pointer, then new one, and drop it when return the result.
	return head.Next
}
