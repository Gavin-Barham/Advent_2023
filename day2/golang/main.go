package main

import (
	"fmt"
	"os"
	s "strings"
)

func print(val interface{}) {
	fmt.Printf("%v \n", val)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func inputFileToArray(filePath string, separator string) []string {
	file, err := os.ReadFile(filePath)
	check(err)
	arr := s.Split(string(file), separator)
	return arr
}

func main() {
	input := inputFileToArray("input.txt", "\n")

	fmt.Print("Answer 1: ")
	print(problemOne(input))
	fmt.Print("Answer 2: ")
	print(problemTwo(input))
}
