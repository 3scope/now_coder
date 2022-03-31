package main

// Store the minimum number, if a new one if smaller than than the root, it should be pushed into max heap.
type IntArrayMaxHeap struct {
	Data   []int
	Length int
}

func (this IntArrayMaxHeap) Len() int {
	return this.Length
}

func (this IntArrayMaxHeap) Less(i, j int) bool {
	return this.Data[i] > this.Data[j]
}

func (this IntArrayMaxHeap) Swap(i, j int) {
	this.Data[i], this.Data[j] = this.Data[j], this.Data[i]
}

func (this *IntArrayMaxHeap) Push(x interface{}) {
	this.Data = append(this.Data, x.(int))
	this.Length++
}

func (this *IntArrayMaxHeap) Pop() interface{} {
	result := this.Data[this.Length-1]
	this.Data = this.Data[:this.Length-1]

	return result
}

// var middleNumberArray = make([]int, 0)

// func Insert(num int) {
// 	middleNumberArray = append(middleNumberArray, num)
// 	// Insert sort.
// 	index := len(middleNumberArray) - 2
// 	for ; index >= 0; index-- {
// 		if middleNumberArray[index] <= num {
// 			break
// 		} else {
// 			// Move to the next backward.
// 			middleNumberArray[index+1] = middleNumberArray[index]
// 		}
// 	}
// 	middleNumberArray[index+1] = num
// }

// func GetMedian() float64 {
// 	length := len(middleNumberArray)
// 	if length%2 == 1 {
// 		return float64(middleNumberArray[length/2])
// 	} else {
// 		left := length/2 - 1
// 		right := length / 2
// 		return (float64(middleNumberArray[left]) + float64(middleNumberArray[right])) / 2
// 	}
// }
