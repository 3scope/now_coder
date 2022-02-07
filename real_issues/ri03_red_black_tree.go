package main

type NodeColor uint16

const (
	Red   NodeColor = 1
	Black NodeColor = 2
)

type RBNode struct {
	Color NodeColor
	Key   int
	Value int

	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
}

type RBTree struct {
	Root *RBNode
}

func (rbt *RBTree) LeftRotation(p *RBNode) {
	if p == nil {
		return
	}

	rightChild := p.Right
	if rightChild != nil {
		p.Right = rightChild.Left
		rightChild.Left.Parent = p
		rightChild.Left = p
	} else {
		return
	}

	rightChild.Parent = p.Parent
	if p.Parent == nil {
		p.Parent = rightChild
		rbt.Root = rightChild
		return
	}
	if p.Parent.Left == p {
		p.Parent.Left = rightChild
		p.Parent = rightChild
	} else if p.Parent.Right == p {
		p.Parent.Right = rightChild
		p.Parent = rightChild
	}
}

func (rbt *RBTree) RightRotation(p *RBNode) {
	if p == nil {
		return
	}

	leftChild := p.Left
	if leftChild != nil {
		p.Left = leftChild.Right
		leftChild.Right.Parent = p
		leftChild.Right = p
	} else {
		return
	}
	leftChild.Parent = p.Parent
	if p.Parent == nil {
		p.Parent = leftChild
		rbt.Root = leftChild
		return
	}
	if p.Parent.Left == p {
		p.Parent.Left = leftChild
		p.Parent = leftChild
	} else if p.Parent.Right == p {
		p.Parent.Right = leftChild
		p.Parent = leftChild
	}
}

func (rbt *RBTree) InsertNode(k, v int) {
	if rbt.Root == nil {
		rbt.Root = &RBNode{
			Key:   k,
			Value: v,
			Color: Black,
		}
		return
	}

	current := rbt.Root
	var pre *RBNode
	for current != nil {
		pre = current
		if k > current.Key {
			current = current.Right
		} else if k < current.Key {
			current = current.Left
		} else {
			current.Value = v
			return
		}
	}
	newNode := &RBNode{
		Key:   k,
		Value: v,
		Color: Red,
	}
	if k > pre.Key {
		pre.Right = newNode
		newNode.Parent = pre
	} else {
		pre.Left = newNode
		newNode.Parent = pre
	}

	rbt.InsertBalance(newNode)
}

// The node p is a just modified node that needs to be balanced.
func (rbt *RBTree) InsertBalance(p *RBNode) {
	if p.Parent.Color != Red {
		return
	}
	// If the modified node p is root node.
	if p.Parent == nil {
		p.Color = Black
		return
	}

	// If the parent node leans to the left.
	if p.Parent.Parent.Left == p.Parent {
		parent := p.Parent
		// Sibling node(uncle node) is red.
		if parent.Parent.Right != nil {
			uncle := parent.Parent.Right
			parent.Color = Black
			uncle.Color = Black
			parent.Parent.Color = Red
			rbt.InsertBalance(parent.Parent)
			return
		}
		if parent.Left == p {
			parent.Color = Black
			parent.Parent.Color = Red
			rbt.RightRotation(parent.Parent)
			return
		} else {
			parent.Parent.Color = Red
			p.Color = Black
			rbt.LeftRotation(parent)
			rbt.RightRotation(p.Parent)
			return
		}
	} else {
		parent := p.Parent
		// Sibling node(uncle node) is red.
		if parent.Parent.Left != nil {
			uncle := parent.Parent.Left
			parent.Color = Black
			uncle.Color = Black
			parent.Parent.Color = Red
			rbt.InsertBalance(parent.Parent)
			return
		}
		if parent.Right == p {
			parent.Color = Black
			parent.Parent.Color = Red
			rbt.LeftRotation(parent.Parent)
			return
		} else {
			parent.Parent.Color = Red
			p.Color = Black
			rbt.RightRotation(parent)
			rbt.LeftRotation(p.Parent)
			return
		}
	}
}

func (rbt *RBTree) DeleteNode(key int) {
	node := rbt.SearchNode(key)
	rbt.DeleteAndBalance(node)
}

