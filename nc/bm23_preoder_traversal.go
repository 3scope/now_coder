package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// write code here
	result := make([]int, 0)
	binaryPreoder(root, &result)

	return result
}

func binaryPreoder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	binaryPreoder(root.Left, result)
	binaryPreoder(root.Right, result)
}
