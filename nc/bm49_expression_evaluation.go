package main

import (
	"strconv"
	"strings"
)

func solveExpression(s string) int {
	// write code here
	priority := make(map[byte]int)
	priority['+'] = 1
	priority['-'] = 1
	priority['*'] = 2
	priority['/'] = 2
	// "s = strings.ReplaceAll(s, " ", "")"
	// Remove spaces.
	s = strings.Replace(s, " ", "", -1)

	oprators := make([]byte, 0)
	numbers := make([]int, 1)
	// To prevent first number is minus.
	numbers[0] = 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			oprators = append(oprators, '(')
		} else if s[i] >= '0' && s[i] <= '9' {
			begin := i
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
			number, _ := strconv.Atoi(s[begin:i])
			i--
			numbers = append(numbers, number)
		} else if s[i] == ')' {
			for oprators[len(oprators)-1] != '(' {
				oprator := oprators[len(oprators)-1]
				oprators = oprators[:len(oprators)-1]

				right := numbers[len(numbers)-1]
				left := numbers[len(numbers)-2]
				numbers = numbers[:len(numbers)-2]
				temp := evaluate(left, right, oprator)
				// Add the answer into stack.
				numbers = append(numbers, temp)
			}
			// Pop the "(".
			oprators = oprators[:len(oprators)-1]
		} else if priority[s[i]] > 0 {
			cur := s[i]
			for len(oprators) > 0 && priority[cur] <= priority[oprators[len(oprators)-1]] {
				oprator := oprators[len(oprators)-1]
				oprators = oprators[:len(oprators)-1]

				right := numbers[len(numbers)-1]
				left := numbers[len(numbers)-2]
				numbers = numbers[:len(numbers)-2]
				temp := evaluate(left, right, oprator)
				// Add the answer into stack.
				numbers = append(numbers, temp)
			}
			oprators = append(oprators, cur)
		}
	}
	// Complete the rest oprations.
	for len(oprators) != 0 {
		oprator := oprators[len(oprators)-1]
		oprators = oprators[:len(oprators)-1]

		right := numbers[len(numbers)-1]
		left := numbers[len(numbers)-2]
		numbers = numbers[:len(numbers)-2]
		temp := evaluate(left, right, oprator)
		// Add the answer into stack.
		numbers = append(numbers, temp)

	}

	return numbers[len(numbers)-1]
}

func evaluate(left, right int, oprator byte) int {
	if oprator == '+' {
		return left + right
	} else if oprator == '-' {
		return left - right
	} else if oprator == '*' {
		return left * right
	} else {
		return left / right
	}
}
