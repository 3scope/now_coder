package main

func mergeArray(A []int, m int, B []int, n int) {
	// 可以保证“A”有足够的空间放下“B”，即“A”的长度为“m+n”。
	// 从后往前遍历，可以不使用额外的空间。
	pointerA := m - 1
	pointerB := n - 1
	// 指向需要存放的位置。
	index := m + n - 1
	for pointerA >= 0 && pointerB >= 0 {
		// 当“A”数组的值大于“B”数组时，此时应该将“A”数组的值放入结果中。
		if A[pointerA] >= B[pointerB] {
			A[index] = A[pointerA]
			pointerA--
		} else {
			A[index] = B[pointerB]
			pointerB--
		}
		index--
	}

	// for pointerA >= 0 {
	// 	A[index] = A[pointerA]
	// }
	// 如果“B”数组没有放完，则将之后所有的数据放入“A”数组中。
	for pointerB >= 0 {
		A[index] = B[pointerB]
		index--
		pointerB--
	}
}
