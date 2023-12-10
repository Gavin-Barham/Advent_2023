package main

import (
	"regexp"
	s "strings"
)

func problemOne(input []string) int {
	sum := 0
	for _, str := range input {
		cardSum := 0
		card := s.Split(s.Split(str, ":")[1], "|")
		re := regexp.MustCompile(`\s+`)
		results := re.Split(card[1], -1)
		winningNums := re.Split(card[0], -1)

		filteredResult := filter(results, func(x string) bool {
			return x == ""
		})
		for _, numStr := range filteredResult {
			if contains(winningNums, numStr) {
				if cardSum == 0 {
					cardSum++
				} else {
					cardSum *= 2
				}
			}

		}
		sum += cardSum
	}
	return sum
}
