package main

func replaceSpace(s string) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		}
	}
	if count == 0 {
		return s
	}
	result := make([]byte, len(s)+2*count)
	for i, j := len(result)-1, len(s)-1; i >= 0; {
		if s[j] == ' ' {
			result[i] = '0'
			result[i-1] = '2'
			result[i-2] = '%'
			i = i - 3
			j--
		} else {
			result[i] = s[j]
			i--
			j--
		}
	}

	return string(result)
}
