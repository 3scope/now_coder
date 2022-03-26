package main

// The pre is the head node of the reverse interval.
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
	// write code here
	if k == 1 {
		return head
	}
	newHead := new(ListNode)
	newHead.Next = head

	pre := newHead
	cur := head

	for reverseKNode(pre, cur, k) {
		for i := 0; i < k; i++ {
			pre = pre.Next
		}
		cur = pre.Next
	}
	// When there is no left k nodes, the rest is revesed, so it should reverse again.
	cur = pre.Next
	reverseKNode(pre, cur, k)

	return newHead.Next
}

// func main() {
// 	nodeArray := []int{
// 		1, 2, 3, 4, 5,
// 	}
// 	nodeList := make([]*ListNode, len(nodeArray))
// 	k := 3
// 	for i := 0; i < len(nodeArray); i++ {
// 		nodeList[i] = &ListNode{nodeArray[i], nil}
// 	}

// 	for i := 0; i < len(nodeList)-1; i++ {
// 		nodeList[i].Next = nodeList[i+1]
// 	}
// 	reverseKGroup(nodeList[0], k)
// }
