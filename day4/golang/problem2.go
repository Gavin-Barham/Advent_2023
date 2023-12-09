package main

import (
	"regexp"
	s "strings"
)

func problemTwo(input []string) int {
	sum := 0
	freq := make(map[int]int)
	for i, str := range input {
		card := s.Split(s.Split(str, ":")[1], "|")
		re := regexp.MustCompile(`\s+`)
		results := re.Split(card[1], -1)
		winningNums := re.Split(card[0], -1)

		filteredResult := filter(results, func(x string) bool {
			return x == ""
		})
		gameNum := i + 1
		if freq[gameNum] == 0 {
			freq[gameNum] = 1
		} else {
			freq[gameNum]++
		}
		gameNum++
		for _, numStr := range filteredResult {
			if contains(winningNums, numStr) {
				if freq[gameNum] >= 1 {
					count := freq[i+1]
					for k := 0; k < count; k++ {
						freq[gameNum]++
					}
					gameNum++
				} else {
					freq[gameNum]++
					gameNum++

				}
			}
		}
	}
	for _, v := range freq {
		sum += v
	}
	return sum
}
