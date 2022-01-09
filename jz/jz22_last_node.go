package main

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	fast, slow := pHead, pHead

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil && slow != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
