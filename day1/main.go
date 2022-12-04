package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("You must provide an step argument")
		os.Exit(1)
	}

	step, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Step argument must be a valid integer")
		os.Exit(1)
	}

	switch step {
	case 1:
		mostCalories := getMostCalories()
		fmt.Printf("The elf carrying the most calories is carrying %d calories\n", mostCalories)
	case 2:
		topThreeCalories := getTopThreeCalories()
		total := 0
		for _, calories := range topThreeCalories {
			total += calories
		}
		fmt.Printf("The top three elves carrying the most calories are carrying %v\nTotal: %d\n", topThreeCalories, total)
		os.Exit(1)
	default:
		fmt.Println("This step is not valid")
		os.Exit(1)
	}
}
