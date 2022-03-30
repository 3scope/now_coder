package main

type Solution struct {
	// write code here
	Capacity int
	Head     *DoubleListNode
	Tail     *DoubleListNode
	// Use for get method
	KeyNode map[int]*DoubleListNode
}

type DoubleListNode struct {
	Key   int
	Value int
	Pre   *DoubleListNode
	Next  *DoubleListNode
}

func Constructor(capacity int) Solution {
	// write code here
	solution := Solution{
		Capacity: capacity,
		Head: &DoubleListNode{
			Pre:  nil,
			Next: nil,
		},
		KeyNode: make(map[int]*DoubleListNode),
	}
	// Circular linked list.
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
	// If use head node and tail node, the next of head is always exists.
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
	// write code here
	if this.Capacity > 0 {
		this.Capacity--
		this.insertAfterHead(key, value)
	} else {
		this.deleteFromTail()
		this.insertAfterHead(key, value)
	}
}

/**
* Your Solution object will be instantiated and called as such:
* solution := Constructor(capacity);
* output := solution.get(key);
* solution.set(key,value);
 */
