package main

import (
	"github.com/RonaldEAM/aoc-2022/inputhelper"
)

func getOverlapsCount() int {
	count := 0

	inputhelper.ProcessByLine("./input.txt", func(line string) {
		firstRange, secondRange := parseRanges(line)
		if firstRange.start <= secondRange.end && secondRange.start <= firstRange.end {
			count++
		}
	})

	return count
}
