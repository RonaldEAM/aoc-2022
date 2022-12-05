package main

import (
	"github.com/RonaldEAM/aoc-2022/inputhelper"
)

func calculateScore2() int {
	// A: Rock, B: Paper, C: Scissors
	// X: Need to lose, Y: draw, Z: win

	winConditions := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}
	choicePoints := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	roundResultPoints := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	totalScore := 0

	inputhelper.ProcessByLine("./input.txt", func(line string) {
		opponentHand, neededResult := getTurnChoices(line)

		currRoundScore := roundResultPoints[neededResult]
		switch neededResult {
		case "X":
			currRoundScore += choicePoints[winConditions[opponentHand]]
		case "Y":
			currRoundScore += choicePoints[opponentHand]
		case "Z":
			currRoundScore += choicePoints[getWinnerOf(opponentHand)]
		}

		totalScore += currRoundScore
	})

	return totalScore
}

type node struct {
	val     string
	winOver *node
}

func getWinnerOf(hand string) string {
	var rock node
	var paper node
	var scissor node

	rock.val = "A"
	rock.winOver = &scissor

	paper.val = "B"
	paper.winOver = &rock

	scissor.val = "C"
	scissor.winOver = &paper

	hands := map[string]node{
		"A": rock,
		"B": paper,
		"C": scissor,
	}

	// On the premise that:
	// rock ---wins over--> scissor ---wins over--> paper ---wins over--> rock
	// If we want the winner hand of another hand, we need to follow the links two times from the given hand
	return hands[hand].winOver.winOver.val
}
