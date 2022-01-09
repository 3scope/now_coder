package main

func EntryNodeOfLoopDoublePointer(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	fast := pHead
	slow := pHead

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = pHead
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	// When break the loop, there is no iteration in link.
	return nil
}
