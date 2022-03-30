package main

func postorderTraversal(root *TreeNode) []int {
	// write code here
	if root == nil {
		return nil
	}

	result := make([]int, 0, 1)
	stack := make([]*TreeNode, 1)
	stack[0] = root

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
