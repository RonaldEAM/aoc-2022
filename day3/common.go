package main

func buildCharValues() func(rune) int {
	i := 1
	charValueMap := make(map[rune]int, 52)
	for ch := 'a'; ch <= 'z'; ch++ {
		charValueMap[ch] = i
		i++
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		charValueMap[ch] = i
		i++
	}

	return func(char rune) int {
		return charValueMap[char]
	}
}
