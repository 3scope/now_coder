package main

func MySort(arr []int) []int {
	return fastSort(arr, 0, len(arr)-1)
}

func fastSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	leftCurrent := left
	rightCurrent := right
	pivot := arr[left]

	for leftCurrent < rightCurrent {
		for leftCurrent < rightCurrent && arr[rightCurrent] >= pivot {
			rightCurrent--
		}
		// If the value at right point is less than pivot, then set it to the left point, and go on.
		if leftCurrent < rightCurrent {
			arr[leftCurrent] = arr[rightCurrent]
			leftCurrent++
		}
		for leftCurrent < rightCurrent && arr[leftCurrent] <= pivot {
			leftCurrent++
		}
		if leftCurrent < rightCurrent {
			arr[rightCurrent] = arr[leftCurrent]
			rightCurrent--
		}
	}

	// The pivot point.
	arr[leftCurrent] = pivot
	// Recursion.
	arr = fastSort(arr, left, leftCurrent-1)
	arr = fastSort(arr, rightCurrent+1, right)

	return arr
}
