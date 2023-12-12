package main

import (
	"fmt"
	"os"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func filter(arr []string, condition func(s string) bool) []string {
	var filtered []string

	for _, s := range arr {
		if !condition(s) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}
func main() {
	// TESTCASE
	// test, err := os.ReadFile("test.txt")
	// check(err)
	// testArr := s.Split(string(test), "\n")

	// INPUT
	input, err := os.ReadFile("input.txt")
	check(err)
	inputArr := s.Split(string(input), "\n")

	// TESTCASE
	// filteredTestArr := filter(testArr, func(s string) bool {
	// 	return s == "" || s == " "
	// })

	// INPUT
	filteredInputArr := filter(inputArr, func(s string) bool {
		return s == "" || s == " "
	})

	// PROBLEM ONE

	// TESTCASE
	// fmt.Printf("test1: Passed:{%v} Answer:{%v}\n", problemOne(filteredTestArr) == 35, problemOne(filteredTestArr))

	// INPUT
	// fmt.Printf("answer1: Passed:{%v} Answer:{%v}\n", problemOne(filteredInputArr) == 3374647, problemOne(filteredInputArr))

	// PROBLEM TWO

	// TESTCASE
	// fmt.Printf("test2: Passed:{%v} Answer:{%v}\n", problemTwo(filteredTestArr) == 46, problemTwo(filteredTestArr))

	// INPUT
	fmt.Printf("answer2: Passed:{%v} Answer:{%v}\n", problemTwo(filteredInputArr) == 0, problemTwo(filteredInputArr))
}
