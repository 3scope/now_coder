package main

func solveBigNumber(s string, t string) string {
	// 翻转字符串，从最末尾加起。
	left := reverseByte(s)
	right := reverseByte(t)
	var result []byte
	// “result”存放最终结果。
	leftMax := len(left) > len(right)
	if leftMax {
		// 如果“left”比较长，那么初始化长度按照“left”来设置。
		result = make([]byte, len(left))
	} else {
		result = make([]byte, len(right))
	}

	var carry uint8
	i := 0
	for ; i < len(left) && i < len(right); i++ {
		sum := ((left[i] - '0') + (right[i] - '0') + carry) % 10
		// “carry”是偏移量，不需要加上“0”。
		carry = ((left[i] - '0') + (right[i] - '0') + carry) / 10
		// “result”需要以字符串的形式来存放结果，所以“sum”需要加上“0”，来让最后的结果显示数字。
		result[i] = sum + '0'
	}

	if leftMax {
		for ; i < len(left); i++ {
			sum := (left[i] - '0' + carry) % 10
			carry = (left[i] - '0' + carry) / 10
			result[i] = sum + '0'
		}
	} else {
		for ; i < len(right); i++ {
			sum := (right[i] - '0' + carry) % 10
			carry = (right[i] - '0' + carry) / 10
			result[i] = sum + '0'
		}
	}
	// 最后进位不为0，需要存放到“result”中
	if carry != 0 {
		result = append(result, carry+'0')
	}

	// 最后需要翻转一次。
	result = reverseByte(string(result))
	return string(result)
}

// 翻转字符串。
func reverseByte(s string) []byte {
	byteSlice := []byte(s)
	for i, j := 0, len(byteSlice)-1; i < j; i, j = i+1, j-1 {
		byteSlice[i], byteSlice[j] = byteSlice[j], byteSlice[i]
	}
	return byteSlice
}
