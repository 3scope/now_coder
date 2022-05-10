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

// 大根堆，因此将大的元素往前移。
func (this ArrayHeap) Less(i, j int) bool {
	return this.Data[i] > this.Data[j]
}

func (this ArrayHeap) Swap(i, j int) {
	this.Data[i], this.Data[j] = this.Data[j], this.Data[i]
}

// Push 的参数是空接口，需要进行断言。
func (ah *ArrayHeap) Push(x interface{}) {
	ah.Data = append(ah.Data, x.(int))
	ah.Length++
}

func (ah *ArrayHeap) Pop() interface{} {
	// 每次 Pop 先存最后一个值，之后将它删掉，因为 Heap 包会交换第一个值和最后一个值的位置。
	result := ah.Data[ah.Length-1]
	ah.Data = ah.Data[:ah.Length-1]
	ah.Length--

	return result
}

func GetLeastNumbers_Solution(input []int, k int) []int {
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
