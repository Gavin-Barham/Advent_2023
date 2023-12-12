package main

import (
	sc "strconv"
	s "strings"
)

func formatSingleRaceInput(input []string) map[string]int {
	out := make(map[string]int)
	for i, arr := range input {
		times := filter(s.Split(s.Split(arr, ":")[1], " "), func(s string) bool { return s == "" || s == " " })
		var numStr string
		for _, v := range times {
			numStr += v
		}
		num, err := sc.Atoi(numStr)
		check(err)
		if i == 0 {
			out["time"] = num
		} else {
			out["distance"] = num
		}
	}
	return out
}
func problemTwo(input []string) int {
	out := 0
	myMap := formatSingleRaceInput(input)
	totalTime := myMap["time"]
	record := myMap["distance"]
	marginOfErr := 0
	moveSpeed := 0
	for currTimer := 0; currTimer < totalTime; currTimer++ {
		remainingTime := totalTime - currTimer
		moveSpeed = currTimer
		distanceTraveled := remainingTime * moveSpeed
		if distanceTraveled > record {
			marginOfErr++
		}
	}
	if out == 0 {
		out = marginOfErr
	} else {
		out *= marginOfErr
	}

	return out
}
