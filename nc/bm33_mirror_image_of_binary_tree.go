package main

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	PreorderTraversal(pRoot)

	return pRoot
}

func PreorderTraversal(pRoot *TreeNode) {
	if pRoot == nil {
		return
	}

	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	PreorderTraversal(pRoot.Left)
	PreorderTraversal(pRoot.Right)
}
