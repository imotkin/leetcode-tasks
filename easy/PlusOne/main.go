package main

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	} else {
		digits[len(digits)-1] = 0
		nextDigit := digits[:len(digits)-1]
		if len(nextDigit) != 0 {
			plusOne(nextDigit)
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
