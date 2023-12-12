package main

import (
	"strconv"
	s "strings"
	"unicode"
)

func createPlotMap(input []string) map[string][][]int {
	plotMap := make(map[string][][]int)
	mapName := ""
	var matrix [][]int
	for i, str := range input {
		var nums []int
		if !unicode.IsNumber(rune(str[0])) {
			if len(matrix) != 0 {
				plotMap[mapName] = matrix
				var newMatrix [][]int
				matrix = newMatrix
			}
			mapName = s.Split(str, " ")[0]
		} else {
			numArr := s.Split(str, " ")
			for _, num := range numArr {
				x, err := strconv.Atoi(num)
				check(err)
				nums = append(nums, x)
			}
			matrix = append(matrix, nums)
			if i >= len(input)-1 {
				plotMap[mapName] = matrix
				var newMatrix [][]int
				matrix = newMatrix
			}
		}
	}
	return plotMap
}

// func problemOne(input []string) int {
// 	reverseMap := []string{
// 		"seed-to-soil",
// 		"soil-to-fertilizer",
// 		"fertilizer-to-water",
// 		"water-to-light",
// 		"light-to-temperature",
// 		"temperature-to-humidity",
// 		"humidity-to-location",
// 	}
// 	seedsArr := s.Split(input[0][7:len(input[0])], " ")
// 	input = input[1:]
// 	plotMap := createPlotMap(input)
// 	out := 1000000000000
// 	for _, seed := range seedsArr {
// 		nextDestination, err := strconv.Atoi(seed)
// 		check(err)

// 		for i := range reverseMap {
// 			key := reverseMap[i]
// 			matrix := plotMap[key]
// 			for _, row := range matrix {
// 				offset := nextDestination
// 				destination := row[0]
// 				sourceStart := row[1]
// 				sourceLength := sourceStart + row[2] - 1

// 				if nextDestination >= sourceStart && nextDestination <= sourceLength {
// 					offset = nextDestination - sourceStart
// 					nextDestination = destination + offset
// 					break
// 				}
// 			}
// 		}
// 		out = min(nextDestination, out)
// 	}
// 	return out
// }
