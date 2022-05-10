package main

import (
	"strconv"
	"strings"
)

func Serialize(root *TreeNode) string {
	// write code here
	if root == nil {
		return "#"
	}

	// The queue only holds non-empty nodes and the next level of it.
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
	// Delete the last char.
	result = result[:len(result)-1]

	return result
}

func Deserialize(s string) *TreeNode {
	// write code here
	nodeData := strings.Split(s, "-")
	node := make([]*TreeNode, len(nodeData))
	if len(node) < 1 || nodeData[0] == "#" {
		return nil
	}

	// Create node.
	for i := 0; i < len(node); i++ {
		if nodeData[i] != "#" {
			node[i] = new(TreeNode)
			node[i].Val, _ = strconv.Atoi(nodeData[i])
		} else {
			node[i] = nil
		}
	}

	// Create binary tree.
	queue := make([]*TreeNode, 1)
	// Store the root.
	queue[0] = node[0]
	for i := 1; i < len(node); i += 2 {
		cur := queue[0]
		queue = queue[1:]
		if node[i] != nil {
			queue = append(queue, node[i])
		}
		cur.Left = node[i]
		if node[i+1] != nil {
			queue = append(queue, node[i+1])
		}
		cur.Right = node[i+1]
	}

	return node[0]
}
