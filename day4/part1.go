package main

import (
	"github.com/RonaldEAM/aoc-2022/inputhelper"
)

func getFullyContainRangesCount() int {
	count := 0

	inputhelper.ProcessByLine("./input.txt", func(line string) {
		firstRange, secondRange := parseRanges(line)
		if (firstRange.start <= secondRange.start && firstRange.end >= secondRange.end) || (firstRange.start >= secondRange.start && firstRange.end <= secondRange.end) {
			count++
		}
	})

	return count
}
