package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	level := 0
	flag := make(map[*TreeNode]int)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	flag[root] = level
	result = append(result, make([]int, 0))

	for len(queue) != 0 {
		root = queue[0]
		queue = queue[1:]
		if flag[root] != level {
			level++
			result = append(result, make([]int, 0))
		}
		result[level] = append(result[level], root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
			flag[root.Left] = flag[root] + 1
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
			flag[root.Right] = flag[root] + 1
		}
	}

	return result
}
