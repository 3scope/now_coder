package main

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}

	return hasPathSumBacktracking(root, sum, root.Val)
}

func hasPathSumBacktracking(root *TreeNode, sum int, curSum int) bool {
	if sum == curSum && root != nil && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		curSum += root.Left.Val
		if hasPathSumBacktracking(root.Left, sum, curSum) {
			return true
		}
		curSum -= root.Left.Val
	}

	if root.Right != nil {
		curSum += root.Right.Val
		if hasPathSumBacktracking(root.Right, sum, curSum) {
			return true
		}
		curSum -= root.Right.Val
	}

	return false
}
