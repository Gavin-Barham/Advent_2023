package main

func problemTwo(input []string) int {
	var cardValMap = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}
	out := 0
	formattedInput := formatInput(input, true, cardValMap)
	sortedHands := sortAllHandsByRank(formattedInput, cardValMap)
	for i, hand := range sortedHands {
		rank := len(sortedHands) - i
		out += rank * hand.bet
	}
	return out
}

func getCardFreqProblemTwo(hand string, cardValMap map[string]int) map[string]int {
	freq := make(map[string]int)
	for _, card := range hand {
		strCard := string(card)
		if freq[strCard] == 0 {
			freq[strCard] = 1
		} else {
			freq[strCard]++
		}
	}
	if freq["J"] > 0 {
		numOfJacks := freq["J"]
		delete(freq, "J")
		highestFreq := findMostFrequent(freq)
		var highestValues []string
		for k, v := range freq {
			if v == highestFreq {
				highestValues = append(highestValues, k)
			}
		}
		if len(highestValues) == 1 {
			freq[highestValues[0]] += numOfJacks
		} else {
			highestValueCard := 0
			currHighest := ""
			for _, s := range highestValues {
				if cardValMap[s] > highestValueCard {
					highestValueCard = cardValMap[s]
					currHighest = s
				}
			}
			freq[currHighest] += numOfJacks
		}
		return freq
	}
	return freq
}

func findMostFrequent(inputMap map[string]int) int {
	var highestValue int

	for _, value := range inputMap {
		if value > highestValue {
			highestValue = value
		}
	}

	return highestValue
}
