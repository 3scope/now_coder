package main

func Print(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 1)
	queue[0] = pRoot
	level := 0

	for len(queue) != 0 {
		level++
		size := len(queue)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// Zigzag.
			if size%2 == 1 {
				temp[i] = cur.Val
			} else {
				temp[size-i-1] = cur.Val
			}

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
