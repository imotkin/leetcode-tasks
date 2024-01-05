package main

func addDigits(num int) int {
	if num < 10 {
		return num
	} else {
		var digits []int

		for num > 0 {
			digits = append(digits, (num % 10))
			num /= 10
		}

		for _, digit := range digits {
			num += digit
		}

		return addDigits(num)
	}
}
