package main

func solveString(str string) string {
	byteSlice := []byte(str)
	for i, j := 0, len(byteSlice)-1; i < j; i, j = i+1, j-1 {
		byteSlice[i], byteSlice[j] = byteSlice[j], byteSlice[i]
	}

	return string(byteSlice)
}
