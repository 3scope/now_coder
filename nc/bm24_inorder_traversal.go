package main

func inorderTraversal(root *TreeNode) []int {
	// write code here
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0, 1)
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		for node := cur.Right; node != nil; node = node.Left {
			stack = append(stack, node)
		}
	}

	return result
}
