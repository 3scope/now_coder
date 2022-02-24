package main

func LeftRotateString(str string, n int) string {
	if str == "" {
		return ""
	}
	result := []byte(str)
	n = n % len(result)
	if n == 0 {
		return str
	}
	Reverse(result, 0, len(result)-1)
	Reverse(result, 0, len(result)-n-1)
	Reverse(result, len(result)-n, len(result)-1)
	return string(result)
}

func Reverse(str []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
