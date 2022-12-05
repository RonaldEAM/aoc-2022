package main

import (
	"github.com/RonaldEAM/aoc-2022/inputhelper"
)

func calculateScore() int {
	// X: Rock, Y: Paper, Z: Scissors

	winConditions := map[string]string{
		"X": "Z",
		"Z": "Y",
		"Y": "X",
	}
	opponentChoiceMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	choicePoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	totalScore := 0

	inputhelper.ProcessByLine("./input.txt", func(line string) {
		originalOpponentHand, myHand := getTurnChoices(line)

		// Normalize opponent choice
		opponentHand := opponentChoiceMap[originalOpponentHand]

		currRoundScore := choicePoints[myHand]
		if winConditions[myHand] == opponentHand {
			currRoundScore += 6
		} else if myHand == opponentHand {
			currRoundScore += 3
		}

		totalScore += currRoundScore
	})

	return totalScore
}
