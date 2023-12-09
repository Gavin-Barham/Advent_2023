package main

import (
	sc "strconv"
	uni "unicode"
)

func scanForSpecialChars(s string) bool {
	for _, c := range s {
		if !uni.IsNumber(c) && string(c) != "." {
			return true
		}
	}
	return false
}

func problemOne(input []string) int {
	sum := 0
	for i, str := range input {
		numberSequence := ""
		for j, char := range str {
			if uni.IsNumber(char) {
				numberSequence += string(char)
			}
			if !uni.IsNumber(char) || j == len(str)-1 {
				if len(numberSequence) == 0 {
					continue
				}
				var start int
				var end int

				if j-len(numberSequence)-1 <= 0 {
					start = j - len(numberSequence)
				} else {
					start = j - len(numberSequence) - 1
				}

				if j >= len(str)-1 {
					end = j
				} else {
					end = j + 1
				}

				if i == 0 {
					if scanForSpecialChars(input[i][start:end]) || scanForSpecialChars(input[i+1][start:end]) {
						num, err := sc.Atoi(numberSequence)
						check(err)
						// fmt.Printf("i: %v, number: %v, length: %v, start:%v, end:%v \n\n", i, num, len(numberSequence), start, end)
						sum += num
					}

				} else if i == len(input)-1 {
					if scanForSpecialChars(input[i-1][start:end]) || scanForSpecialChars(input[i][start:end]) {
						num, err := sc.Atoi(numberSequence)
						check(err)
						// fmt.Printf("i: %v, number: %v, length: %v, start:%v, end:%v \n\n", i, num, len(numberSequence), start, end)
						sum += num
					}
				} else {
					if scanForSpecialChars(input[i-1][start:end]) || scanForSpecialChars(input[i][start:end]) || scanForSpecialChars(input[i+1][start:end]) {
						num, err := sc.Atoi(numberSequence)
						check(err)
						// fmt.Printf("i: %v, number: %v, length: %v, start:%v, end:%v \n\n", i, num, len(numberSequence), start, end)
						sum += num
					}
				}
				numberSequence = ""
			}
		}
	}
	return sum
}
