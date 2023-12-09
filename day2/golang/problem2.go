package main

import (
	sc "strconv"
	s "strings"
)

func problemTwo(input []string) int {
	total := 0

	for _, str := range input {
		gameResults := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		splitStr := s.Split(str, ":")[1]
		game := s.Split(splitStr, ";")
		for _, val := range game {
			val = s.Trim(val, " ")
			round := s.Split(s.Join(s.Split(val, ","), ""), " ")
			for i := 0; i < len(round)-1; i += 2 {
				count, err := sc.Atoi(round[i])
				check(err)
				color := round[i+1]
				if gameResults[color] < count {
					gameResults[color] = count
				}
			}
		}
		total += gameResults["red"] * gameResults["green"] * gameResults["blue"]
	}
	return total
}
