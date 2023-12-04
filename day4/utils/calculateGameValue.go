package utils

import "fmt"

func CalculateGameValue(matches int) int {
	fmt.Println("number of matches: ", matches)
	if matches == 0 {
		return 0
	}

	pointsTotal := 1
	for i := 1; i < matches; i++ {
		pointsTotal *= 2
	}

	return pointsTotal
}
