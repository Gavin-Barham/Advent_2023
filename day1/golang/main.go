package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func isNumeric(char byte) bool {
	return char <= 57 && char >= 48
}
func print(val interface{}) {
	fmt.Printf("%v \n", val)
}

func main() {
	answerOne := problemOne()
	answerTwo := problemTwo()
	print(answerOne)
	print(answerTwo)
}
