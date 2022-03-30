package main

func solveBigNumber(s string, t string) string {
	// write code here
	left := reverseByte(s)
	right := reverseByte(t)
	var result []byte
	leftMax := len(left) > len(right)
	if leftMax {
		result = left
	} else {
		result = right
	}

	var carry uint8
	i := 0
	for ; i < len(left) && i < len(right); i++ {
		// Convert sum to byte but carry not.
		sum := ((left[i]-'0')+(right[i]-'0')+carry)%10 + '0'
		carry = ((left[i] - '0') + (right[i] - '0') + carry) / 10
		result[i] = sum
	}

	if leftMax {
		for ; i < len(left); i++ {
			sum := (left[i]-'0'+carry)%10 + '0'
			carry = (left[i] - '0' + carry) / 10
			result[i] = sum
		}
	} else {
		for ; i < len(right); i++ {
			sum := (right[i]-'0'+carry)%10 + '0'
			carry = (right[i] - '0' + carry) / 10
			result[i] = sum
		}
	}
	if carry != 0 {
		result = append(result, byte(carry+'0'))
	}

	return string(result)
}

func reverseByte(s string) []byte {
	byteSlice := []byte(s)
	for i, j := 0, len(byteSlice)-1; i < j; i, j = i+1, j-1 {
		byteSlice[i], byteSlice[j] = byteSlice[j], byteSlice[i]
	}
	return byteSlice
}
