package main

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	return preorderSymmetical(pRoot.Left, pRoot.Right)
}

func preorderSymmetical(left, right *TreeNode) bool {
	// First to handle with middle nodes(two root nodes).
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	// After processing what the current node can get, to process left and right child.
	outside := preorderSymmetical(left.Left, right.Right)
	inside := preorderSymmetical(left.Right, right.Left)
	return outside && inside
}
