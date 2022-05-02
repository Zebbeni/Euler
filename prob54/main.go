package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

const filename = "poker.txt"

type rank int

const (
	rankHighCard = iota
	rankOnePair
	rankTwoPair
	rankThreeOfAKind
	rankStraight
	rankFlush
	rankFullHouse
	rankFourOfAKind
	rankStraightFlush
	rankRoyalFlush
)

var cardValueMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type hand struct {
	values  []int
	isFlush bool
}

func main() {
	start := time.Now()
	p1WinCount, p2WinCount := 0, 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hand1, hand2 := parseHands(line)
		if p1Won := compareHands(hand1, hand2); p1Won {
			p1WinCount += 1
		} else {
			p2WinCount += 1
		}
	}

	fmt.Printf("exec time: %v", time.Since(start))
	fmt.Printf("\np1: %d, p2: %d", p1WinCount, p2WinCount)
}

func parseHands(text string) (hand, hand) {
	cardStrings := strings.Split(text, " ")
	h1 := parseHand(cardStrings[:5])
	h2 := parseHand(cardStrings[5:])
	return h1, h2
}

func parseHand(cardStrings []string) hand {
	values := make([]int, 5)
	isFlush := true
	suit := ""
	for i, cardString := range cardStrings {
		cardValStrings := strings.Split(cardString, "")

		cardValue := cardValStrings[0]
		values[i] = cardValueMap[cardValue]

		cardSuit := cardValStrings[1]
		if i == 0 {
			suit = cardSuit
		} else if suit != cardSuit {
			isFlush = false
		}
	}

	return hand{values: values, isFlush: isFlush}
}

// return true iff hand 1 beats hand 2
// (problem 54 assumes no ties)
func compareHands(h1, h2 hand) bool {
	rank1, h1v1, h1v2, h1v3, h1v4, h1v5 := evaluateHand(h1)
	rank2, h2v1, h2v2, h2v3, h2v4, h2v5 := evaluateHand(h2)
	if rank1 != rank2 {
		return rank1 > rank2
	} else if h1v1 != h2v1 {
		return h1v1 > h2v1
	} else if h1v2 != h2v2 {
		return h1v2 > h2v2
	} else if h1v3 != h2v3 {
		return h1v3 > h2v3
	} else if h1v4 != h2v4 {
		return h1v4 > h2v4
	}
	return h1v5 > h2v5
}

// return the type, the high card of the cards
func evaluateHand(h hand) (rank, v1, v2, v3, v4, v5 int) {
	sort.Ints(h.values)

	isStraight := true
	// map of values to the number of times they appear
	valueCounts := make(map[int]int)
	// map of counts to the different value(s) - ordered least to greatest - that have that count
	countValues := make(map[int][]int)

	for i, v := range h.values {
		if _, ok := valueCounts[v]; !ok {
			valueCounts[v] = 0
		}
		valueCounts[v]++

		if i > 0 {
			prevValue := h.values[i-1]
			if v > prevValue {
				prevCount := valueCounts[prevValue]
				if _, ok := countValues[prevCount]; ok {
					countValues[prevCount] = append(countValues[prevCount], prevValue)
				} else {
					countValues[prevCount] = []int{prevValue}
				}
			}

			if i == 4 {
				count := valueCounts[v]
				if _, ok := countValues[count]; ok {
					countValues[count] = append(countValues[count], v)
				} else {
					countValues[count] = []int{v}
				}
			}

			if v != prevValue+1 {
				isStraight = false
			}
		}
	}

	if _, found4 := countValues[4]; found4 {
		return rankFourOfAKind, countValues[4][0], countValues[1][0], 0, 0, 0
	}
	if _, found3 := countValues[3]; found3 {
		if _, found2 := countValues[2]; found2 {
			return rankFullHouse, countValues[3][0], countValues[2][0], 0, 0, 0
		} else {
			return rankThreeOfAKind, countValues[3][0], countValues[1][1], countValues[1][0], 0, 0
		}
	}
	if _, found2 := countValues[2]; found2 {
		if len(countValues[2]) > 1 {
			return rankTwoPair, countValues[2][1], countValues[2][0], 0, 0, 0
		} else {
			return rankOnePair, countValues[2][0], countValues[1][2], countValues[1][1], countValues[1][0], 0
		}
	}
	if isStraight {
		if h.isFlush {
			if h.values[4] == cardValueMap["A"] {
				return rankRoyalFlush, 0, 0, 0, 0, 0
			} else {
				return rankStraightFlush, countValues[1][4], countValues[1][3], countValues[1][2], countValues[1][1], countValues[1][0]
			}
		} else {
			return rankStraight, countValues[1][4], countValues[1][3], countValues[1][2], countValues[1][1], countValues[1][0]
		}
	} else if h.isFlush {
		return rankFlush, countValues[1][4], countValues[1][3], countValues[1][2], countValues[1][1], countValues[1][0]
	}
	return rankHighCard, countValues[1][4], countValues[1][3], countValues[1][2], countValues[1][1], countValues[1][0]
}
