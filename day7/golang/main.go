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

func main() {
	// TESTCASE
	test, err := os.ReadFile("test.txt")
	check(err)
	testArr := s.Split(string(test), "\n")

	// INPUT
	input, err := os.ReadFile("input.txt")
	check(err)
	inputArr := s.Split(string(input), "\n")

	// PROBLEM ONE

	// TESTCASE
	fmt.Printf("test1: Passed:{%v} Answer:{%v}\n", problemOne(testArr) == 6440, problemOne(testArr))

	// // INPUT
	fmt.Printf("answer1: Passed:{%v} Answer:{%v}\n", problemOne(inputArr) == 248559379, problemOne(inputArr))

	// PROBLEM TWO

	// TESTCASE
	fmt.Printf("test2: Passed:{%v} Answer:{%v}\n", problemTwo(testArr) == 5905, problemTwo(testArr))

	// INPUT
	fmt.Printf("answer2: Passed:{%v} Answer:{%v}\n", problemTwo(inputArr) == 249631254, problemTwo(inputArr))
}
