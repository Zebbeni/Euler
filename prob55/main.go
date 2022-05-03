package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	maxNum   = 10000
	minNum   = 1
	maxTries = 50
)

// a map of numbers and whether they've been found to be Lychrel or non-Lychrel
// true if Lychrel, false if Non-Lychrel
// (ie. a reverse-and-add iteration resulted in a palindrome)
var lychrelTestMap = make(map[string]bool)

func main() {
	lychrelCount := 0
	for i := maxNum; i >= minNum; i-- {
		fmt.Printf("\ni = %d:", i)
		isLychrel := testLychrel(i)
		lychrelTestMap[strconv.Itoa(i)] = isLychrel
		if isLychrel {
			lychrelCount++
		}
	}
	fmt.Printf("\nLychrel numbers found: %d", lychrelCount)
}

func testLychrel(x int) bool {
	n := big.NewInt(int64(x))
	for tries := 0; tries < maxTries; tries++ {
		nReverse := reverse(n)
		sum := big.NewInt(0).Add(n, nReverse)
		fmt.Printf("\n  (%2d)...%d + %d = %d", tries+1, n, nReverse, sum)

		sumString := sum.String()
		if isLychrel, isKnown := lychrelTestMap[sumString]; isKnown {
			toPrint := "non-lychrel"
			if isLychrel {
				toPrint = "lychrel"
			}
			fmt.Printf("... known %s", toPrint)

			return isLychrel
		}

		if isPalindrome := isPalindrome(sumString); isPalindrome {
			fmt.Printf("... new non-lychrel!")
			return false
		}
		n = sum
	}
	fmt.Printf("... new lychrel found!")
	return true
}

func reverse(x *big.Int) *big.Int {
	xOriginalRunes := []rune(x.String())
	xReversedRunes := make([]rune, len(xOriginalRunes))
	for i, v := range xOriginalRunes {
		xReversedRunes[len(xOriginalRunes)-1-i] = v
	}
	reverseInt := new(big.Int)
	reverseInt, _ = reverseInt.SetString(string(xReversedRunes), 10)
	return reverseInt
}

// return true or false if the given string is palindromic,
// and a bool saying whether the value was previously tested
func isPalindrome(x string) bool {
	runes := []rune(x)
	runeCount := len(runes)
	for i := 0; i < runeCount/2; i++ {
		if runes[i] != runes[runeCount-1-i] {
			return false
		}
	}
	return true
}
