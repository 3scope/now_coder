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

	queue := make([]*TreeNode, 1)
	queue[0] = root
	result := ""
	// Difference with level order.
	// The "flag" is used to store whether all elements in this level is nil.
	// If flag equals to true, it means all elements is nil.
	flag := true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				result += "#-"
				queue = append(queue, nil)
				queue = append(queue, nil)

				continue
			}
			flag = false
			result += strconv.Itoa(cur.Val)
			result += "-"
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
		// To reset the flag.
		if flag {
			break
		}
		flag = true
	}
	// Delete the last char.
	result = result[:len(result)-1]

	return result
}

func Deserialize(s string) *TreeNode {
	// write code here
	nodeData := strings.Split(s, "-")
	node := make([]*TreeNode, len(nodeData))
	if nodeData[0] == "#" {
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
	for i := 0; i*2+1 < len(node); i++ {
		// Pass over nil.
		if node[i] == nil {
			continue
		}
		node[i].Left = node[i*2+1]
		if i*2+2 < len(node) {
			node[i].Right = node[i*2+2]
		}
	}
	return node[0]
}
