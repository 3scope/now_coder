package dynamic_programming

func Backpack(weight, value []int, size int) int {
	number := len(weight)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, number)
	}

	// Init.
	for j := 0; j < size; j++ {
		if j < weight[0] {
			dp[0][j] = 0
		} else {
			dp[0][j] = weight[0]
		}
	}

	for i := 1; i < number; i++ {
		for j := 0; j < size; j++ {
			if j-weight[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+weight[i])
			}
		}
	}
	return dp[number-1][size-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func RollingArray(weight, value []int, size int) int {
	number := len(weight)
	dp := make([]int, size)

	for j := 0; j < size; j++ {
		if j < weight[0] {
			dp[j] = 0
		} else {
			dp[j] = weight[0]
		}
	}

	for i := 1; i < number; i++ {
		for j := size - 1; j >= 0; j-- {
			if j-weight[i] >= 0 {
				dp[j] = max(dp[j], dp[j-weight[i]]+weight[i])
			}
		}
	}

	return dp[size-1]
}
