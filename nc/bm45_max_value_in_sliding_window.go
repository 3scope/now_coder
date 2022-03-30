package main

func maxInWindows(num []int, size int) []int {
	// write code here
	queue := &MonotonicQueue{
		Data:   make([]int, 0, 1),
		Length: 0,
	}
	// The first round.
	for i := 0; i < size; i++ {
		queue.Push(num[i])
	}
	result := []int{queue.Data[0]}
	for i := size; i < len(num); i++ {
		// To check if i-size is in the queue.
		queue.PopFront(num[i-size])
		queue.Push(num[i])
		result = append(result, queue.Data[0])
	}

	return result
}

type MonotonicQueue struct {
	// The front of the queue is at index of 0.
	Data   []int
	Length int
}

func (mq *MonotonicQueue) PopFront(value int) {
	if mq.Length > 0 && mq.Data[0] == value {
		mq.Data = mq.Data[1:]
		mq.Length--
	}
}

func (mq *MonotonicQueue) PopBack() {
	if mq.Length > 0 {
		mq.Data = mq.Data[:mq.Length-1]
		mq.Length--
	}
}

func (mq *MonotonicQueue) Push(value int) {
	// The max element is on the front.
	// When the value equals to tail element, push it in the queue.
	for mq.Length != 0 && mq.Data[mq.Length-1] < value {
		mq.PopBack()
	}
	mq.Data = append(mq.Data, value)
	mq.Length++
}
