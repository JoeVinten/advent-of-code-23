package main

import (
	"adventofcode/day1/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("data/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		numberRegex := regexp.MustCompile(`[1-9]`)
		input := scanner.Text()

		numbersIndex := numberRegex.FindAllStringIndex(input, -1)

		wordsIndex := utils.FindWordIndex(words, input)

		// combines numbers and digits together
		allMatches := utils.AppendWordsAndNumbers(wordsIndex, numbersIndex, input)
		sort.Slice(allMatches, func(i, j int) bool {
			return allMatches[i].Index < allMatches[j].Index
		})

		if len(allMatches) > 0 {
			firstElement := utils.TextToNumber(allMatches[0].Value)
			lastElement := utils.TextToNumber(allMatches[len(allMatches)-1].Value)

			combineNumbers := firstElement + lastElement

			n, err := strconv.Atoi(combineNumbers)
			if err != nil {
				panic(err)
			}

			total = total + n
		}

	}

	fmt.Println("🎄Code Total 🎄 ", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
