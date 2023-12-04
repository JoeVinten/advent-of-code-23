package main

import (
	"adventofcode/day4/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	cardNum string
	amount  int
}

func split(c rune) bool {
	return c == ':' || c == '|'
}

func winningCards(input []string) int {
	cards := []card{}

	for i := 1; i <= len(input); i++ {
		cards = append(cards, card{cardNum: strconv.Itoa(i), amount: 1})
	}

	for i, line := range input {
		fields := strings.FieldsFunc(line, split)
		winning := strings.Fields(fields[1])
		numbers := strings.Fields(fields[2])

		count := 0

		for _, num := range numbers {
			if slices.Contains(winning, num) {
				count++
				cards[i+count].amount = cards[i+count].amount + 1*cards[i].amount
			}
		}
	}

	totalCards := 0
	for _, card := range cards {
		totalCards += card.amount
	}

	return totalCards
}

func main() {
	lines, err := utils.ReadLine("./data/smallExample.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	scratchcardPileTotal := 0
	totalCards := 0

	for _, line := range lines {
		totalCards++

		fmt.Printf("🎄🎅🎄🎅🎄🎅🎄🎅🎄🎅 -- GAME %d -- 🎅🎄🎅🎄🎅🎄🎅🎄🎅🎄\n", totalCards)

		numberOfMatches := utils.CalculateMatches(line)

		scratchcardPileTotal += utils.CalculateGameValue(numberOfMatches)
	}

	fmt.Println("****************************************")
	fmt.Println("Part 1: ")
	fmt.Println("Scratchcard Points Total:", scratchcardPileTotal)
	fmt.Println("****************************************")

	// Part 2
	totalCards = winningCards(lines)
	fmt.Println("Part 2: ")
	fmt.Println("Total Cards", totalCards)
}
