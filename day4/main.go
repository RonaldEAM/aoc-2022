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
		count := getFullyContainRangesCount()
		fmt.Printf("Assignment pairs where one range fully contain the other: %d\n", count)
	case 2:
		count := getOverlapsCount()
		fmt.Printf("Number of overlapping assignment pairs: %d\n", count)
	default:
		fmt.Println("This step is not valid")
		os.Exit(1)
	}
}
