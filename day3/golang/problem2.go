package main

import (
	sc "strconv"
	uni "unicode"
)

func isAdjacentNum(start int, end int, target int) bool {
	return start == target || start == target-1 || start == target+1 || end == target || end == target-1 || end == target+1
}
func addNumToList(str string, mid int) []int {
	var numList []int
	numSequence := ""
	seqStart := -1
	seqEnd := -1
	for idx, char := range str {
		if uni.IsNumber(char) {
			if seqStart == -1 {
				seqStart = idx
			}
			seqEnd = idx
			numSequence += string(char)
		}
		if !uni.IsNumber(char) || idx == len(str)-1 {
			if len(numSequence) != 0 {
				if isAdjacentNum(seqStart, seqEnd, mid) {
					num, err := sc.Atoi(numSequence)
					check(err)
					numList = append(numList, num)
				}
				numSequence = ""
				seqStart = -1
				seqEnd = -1
			}
		}
	}
	return numList
}
func getAdjacentNum(matrix []string, i int, j int, numList []int) []int {
	if i == 0 {
		nums := addNumToList(matrix[i], j)
		numList = append(numList, nums...)

		nums = addNumToList(matrix[i+1], j)
		numList = append(numList, nums...)

	} else if i == len(matrix)-1 {
		nums := addNumToList(matrix[i-1], j)
		numList = append(numList, nums...)

		nums = addNumToList(matrix[i], j)
		numList = append(numList, nums...)

	} else {
		nums := addNumToList(matrix[i-1], j)
		numList = append(numList, nums...)

		nums = addNumToList(matrix[i], j)
		numList = append(numList, nums...)

		nums = addNumToList(matrix[i+1], j)
		numList = append(numList, nums...)

	}
	return numList
}

func problemTwo(input []string) int {
	sum := 0
	for i, str := range input {
		for j, char := range str {
			var numList []int
			if string(char) != "*" {
				continue
			} else {
				numList = getAdjacentNum(input, i, j, numList)
				if len(numList) == 2 {
					num := numList[0] * numList[1]
					sum += num
				}

			}
		}
	}
	return sum
}
