package main

import (
	"strconv"
	"strings"
)

// 二叉树序列化。
func Serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	// 当序列化时，遇到了非空节点，就把左右孩子加进队列，而遇到空节点，那么就将其序列化，而忽略其左右孩子。
	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := ""

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				result += "#-"
			} else {
				result += strconv.Itoa(cur.Val) + "-"
				queue = append(queue, cur.Left)
				queue = append(queue, cur.Right)
			}
		}
	}
	// 最后一个“-”需要删除。
	result = result[:len(result)-1]

	return result
}

// 二叉树重建。
func Deserialize(s string) *TreeNode {
	if s == "#" {
		return nil
	}

	// 先得到每个节点的值，建立对应的节点。
	nodeData := strings.Split(s, "-")
	node := make([]*TreeNode, len(nodeData))
	// Create node.
	for i := 0; i < len(node); i++ {
		if nodeData[i] != "#" {
			node[i] = new(TreeNode)
			node[i].Val, _ = strconv.Atoi(nodeData[i])
		} else {
			node[i] = nil
		}
	}

	queue := make([]*TreeNode, 1)
	// 根节点先入队。
	queue[0] = node[0]
	// “index”存放当前节点左右孩子的下标。
	index := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if node[index] != nil {
			cur.Left = node[index]
			queue = append(queue, cur.Left)
		}
		if node[index+1] != nil {
			cur.Right = node[index+1]
			queue = append(queue, cur.Right)
		}
		// “index”每次自增2，刚好指向下个非空节点的左孩子。
		index += 2
	}

	return node[0]
}
