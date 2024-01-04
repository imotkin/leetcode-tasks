package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)

	for _, letter := range magazine {
		if _, ok := letters[letter]; !ok {
			letters[letter] = 1
		} else {
			letters[letter]++
		}
	}

	for _, letter := range ransomNote {
		if _, ok := letters[letter]; !ok {
			return false
		} else if letters[letter] == 0 {
			return false
		} else {
			letters[letter]--
		}
	}

	return true
}
