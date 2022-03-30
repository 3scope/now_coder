package main

func maxArea(height []int) int {
	// write code here
	if len(height) < 2 {
		return 0
	}

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
		// Only moving the short side can increase the total area.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
