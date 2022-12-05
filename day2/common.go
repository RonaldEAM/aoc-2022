package main

import "strings"

func getTurnChoices(line string) (string, string) {
	cleanLine := strings.TrimSpace(line)
	splitStr := strings.Split(cleanLine, " ")
	return splitStr[0], splitStr[1]
}
