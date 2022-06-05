package main

func judgePalindromeStr(str string) bool {
	// 特殊情况。
	if len(str) == 0 {
		return true
	}

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		// 只要出现不相同的字符，就不是回文的。
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
