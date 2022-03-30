package main

func levelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := make([][]int, 0)
	for len(queue) != 0 {
		size := len(queue)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			temp[i] = cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, temp)
	}

	return result
}
