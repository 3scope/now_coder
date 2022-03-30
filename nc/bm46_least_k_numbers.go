package main

import (
	"container/heap"
)

type ArrayHeap struct {
	Data   []int
	Length int
}

func (this ArrayHeap) Len() int {
	return this.Length
}

// Big root heap.
func (this ArrayHeap) Less(i, j int) bool {
	return this.Data[i] > this.Data[j]
}

func (this ArrayHeap) Swap(i, j int) {
	this.Data[i], this.Data[j] = this.Data[j], this.Data[i]
}

func (this *ArrayHeap) Push(x interface{}) {
	this.Data = append(this.Data, x.(int))
	this.Length++
}

func (this *ArrayHeap) Pop() interface{} {
	// Swap the element of index len(old) and 0.
	result := this.Data[this.Length-1]
	this.Data = this.Data[:this.Length-1]
	this.Length--

	return result
}

func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	h := &ArrayHeap{}
	heap.Init(h)
	for i := 0; i < len(input); i++ {
		heap.Push(h, input[i])
		if h.Length > k {
			heap.Pop(h)
		}
	}

	return h.Data
}
