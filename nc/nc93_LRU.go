package main

func LRU(operators [][]int, k int) []int {
	// write code here
	lru := NewLRU(k)
	result := make([]int, 0)
	for _, array := range operators {
		opt := array[0]
		if opt == 1 {
			lru.Set(array[1], array[2])
		} else if opt == 2 {
			result = append(result, lru.Get(array[1]))
		}
	}

	return result
}

type DoubleListNode struct {
	Key  int
	Val  int
	Pre  *DoubleListNode
	Next *DoubleListNode
}

type LRUCache struct {
	// The capacity of cache.(Same as "k")
	Capacity int
	Head     *DoubleListNode
	Tail     *DoubleListNode
	Keys     map[int]*DoubleListNode
}

func NewLRU(capacity int) *LRUCache {
	lru := &LRUCache{
		Capacity: capacity,
		Head:     new(DoubleListNode),
		Tail:     new(DoubleListNode),
		Keys:     make(map[int]*DoubleListNode),
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Pre = lru.Head

	return lru
}

func (lru *LRUCache) InsertIntoHead(node *DoubleListNode) {
	node.Next = lru.Head.Next
	node.Pre = lru.Head
	lru.Head.Next.Pre = node
	lru.Head.Next = node
	lru.Keys[node.Key] = node
}

func (lru *LRUCache) DeleteFromList(node *DoubleListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
	delete(lru.Keys, node.Key)
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.Keys[key]; ok {
		lru.DeleteFromList(node)
		lru.InsertIntoHead(node)
		return node.Val
	} else {
		return -1
	}
}

func (lru *LRUCache) Set(key, value int) {
	if node, ok := lru.Keys[key]; ok {
		node.Val = value
		lru.DeleteFromList(node)
		lru.InsertIntoHead(node)
		return
	}

	if lru.Capacity > 0 {
		node := new(DoubleListNode)
		node.Key = key
		node.Val = value
		lru.InsertIntoHead(node)
		lru.Capacity--
	} else {
		node := new(DoubleListNode)
		node.Key = key
		node.Val = value
		lru.InsertIntoHead(node)
		lru.DeleteFromList(lru.Tail.Pre)
	}
}
