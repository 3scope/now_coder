package main

func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	// Use inorder traversal.
	// Store the pre node and new root node.
	support := make([]*TreeNode, 2)
	InorderConvert(pRootOfTree, support)

	return support[1]
}

func InorderConvert(root *TreeNode, support []*TreeNode) {
	if root == nil {
		return
	}
	InorderConvert(root.Left, support)
	if support[0] == nil {
		support[0] = root
	} else {
		support[0].Right = root
		root.Left = support[0]
		support[0] = root
	}
	// Get the root of double link list.
	if support[1] == nil {
		support[1] = root
	}
	InorderConvert(root.Right, support)
}
