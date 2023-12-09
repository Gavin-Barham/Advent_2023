package main

import (
	"os"
	sc "strconv"
	s "strings"
)

func wordToNumeric(word string) int {
	wordMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	out := ""
	for i := range word {
		var currWord string = string(word[i])
		if isNumeric(word[i]) {
			out += currWord
			continue
		}

		for j := i + 1; j < len(word); j++ {
			if isNumeric(word[j]) {
				break
			}
			currWord += string(word[j])
			if wordMap[currWord] != 0 {
				out += sc.Itoa(wordMap[currWord])
				break
			}
		}
	}
	first := ""
	last := ""
	for i := 0; i < len(out); i++ {
		char := out[i]
		if first == "" {
			first = string(char)
		} else {
			last = string(char)
		}
	}
	if last == "" {
		last = first
	}
	num, err := sc.Atoi(first + last)
	check(err)
	return num
}

func problemTwo() int {
	file, err := os.ReadFile("input1.txt")
	check(err)
	var strArr []string = s.Split(string(file), "\n")
	sum := 0
	for i := 0; i < len(strArr); i++ {
		sum += wordToNumeric(strArr[i])
	}
	return sum
}
