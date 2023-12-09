package main

import (
	sc "strconv"
	s "strings"
)

func problemOne(input []string) int {
	var expectedResults = map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	total := 0

	for _, str := range input {
		gameResults := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		game := s.Split(str, ":")
		gameNum, err := sc.Atoi(s.Split(string(game[0]), " ")[1])
		check(err)
		rounds := s.Split(game[1], ";")
		for _, val := range rounds {
			val = s.Trim(val, " ")
			round := s.Split(val, ",")
			round = s.Split(s.Join(round, ""), " ")
			for i := 0; i < len(round)-1; i += 2 {
				count, err := sc.Atoi(round[i])
				check(err)
				color := round[i+1]

				if gameResults[color] < count {
					gameResults[color] = count
				}
			}

		}
		if expectedResults["red"] >= gameResults["red"] && expectedResults["green"] >= gameResults["green"] && expectedResults["blue"] >= gameResults["blue"] {
			total += gameNum
		}
	}
	return total
}
