package main

func maxDepth(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}
