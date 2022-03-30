package main

import (
	"strconv"
	"strings"
)

func solve(IP string) string {
	// write code here
	if validIPV4(IP) {
		return "IPv4"
	}
	if validIPV6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPV4(ip string) bool {
	part := strings.Split(ip, ":")
	if len(part) != 4 {
		return false
	}
	for i := 0; i < 4; i++ {
		if len(part[i]) > 1 && part[i][0] == '0' {
			return false
		}
		number, err := strconv.Atoi(part[i])
		if err != nil {
			return false
		}
		if number > 255 || number < 0 {
			return false
		}
	}

	return true
}

func validIPV6(ip string) bool {
	part := strings.Split(ip, ":")
	if len(part) != 8 {
		return false
	}

	for i := 0; i < 8; i++ {
		if len(part[i]) > 4 || len(part[i]) == 0 {
			return false
		}
		_, err := strconv.ParseInt(part[i], 16, 64)
		if err != nil {
			return false
		}
	}
	return true
}
