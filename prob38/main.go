package main

import (
	"fmt"
	"strconv"
)

const concatlen = 9

func main() {
	for n := 9; n < 100000; n++ {
		val := getProductConcat(n)
		if isPandigital(val) {
			fmt.Print("\n")
			for _, r := range val {
				fmt.Print(uint64(r - '0'))
			}
			fmt.Print("\n")
		}
	}
}

func isPandigital(runes []rune) bool {
	mapRunes := make(map[uint64]bool)
	if len(runes) != concatlen {
		return false
	}
	for _, runeVal := range runes {
		val := uint64(runeVal - '0')
		if val == 0 || mapRunes[val] {
			return false
		}
		mapRunes[val] = true
	}
	return true
}

func getProductConcat(n int) []rune {
	concString := []rune{}
	runes := []rune{}
	for i := 1; i < 10000; i++ {
		runes = arraytize(strconv.Itoa(n * i))
		concString = append(concString, runes...)
		if len(concString) >= concatlen {
			return concString
		}
	}
	fmt.Println("something bad happened")
	return nil
}

func arraytize(s string) []rune {
	runes := []rune{}
	for _, runeVal := range s {
		runes = append(runes, runeVal)
	}
	return runes
}
