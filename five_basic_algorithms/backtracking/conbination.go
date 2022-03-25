package backtracking

type Result struct {
	Temp   []int
	Answer [][]int
}

func (r *Result) Conbination(size, k, start int) {
	if len(r.Temp) == k {
		r.Answer = append(r.Answer, r.Temp)
		return
	}
	// Pruning.
	// 需要确定i最多到多少，才能满足有k个，因此，将循环条件i < size更改为i < size - (k - len(r.Temp)) + 1
	// 															|<-- k-len() -->|
	//													i最多取值取到这里		size
	// 剪枝可以有两种方法，一种是更改循环条件，另一个种是多一个"if"判断。
	// if len(r.Temp)+(size-start+1) < k {
	// 	return
	// }
	for i := start; i <= size; i++ {
		// Deal with subset in this level.
		r.Temp = append(r.Temp, i)
		// Recursive function.
		r.Conbination(size, k, i+1)
		// Backtracking code.
		r.Temp = r.Temp[:len(r.Temp)-1]
	}
}

func (r *Result) Plus(candidate []int, target int, start int) {
	if r.Temp[0] == target {
		r.Answer = append(r.Answer, r.Temp[1:len(r.Temp)])
		return
	} else if r.Temp[0] > target {
		return
	}

	for i := start; i < len(candidate); i++ {
		r.Temp[0] += candidate[i]
		r.Temp = append(r.Temp, candidate[i])
		r.Plus(candidate, target, i)
		r.Temp[0] -= candidate[i]
		r.Temp = r.Temp[:len(r.Temp)-1]
	}
}

func (r *Result) Palindrome(s string, start int) {
	if start >= len(s) {
		r.Answer = append(r.Answer, r.Temp)
		return
	}
	for i := start; i <= len(s); i++ {
		if isPalindrome(s[start:i]) {
			// Record slice position.
			r.Temp = append(r.Temp, i)
			r.Palindrome(s, i+1)
			r.Temp = r.Temp[:len(r.Temp)-1]
		} else {
			continue
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s); i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
