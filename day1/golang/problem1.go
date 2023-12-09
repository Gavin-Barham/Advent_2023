package main

import (
	"os"
	sc "strconv"
	s "strings"
)

func problemOne() int {
	file, err := os.ReadFile("input1.txt")
	check(err)
	var strArr []string = s.Split(string(file), "\n")
	var out []string
	sum := 0
	for i := 0; i < len(strArr); i++ {
		str := strArr[i]
		first := ""
		last := ""
		for idx := 0; idx < len(str); idx++ {
			char := str[idx]
			if isNumeric(char) {
				if first == "" {
					first = string(char)
				} else {
					last = string(char)
				}
			}

		}
		if last == "" {
			last = first
		}
		out = append(out, first+last)
	}
	for i := 0; i < len(out); i++ {
		num, err := sc.Atoi(out[i])
		check(err)
		sum = sum + num
	}
	return sum
}
