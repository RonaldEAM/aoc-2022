package main

import (
	"log"
	"strconv"
	"strings"
)

func getMostCalories() int {
	elfSum := 0
	maxCalories := 0

	processInputByLine("./input.txt", func(line string) {
		value := strings.TrimSpace(line)

		if value == "" {
			if elfSum > maxCalories {
				maxCalories = elfSum
			}
			elfSum = 0
			return
		}

		calories, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		elfSum += calories
	})

	return maxCalories
}
