package main

func GetUglyNumber_Solution(index int) int {
	// “dp[i]”代表下标从0开始的第“i”个丑数的值。
	dp := make([]int, index+1)
	dp[0] = 0
	if index == 0 {
		return dp[index]
	}
	// 下标为1的丑数是1。
	dp[1] = 1
	// 定义三个指针，分别是以2 / 3 / 5为质因子的当前丑数的下标。
	p1, p2, p3 := 1, 1, 1
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	// 求从2到“index”的丑数。
	for i := 2; i <= index; i++ {
		// 选取从当前丑数生成的下一个丑数中最小的一个。
		next := min(min(dp[p1]*2, dp[p2]*3), dp[p3]*5)
		dp[i] = next
		// 如果生成的丑数被选取，那么对应的指针自增。
		if next == dp[p1]*2 {
			p1++
		}
		if next == dp[p2]*3 {
			p2++
		}
		if next == dp[p3]*5 {
			p3++
		}
	}

	return dp[index]
}
