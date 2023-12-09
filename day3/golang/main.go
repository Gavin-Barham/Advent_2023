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

func main() {
	input := readFileToStrArray("input.txt", "\n")
	test := readFileToStrArray("test.txt", "\n")

	test1 := problemOne(test)
	answer1 := problemOne(input)
	test2 := problemTwo(test)
	answer2 := problemTwo(input)

	fmt.Println()
	fmt.Printf("Test 1: passed:{%v}, value:{%v} \n", test1 == 4361, test1)
	fmt.Printf("Answer 1: passed:{%v}, value:{%v} \n", answer1 == 525119, answer1)
	fmt.Println()
	fmt.Printf("Test 2: passed:{%v}, value:{%v} \n", test2 == 467835, test2)
	fmt.Printf("Answer 2: passed:{%v}, value:{%v} \n", answer2 == 76504829, answer2)
	fmt.Println()
}
