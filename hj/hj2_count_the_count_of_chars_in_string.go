package main

func countChar(s string) map[byte]int {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	return count
}
