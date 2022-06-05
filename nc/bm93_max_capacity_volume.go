package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	// 贪心算法，使用双指针从最外层向内部逼近。
	left := 0
	right := len(height) - 1
	result := 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	for left < right {
		sum := min(height[left], height[right]) * (right - left)
		result = max(result, sum)
		// 当一边比另一边要矮时，那么就需要移动矮的一边，因为这样才能
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
