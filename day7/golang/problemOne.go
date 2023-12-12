package main

import (
	"sort"
	"strconv"
	s "strings"
)

type Game struct {
	hand   string
	rank   int
	bet    int
	spread map[string]int
}

func problemOne(input []string) int {
	var cardValMap = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
	out := 0
	formattedInput := formatInput(input, false, cardValMap)
	sortedHands := sortAllHandsByRank(formattedInput, cardValMap)
	for i, hand := range sortedHands {
		rank := len(sortedHands) - i
		out += rank * hand.bet
	}
	return out
}
func formatInput(input []string, isProblemTwo bool, cardValMap map[string]int) []Game {
	var out []Game
	for _, gameStr := range input {
		game := s.Split(gameStr, " ")
		hand := game[0]
		bet, err := strconv.Atoi(game[1])
		check(err)
		var cardFreq map[string]int
		if isProblemTwo {
			cardFreq = getCardFreqProblemTwo(hand, cardValMap)
		} else {
			cardFreq = getCardFreq(hand)
		}

		rank := determineHandType(cardFreq)
		currGame := Game{
			hand:   hand,
			bet:    bet,
			rank:   rank,
			spread: cardFreq,
		}
		out = append(out, currGame)
	}
	return out
}

func getCardFreq(hand string) map[string]int {
	freq := make(map[string]int)
	for _, card := range hand {
		strCard := string(card)
		if freq[strCard] == 0 {
			freq[strCard] = 1
		} else {
			freq[strCard]++
		}
	}
	return freq
}

func determineHandType(freq map[string]int) int {
	if len(freq) == 1 {
		return 7
	} else if len(freq) == 2 {
		for _, v := range freq {
			if v == 1 || v == 4 {
				return 6
			} else {
				return 5
			}
		}
	} else if len(freq) == 3 {
		for _, v := range freq {
			if v == 3 {
				return 4
			} else if v == 2 {
				return 3
			}
		}
	} else if len(freq) == 4 {
		return 2
	}
	return 1
}

func sortAllHandsByRank(hands []Game, cardValMap map[string]int) []Game {
	sortByRank := func(i int, j int) bool {
		if hands[i].rank != hands[j].rank {
			return hands[i].rank > hands[j].rank
		} else {
			for k, v := range hands[i].hand {
				ACard := string(v)
				Bcard := string(hands[j].hand[k])
				if cardValMap[ACard] == cardValMap[Bcard] {
					continue
				}
				return cardValMap[ACard] > cardValMap[Bcard]
			}
		}
		return false
	}
	sort.SliceStable(hands, sortByRank)
	return hands
}
