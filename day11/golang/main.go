package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "input.txt"
	test := "test.txt"

	testAnswer1 := problemOne(test)
	answer1 := problemOne(input)
	testAnswer2 := problemTwo(test)
	answer2 := problemTwo(input)

	fmt.Printf("testAnswer1:{%v}\nanswer1:{%v}\n\n", testAnswer1, answer1)
	fmt.Printf("testAnswer2:{%v}\nanswer2:{%v}\n\n", testAnswer2, answer2)
}
func getInputFromFile(file string) string {
	input, err := os.ReadFile(file)
	check(err)
	return string(input)
}
func formatInputTo2DArray(input string) [][]string {
	strArr := strings.Split(input, "\n")
	var out [][]string
	for _, str := range strArr {
		out = append(out, strings.Split(str, ""))
	}
	return out
}

func expandUniverse(input [][]string) [][]string {
	emptyColumnsIndexArr := getEmptySpaceColumns(input)
	emptyRowsIndexArr := getEmptySpaceRows(input)

	var extendedUniverse [][]string
	for i, row := range input {
		var newRow []string
		for j, col := range row {
			if contains(emptyColumnsIndexArr, j) {
				newRow = append(newRow, col)
				newRow = append(newRow, col)
			} else {
				newRow = append(newRow, col)
			}
		}

		if contains(emptyRowsIndexArr, i) {
			extendedUniverse = append(extendedUniverse, newRow)
			extendedUniverse = append(extendedUniverse, newRow)
		} else {
			extendedUniverse = append(extendedUniverse, newRow)
		}
	}
	return extendedUniverse
}

func getEmptySpaceRows(input [][]string) []int {
	var idxArr []int
	isEmptySpace := false
	for i, row := range input {
		for _, col := range row {
			if isNotEmptySpace(col) {
				isEmptySpace = false
				break
			} else {
				isEmptySpace = true
			}
		}
		if isEmptySpace {
			idxArr = append(idxArr, i)
		}
	}
	return idxArr
}
func getEmptySpaceColumns(input [][]string) []int {
	var out []int
	isEmptySpace := false
	for i := range input[0] {
		for _, row := range input {
			if isNotEmptySpace(row[i]) {
				isEmptySpace = false
				break
			} else {
				isEmptySpace = true
			}

		}
		if isEmptySpace {
			out = append(out, i)
		}
	}
	return out
}
func findAllGalaxies(input [][]string) []map[string]int {
	var out []map[string]int
	for i, row := range input {
		for j, col := range row {
			if col == "#" {
				coords := make(map[string]int)
				coords["x"] = i
				coords["y"] = j
				out = append(out, coords)
			}
		}
	}
	return out
}
func contains(arr []int, i int) bool {
	for _, num := range arr {
		if num == i {
			return true
		}
	}
	return false
}
func isNotEmptySpace(input string) bool {
	return input != "."
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
