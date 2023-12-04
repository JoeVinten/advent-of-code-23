package utils

import (
	"regexp"
	"strings"
)

func CalculateMatches(line string) int {
	cardNumberRegex := regexp.MustCompile(`Card\s*\d+:`)

	removeGame := cardNumberRegex.ReplaceAllString(line, "")

	splitString := strings.Split(removeGame, "|")

	winningNumbers := strings.Fields(splitString[0])
	scratchCardNumbers := strings.Fields(splitString[1])

	numberOfMatches := 0

	for i := 0; i < len(winningNumbers); i++ {
		if ArrayContains(scratchCardNumbers, winningNumbers[i]) {
			numberOfMatches += 1
		}
	}

	return numberOfMatches
}
