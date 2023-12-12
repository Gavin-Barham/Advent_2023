package main

// import (
// 	sc "strconv"
// 	s "strings"
// )

// func formatInput(input []string) map[string][]int {
// 	out := make(map[string][]int)
// 	for i, arr := range input {
// 		times := filter(s.Split(s.Split(arr, ":")[1], " "), func(s string) bool { return s == "" || s == " " })
// 		var nums []int
// 		for _, v := range times {
// 			num, err := sc.Atoi(v)
// 			check(err)
// 			nums = append(nums, num)
// 		}
// 		if i == 0 {
// 			out["times"] = nums
// 		} else {
// 			out["distances"] = nums
// 		}
// 	}
// 	return out
// }
// func problemOne(input []string) int {
// 	out := 0
// 	myMap := formatInput(input)
// 	for i, totalTime := range myMap["times"] {
// 		marginOfErr := 0
// 		record := myMap["distances"][i]
// 		moveSpeed := 0
// 		for currTimer := 0; currTimer < totalTime; currTimer++ {
// 			remainingTime := totalTime - currTimer
// 			moveSpeed = currTimer
// 			distanceTraveled := remainingTime * moveSpeed
// 			if distanceTraveled > record {
// 				marginOfErr++
// 			}
// 		}
// 		if out == 0 {
// 			out = marginOfErr
// 		} else {
// 			out *= marginOfErr
// 		}
// 	}
// 	return out
// }
