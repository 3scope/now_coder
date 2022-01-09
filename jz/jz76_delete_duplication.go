package main

func DeleteDuplication(pHead *ListNode) *ListNode {
	nHead := new(ListNode)
	nHead.Next = pHead

	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	ptr := pHead
	pre := nHead
	for ptr != nil {
		if ptr.Next == nil {
			break
		}
		if ptr.Next.Val == ptr.Val {
			ptr = ptr.Next
			for ptr.Next != nil && ptr.Next.Val == ptr.Val {
				ptr = ptr.Next
			}
			ptr = ptr.Next
			pre.Next = ptr
		} else {
			pre = ptr
			ptr = ptr.Next
		}
	}
	return nHead.Next
}
