package main

import (
	"github.com/RonaldEAM/aoc-2022/inputhelper"
	"strings"
)

func getPrioritiesSum() int {
	totalSum := 0
	getCharValue := buildCharValues()

	inputhelper.ProcessByLine("./input.txt", func(line string) {
		value := strings.TrimSpace(line)
		half := len(value) / 2
		currCharMap := map[rune]int{}
		var sharedItemType rune
		for i, char := range value {
			if i >= half {
				if _, ok := currCharMap[char]; ok {
					sharedItemType = char
					break
				}
				continue
			}

			if _, ok := currCharMap[char]; !ok {
				currCharMap[char] = 0
			}
			currCharMap[char] += 1
		}

		totalSum += getCharValue(sharedItemType)
	})

	return totalSum
}
