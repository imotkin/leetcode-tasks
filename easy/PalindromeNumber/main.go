package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}
