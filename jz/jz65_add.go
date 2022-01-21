package main

func Add(num1 int, num2 int) int {
	// write code here
	sum := num1
	flag := num2
	for flag != 0 {
		sum = num1 ^ num2
		flag = num1 & num2 << 1
		num1 = sum
		num2 = flag
	}
	return sum
}
