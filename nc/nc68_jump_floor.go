package main

func JumpFloor(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else if number == 2 {
		return 2
	} else {
		return JumpFloor(number-1) + JumpFloor(number-2)
	}
}

func DynamicProgramming(number int) int {
	// The start state.
	dp1, dp2 := 1, 2
	if number <= 2 {
		return number
	} else {
		for i := 3; i <= number; i++ {
			// State Compression.
			temp := dp1 + dp2
			dp1 = dp2
			dp2 = temp
		}
	}
	return dp2
}
