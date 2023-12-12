package main

import (
	"strconv"
	s "strings"
)

func problemTwo(input []string) int {
	reverseMap := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}
	seedsArr := s.Split(input[0][7:len(input[0])], " ")
	input = input[1:]
	plotMap := createPlotMap(input)
	out := 1000000000000
	for seed := 0; seed < len(seedsArr); seed += 2 {
		startSeed, err := strconv.Atoi(seedsArr[seed])
		check(err)
		rangeLength, err := strconv.Atoi(seedsArr[seed+1])
		check(err)
		seedRange := startSeed + rangeLength
		nextDestination := startSeed
		for r := startSeed; r < seedRange; r++ {
			nextDestination = r
			for i := range reverseMap {
				key := reverseMap[i]
				matrix := plotMap[key]
				for _, row := range matrix {
					offset := nextDestination
					destination := row[0]
					sourceStart := row[1]
					sourceLength := sourceStart + row[2] - 1

					if nextDestination >= sourceStart && nextDestination <= sourceLength {
						offset = nextDestination - sourceStart
						nextDestination = destination + offset
						break
					}
				}
			}
			out = min(nextDestination, out)
		}
	}
	return out
}
