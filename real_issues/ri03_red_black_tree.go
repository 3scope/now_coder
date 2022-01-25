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

	rbt.Balance(newNode)
}

func (rbt *RBTree) Balance(p *RBNode) {
	if p.Parent.Color != Red {
		return
	}
	if p.Parent.Parent.Left == p.Parent {

	} else {

	}
}
