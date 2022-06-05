package main

func isPail(head *ListNode) bool {
	// write code here
	if head == nil || head.Next == nil {
		return true
	}

	// 找中点，需要使用快慢指针，将“fast”初始化到“slow”后一个可以保证慢指针出现在前半段。
	slow, fast := head, head.Next
	stack := make([]*ListNode, 1)
	stack[0] = slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		stack = append(stack, slow)
	}

	// 当最后“fast”指针指向“nil”时，“slow”指针刚好位于中间位置，那么此时中间位置不需要进行比较，直接删除。
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
