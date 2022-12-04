package main

import (
	"github.com/RonaldEAM/aoc-2022/inputHelper"
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
	"log"
	"strconv"
	"strings"
)

func getTopThreeCalories() []int {
	inverseIntComparator := func(a, b interface{}) int {
		return -utils.IntComparator(a, b)
	}
	maxHeap := binaryheap.NewWith(inverseIntComparator)

	elfSum := 0

	inputHelper.ProcessByLine("./input.txt", func(line string) {
		value := strings.TrimSpace(line)

		if value == "" {
			maxHeap.Push(elfSum)
			elfSum = 0
			return
		}

		calories, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		elfSum += calories
	})

	topThreeCalories := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		v, ok := maxHeap.Pop()
		if !ok {
			break
		}
		topThreeCalories = append(topThreeCalories, v.(int))
	}

	return topThreeCalories
}
