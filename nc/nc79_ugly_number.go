package main

func GetUglyNumber_Solution(index int) int {
	// write code here
	if index == 0 {
		return 0
	}

	uglyNumber := make([]int, 1)
	uglyNumber[0] = 1
	p1, p2, p3 := 0, 0, 0
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	// Every element
	for i := 1; i < index; i++ {
		next := min(min(uglyNumber[p1]*2, uglyNumber[p2]*3), uglyNumber[p3]*5)
		uglyNumber = append(uglyNumber, next)
		if next == uglyNumber[p1]*2 {
			p1++
		}
		// Next may be same twice.
		if next == uglyNumber[p2]*3 {
			p2++
		}
		if next == uglyNumber[p3]*5 {
			p3++
		}
	}

	return uglyNumber[len(uglyNumber)-1]
}
