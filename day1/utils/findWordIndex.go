package utils

import "regexp"

func FindWordIndex(words []string, input string) []Match {
	var wordsIndex []Match
	for _, word := range words {
		writtenNumberRegex := regexp.MustCompile(word)
		indexes := writtenNumberRegex.FindAllStringIndex(input, -1)
		for _, indexPair := range indexes {
			wordsIndex = append(wordsIndex, Match{Value: word, Index: indexPair[0]})
		}
	}
	return wordsIndex
}
