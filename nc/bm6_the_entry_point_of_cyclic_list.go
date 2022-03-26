package main

func hasCycle(head *ListNode) bool {
	// write code here
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
			// slow = head
			// for slow != fast {
			// 	slow = slow.Next
			// 	fast = fast.Next
			// }
			// return fast
		}
	}
	// If fast equals to false, there is no cycle.
	return false
}