func (rbt *RBTree) SearchNode(key int) *RBNode {
	// Leaf node.
	p := rbt.Root
	for p != nil {
		if p.Key == key {
			break
		} else if p.Key < key {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return p
}

func (rbt *RBTree) DeleteAndBalance(p *RBNode) {
	// Leaf node.
	if p.Right == nil && p.Left == nil {
		// If delete the root node, set the root pointer into nil.
		if rbt.Root == p {
			rbt.Root = nil
			return
		}

		// Adjust first, Then delete node.
		if p.Color == Black {
			rbt.DeleteBalance(p)
		}
		parent := p.Parent
		if parent.Left == p {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		p.Parent = nil
		return

	} else if p.Left != nil && p.Right != nil {
		// It has two children.
		successor := p.Right
		pre := (*RBNode)(nil)
		for successor != nil {
			pre = successor
			successor = successor.Left
		}
		// The successor of p node.
		successor = pre
		p.Key = successor.Key
		p.Value = successor.Value

		// Removed nodes are moved to the leaf of 234 tree.
		rbt.DeleteAndBalance(successor)
	} else if p.Left == nil {
		// Delete first, then adjust.
		// It has no left child.
		// It is root node.
		replacement := p.Right
		if p.Parent == nil {
			rbt.Root = replacement
		} else {
			if p.Parent.Left == p {
				p.Parent.Left = replacement
			} else {
				p.Parent.Right = replacement
			}
		}
		replacement.Parent = p.Parent
		p.Parent = nil
		p.Right = nil
		// Actually, the node has only one child can be only black color.
		if p.Color == Black {
			rbt.DeleteBalance(replacement)
		}
		return
	} else {
		// Delete first, the adjust.
		// It has no right child.
		// It is root node.
		replacement := p.Left
		if p.Parent == nil {
			rbt.Root = replacement
		} else {
			if p.Parent.Left == p {
				p.Parent.Left = replacement
			} else {
				p.Parent.Right = replacement

			}
		}
		replacement.Parent = p.Parent
		p.Parent = nil
		p.Left = nil
		// Actually, the node has only one child can be only black color.
		if p.Color == Black {
			rbt.DeleteBalance(replacement)
		}
		return
	}
}

func (rbt *RBTree) DeleteBalance(p *RBNode) {
	// In the first case, the deleted node has a child, and the child must be red.
	if p.Color == Red {
		p.Color = Black
		return
		// The color of the node is black. That's mean the delete node is leaf.
	} else {
		// The leaf node is the left child.
		if p.Parent.Left == p {
			right := p.Parent.Right
			// To judge whether the right node is the sibling node.
			if right.Color == Red {
				// In this condition, the right node is not the sibling node of p in 234 tree.
				right.Color = Black
				right.Parent.Color = Red
				rbt.LeftRotation(p.Parent)
				// The action above can change the tree into another status of balance.
			}
			if right.Left == nil && right.Right == nil {
				right.Color = Red

				// In this situation, the sibling node has extra child to lend to p.
			} else if right.Right == nil && right.Left != nil {
				right.Left.Color = Black         //			Red 7					 Black 8
				right.Color = right.Parent.Color //		   /	\					/	  \
				rbt.RightRotation(right)         // X Black 6(p) Black 9(right) => Red 7  Red 9
				rbt.LeftRotation(p.Parent)       //				/
			} else { // 									Red 8
				right.Color = right.Parent.Color
				right.Right.Color = Black
				right.Parent.Color = Black
				rbt.LeftRotation(p.Parent)
			}
		} else {
			left := p.Parent.Left
			if left.Color == Red {
				left.Color = Black
				left.Parent.Color = Red
				rbt.RightRotation(p.Parent)
			}
		}
	}
	//
}

// func (rbt *RBTree) DirectDelete(p *RBNode) {
// 	// Leaf node.
// 	if p.Right == nil && p.Left == nil {
// 		// If delete the root node, set the root pointer into nil.
// 		if rbt.Root == p {
// 			rbt.Root = nil
// 			return
// 		}

// 		parent := p.Parent
// 		if parent.Left == p {
// 			parent.Left = nil
// 		} else {
// 			parent.Right = nil
// 		}
// 		p.Parent = nil
// 		return
// 	} else if p.Left != nil && p.Right != nil {
// 		// It has two children.
// 		successor := p.Right
// 		pre := (*RBNode)(nil)
// 		for successor != nil {
// 			pre = successor
// 			successor = successor.Left
// 		}
// 		// The successor of p node.
// 		successor = pre
// 		p.Key = successor.Key
// 		p.Value = successor.Value
// 		rbt.DirectDelete(successor)
// 	} else if p.Left == nil {
// 		// It has no left child.
// 		// It is root node.
// 		if p.Parent == nil {
// 			rbt.Root = p.Right
// 		} else {
// 			if p.Parent.Left == p {
// 				p.Parent.Left = p.Right
// 			} else {
// 				p.Parent.Right = p.Right
// 			}
// 		}
// 		p.Right.Parent = p.Parent
// 		p.Parent = nil
// 		p.Right = nil
// 		return
// 	} else {
// 		// It has no right child.
// 		// It is root node.
// 		if p.Parent == nil {
// 			rbt.Root = p.Left
// 		} else {
// 			if p.Parent.Left == p {
// 				p.Parent.Left = p.Right
// 			} else {
// 				p.Parent.Right = p.Right

// 			}
// 		}
// 		p.Left.Parent = p.Parent
// 		p.Parent = nil
// 		p.Left = nil
// 		return
// 	}
// }
