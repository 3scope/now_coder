package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func threeOrders(root *TreeNode) [][]int {
	result := make([][]int, 0, 3)
	preOrder := PreOrderTraversal(root)
	inOrder := InOrderTraversal(root)
	postOrder := PostOrderTraversal(root)

	result = append(result, preOrder)
	result = append(result, inOrder)
	result = append(result, postOrder)
	return result
}

// Recursion.
func (tn *TreeNode) PreOrderTraversal(root *TreeNode, temp *[]int) {
	if root == nil {
		return
	}
	*temp = append(*temp, root.Val)
	tn.PreOrderTraversal(root.Left, temp)
	tn.PreOrderTraversal(root.Right, temp)
}

func (tn *TreeNode) InOrderTraversal(root *TreeNode, temp *[]int) {
	if root == nil {
		return
	}
	tn.InOrderTraversal(root.Left, temp)
	*temp = append(*temp, root.Val)
	tn.InOrderTraversal(root.Right, temp)
}

func (tn *TreeNode) PostOrderTraversal(root *TreeNode, temp *[]int) {
	if root == nil {
		return
	}
	tn.PostOrderTraversal(root.Left, temp)
	tn.PostOrderTraversal(root.Right, temp)
	*temp = append(*temp, root.Val)
}

// Iterate.
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0, 8)
	stack = append(stack, root)
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		result = append(result, root.Val)
		stack = stack[0 : len(stack)-1]
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return result
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0, 8)

	stack = append(stack, root)
	for root.Left != nil {
		root = root.Left
		stack = append(stack, root)
	}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		result = append(result, root.Val)
		stack = stack[0 : len(stack)-1]
		if root.Right != nil {
			stack = append(stack, root.Right)
			root = root.Right
			for root.Left != nil {
				root = root.Left
				stack = append(stack, root)
			}
		}
	}
	return result
}

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0, 8)
	stack = append(stack, root)
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		result = append(result, root.Val)
		stack = stack[0 : len(stack)-1]
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
