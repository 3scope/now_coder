package main

// 翻转“k”个节点，成功翻转返回“true”，不足“k”个节点返回“false”。
func reverseKNode(pre, cur *ListNode, k int) bool {
	for i := 1; i < k; i++ {
		if cur == nil || cur.Next == nil {
			return false
		}
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 如果“k”为1，相当于不进行翻转。
	if k == 1 {
		return head
	}
	newHead := new(ListNode)
	newHead.Next = head
	// 从第一组开始翻转。
	pre := newHead
	cur := head

	for reverseKNode(pre, cur, k) {
		// 每次成功翻转后，需要重新定位。
		for i := 0; i < k; i++ {
			pre = pre.Next
		}
		cur = pre.Next
	}
	// 最后一次翻转失败的话，需要重新翻转回来。
	cur = pre.Next
	reverseKNode(pre, cur, k)

	return newHead.Next
}

func main() {
	nodeArray := []int{
		1, 2, 3, 4, 5,
	}
	nodeList := make([]*ListNode, len(nodeArray))
	k := 3
	for i := 0; i < len(nodeArray); i++ {
		nodeList[i] = &ListNode{nodeArray[i], nil}
	}

	for i := 0; i < len(nodeList)-1; i++ {
		nodeList[i].Next = nodeList[i+1]
	}
	reverseKGroup(nodeList[0], k)
}
