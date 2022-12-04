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
		totalScore := calculateScore()
		fmt.Printf("Total score is %d\n", totalScore)
	case 2:
		totalScore := calculateScore2()
		fmt.Printf("Total score is %d\n", totalScore)
	default:
		fmt.Println("This step is not valid")
		os.Exit(1)
	}
}
