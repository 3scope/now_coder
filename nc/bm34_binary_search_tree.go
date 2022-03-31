package main

func isValidBST(root *TreeNode) bool {
	// write code here
	support := make([]*TreeNode, 1)
	return isValidBSTInorder(root, support)
}

func isValidBSTInorder(root *TreeNode, support []*TreeNode) bool {
	if root == nil {
		return true
	}

	// Inorder traversal.
	if !isValidBSTInorder(root.Left, support) {
		return false
	}

	if support[0] == nil {
		support[0] = root
	} else {
		if support[0].Val >= root.Val {
			return false
		}
		support[0] = root
	}

	if !isValidBSTInorder(root.Right, support) {
		return false
	}
	return true
}
