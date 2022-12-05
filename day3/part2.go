package main

import (
	"github.com/RonaldEAM/aoc-2022/inputHelper"
	"strings"
)

func getGroupPrioritiesSum() int {
	totalSum := 0
	getCharValue := buildCharValues()
	currCharMap := map[rune]int{}
	i := 1

	inputHelper.ProcessByLine("./input.txt", func(line string) {
		var groupSharedItem rune
		var seenInLine [52]bool

		value := strings.TrimSpace(line)

		for _, char := range value {
			pos := getCharValue(char) - 1
			if _, ok := currCharMap[char]; !ok {
				currCharMap[char] = 1
			} else if !seenInLine[pos] {
				currCharMap[char] += 1
			}
			seenInLine[pos] = true

			if currCharMap[char] == 3 {
				groupSharedItem = char
				break
			}
		}

		if i%3 == 0 {
			totalSum += getCharValue(groupSharedItem)
			currCharMap = map[rune]int{}
		}

		i += 1
	})

	return totalSum
}
