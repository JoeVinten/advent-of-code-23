package utils

type Match struct {
	Value string
	Index int
}

func AppendWordsAndNumbers(words []Match, numbers [][]int, input string) []Match {
	var allMatches []Match
	for _, word := range words {
		allMatches = append(allMatches, word)
	}
	for _, number := range numbers {
		allMatches = append(allMatches, Match{Value: numberToString(number, input), Index: number[0]})
	}
	return allMatches
}
