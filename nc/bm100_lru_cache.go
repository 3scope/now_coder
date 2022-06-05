package main

type Solution struct {
	Capacity int
	Head     *DoubleListNode
	Tail     *DoubleListNode
	// “Get”方法的时间复制度要求是 O(1)，因此使用一个键值对存储。
	KeyNode map[int]*DoubleListNode
}

type DoubleListNode struct {
	Key   int
	Value int
	Pre   *DoubleListNode
	Next  *DoubleListNode
}

func Constructor(capacity int) Solution {
	solution := Solution{
		Capacity: capacity,
		Head: &DoubleListNode{
			Pre:  nil,
			Next: nil,
		},
		KeyNode: make(map[int]*DoubleListNode),
	}
	// 使用一个环形双链表，初始时只有一个节点。
	solution.Tail = solution.Head
	solution.Head.Next = solution.Tail
	solution.Tail.Pre = solution.Head

	return solution
}

func (this *Solution) insertAfterHead(key, value int) {
	node := &DoubleListNode{
		Key:   key,
		Value: value,
		Pre:   this.Head,
		Next:  this.Head.Next,
	}
	this.Head.Next.Pre = node
	this.Head.Next = node
	this.KeyNode[key] = node
}

func (this *Solution) deleteFromTail() {
	deleteNode := this.Tail.Pre
	pre := deleteNode.Pre

	deleteNode.Pre = nil
	pre.Next = this.Tail
	deleteNode.Next = nil
	this.Tail.Pre = pre
	delete(this.KeyNode, deleteNode.Key)
}

func (this *Solution) get(key int) int {
	// write code here
	node, ok := this.KeyNode[key]
	if !ok {
		return -1
	}

	// 首先通过“map”快速得到对应的节点，之后将将节点从链表上删除，再插入到头节点后面。
	pre := node.Pre
	next := node.Next
	node.Pre = nil
	node.Next = nil
	pre.Next = next
	next.Pre = pre
	this.insertAfterHead(node.Key, node.Value)

	return node.Value
}

func (this *Solution) set(key int, value int) {
	if this.Capacity > 0 {
		this.Capacity--
		this.insertAfterHead(key, value)
	} else {
		// 空间不够的话，需要先删除尾部的节点。
		this.deleteFromTail()
		this.insertAfterHead(key, value)
	}
}
