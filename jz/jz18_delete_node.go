package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @param val int整型
 * @return ListNode类
 */
func DeleteNode(head *ListNode, val int) *ListNode {
	node := head
	var pre *ListNode = nil
	for node != nil {
		if node.Val == val {
			if pre == nil {
				head = node.Next
			} else {
				pre.Next = node.Next
			}
		}
		pre = node
		node = node.Next
	}
	return head
}
