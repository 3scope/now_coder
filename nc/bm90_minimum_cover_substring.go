package main

func minWindow(S string, T string) string {
	// T 是模式串，先得到 T 中所有字符的数量。
	tByteCount := make(map[byte]int)
	for i := 0; i < len(T); i++ {
		tByteCount[T[i]]++
	}

	// 存放已经满足要求的字符的数量。
	isOK := 0
	// 使用滑动窗口，左闭右开。
	left, right := 0, 0
	result := ""
	for right <= len(S) && left < len(S) {
		// 如果当前窗口里面的字符串不能满足要求，那么就向右滑动。
		if isOK < len(tByteCount) {
			// 窗口已经不能再次扩大了，跳出循环。
			if right == len(S) {
				break
			}
			// 如果字符在 T 中出现过，那么对应的字符数量减1。
			if v, ok := tByteCount[S[right]]; ok {
				tByteCount[S[right]]--
				// 如果刚好减为0，那么该字符已满足要求。
				if v-1 == 0 {
					isOK++
				}
			}
			right++
		} else {
			if result == "" {
				result = S[left:right]
			}
			// “result”不断取最小值。
			if len(result) > right-left {
				result = S[left:right]
			}
			// 已经满足要求，尝试删除多余的字符。
			// 如果字符在 T 中出现过，那么对应的字符数量加1。
			if v, ok := tByteCount[S[left]]; ok {
				tByteCount[S[left]]++
				// 如果当前窗口中的该字符的数量和原字符串中对应字符数量相同，那么该字符串不满足要求。
				if v == 0 {
					isOK--
				}
			}
			left++
		}
	}

	return result
}
