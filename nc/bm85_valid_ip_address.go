package main

import (
	"strconv"
	"strings"
)

func solve(IP string) string {
	if validIPV4(IP) {
		return "IPv4"
	}
	if validIPV6(IP) {
		return "IPv6"
	}
	return "Neither"
}

// 验证 IPV4 地址。
func validIPV4(ip string) bool {
	part := strings.Split(ip, ".")
	// 不是四组数值不可能是 IPV4 地址。
	if len(part) != 4 {
		return false
	}

	for i := 0; i < 4; i++ {
		// 不能有前缀0。
		if len(part[i]) > 1 && part[i][0] == '0' {
			return false
		}
		number, err := strconv.Atoi(part[i])
		if err != nil {
			return false
		}
		// 如果数值大于255或者小于0，都不会是正确地址。
		if number > 255 || number < 0 {
			return false
		}
	}
	return true
}

func validIPV6(ip string) bool {
	part := strings.Split(ip, ":")
	// 不是8组不可能是 IPV6 地址。
	if len(part) != 8 {
		return false
	}

	for i := 0; i < 8; i++ {
		// 字符串的长度一定是在[1, 4]。
		if len(part[i]) > 4 || len(part[i]) == 0 {
			return false
		}
		_, err := strconv.ParseInt(part[i], 16, 64)
		// 如果转换16进制出错，证明有其他字符，不会是 IPV6 地址。
		if err != nil {
			return false
		}
	}
	return true
}
