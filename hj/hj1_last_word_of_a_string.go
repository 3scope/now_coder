package main

func findSubString(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else {
			break
		}
	}

	return count
}

// func main() {
// 	input := bufio.NewScanner(os.Stdin)
// 	if input.Scan() {
// 		s := input.Text()
// 		fmt.Printf("%d\n", findSubString(s))
// 	}
// }
