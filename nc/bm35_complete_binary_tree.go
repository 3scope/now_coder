package main

func isCompleteTree(root *TreeNode) bool {
	// write code here
	if root == nil {
		return true
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root
	isNil := false
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				isNil = true
				continue
			}
			// Not nil node appears after nil.
			if isNil {
				return false
			}
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}

	return true
}
