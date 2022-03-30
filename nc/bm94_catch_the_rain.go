package main

func maxWater(arr []int) int64 {
	// write code here
	if len(arr) <= 2 {
		return 0
	}

	result := 0
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	stack := make([]int, 0, 1000)
	stack = append(stack, 0)
	for i := 1; i < len(arr); i++ {
		if arr[stack[len(stack)-1]] > arr[i] {
			stack = append(stack, i)
		} else if arr[stack[len(stack)-1]] == arr[i] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) >= 2 {
				floor := stack[len(stack)-1]
				if arr[floor] > arr[i] {
					break
				}
				stack = stack[:len(stack)-1]
				left := stack[len(stack)-1]
				right := i

				h := min(arr[left], arr[right]) - arr[floor]
				w := right - left - 1
				result += h * w
			}
			if arr[stack[len(stack)-1]] < arr[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}

	return int64(result)
}
