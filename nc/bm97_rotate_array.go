package main

func solveArray(n int, m int, a []int) []int {
	// write code here
	if n == 0 {
		return nil
	}
	m = m % n
	reverseArray(a, 0, len(a)-1)
	reverseArray(a, 0, m-1)
	reverseArray(a, m, len(a)-1)

	return a
}

func reverseArray(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
