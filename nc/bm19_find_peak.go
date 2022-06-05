package main

// 峰值要求严格大于左右元素，并且当只有一个元素的时候，就是峰值。
func findPeakElement(nums []int) int {
	// 使用左闭右闭区间，每次比较“middle”所在的元素和右边的一个元素。
	// 因为每次都向下取整，那么只要“left”和“right”不相等，“middle+1”永远是合法的。
	left := 0
	right := len(nums) - 1
	middle := (left + right) / 2
	// 只有可能的区间只有一个元素的话，那么直接确定是山峰。
	for left < right {
		if nums[middle] > nums[middle+1] {
			// 右边是下坡，山峰在左边，并且“middle”有可能是山峰。
			right = middle
		} else if nums[middle] < nums[middle+1] {
			// 右边是上坡，山峰在右边，并且“middle”不可能是山峰。
			left = middle + 1
		} else {
			// "middle"不可能是山峰，因为峰值要求严格大于左右元素。
			left = middle + 1
		}
		middle = (left + right) / 2
	}
	return middle
}
