package main

func reverseString(s []byte) []byte {
	var start, end = 0, len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]

		start++
		end--
	}

	return s
}
