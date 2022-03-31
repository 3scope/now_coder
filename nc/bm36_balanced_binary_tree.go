package main

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if IsBalancedSolution(pRoot) == -1 {
		return false
	}
	return true
}

func IsBalancedSolution(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Postorder.
	// Left child.
	left := IsBalancedSolution(root.Left)
	if left == -1 {
		return -1
	}
	// Right child.
	right := IsBalancedSolution(root.Right)
	if right == -1 {
		return -1
	}

	// Middle.
	sub := left - right
	if sub < -1 || sub > 1 {
		return -1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	return max(left, right) + 1
}
