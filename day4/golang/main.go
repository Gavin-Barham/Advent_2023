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
func readFileToStrArray(filePath string, separator string) []string {
	file, err := os.ReadFile(filePath)
	check(err)
	return s.Split(string(file), separator)
}
func contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}
func filter(numbers []string, condition func(string) bool) []string {
	var filtered []string

	for _, num := range numbers {
		if !condition(num) {
			filtered = append(filtered, num)
		}
	}

	return filtered
}
func main() {
	input := readFileToStrArray("input.txt", "\n")
	test := readFileToStrArray("test.txt", "\n")

	test1 := problemOne(test)
	answer1 := problemOne(input)
	test2 := problemTwo(test)
	answer2 := problemTwo(input)

	fmt.Println()

	// PROBLEM 1:
	fmt.Printf("Test 1: passed:{%v}, value:{%v} \n", test1 == 13, test1)
	fmt.Printf("Answer 1: passed:{%v}, value:{%v} \n", answer1 == 33950, answer1)

	fmt.Println()

	// PROBLEM 2:
	fmt.Printf("Test 2: passed:{%v}, value:{%v} \n", test2 == 30, test2)
	fmt.Printf("Answer 2: passed:{%v}, value:{%v} \n", answer2 == 14814534, answer2)

	fmt.Println()
}
