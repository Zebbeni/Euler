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

// a map of numbers to a value indicating whether that number has been found to be
// Lychrel or non-Lychrel. Value is set to the number of reverse-and-add iterations
// needed to find a palindrome for the given number (-1 for Lychrel numbers)
// (ie. a reverse-and-add iteration resulted in a palindrome)
var lychrelTestMap = make(map[string]int)

func main() {
	lychrelCount := 0
	for i := maxNum; i >= minNum; i-- {
		fmt.Printf("\ni = %d:", i)
		isLychrel, tries := testLychrel(i)
		lychrelTestMap[strconv.Itoa(i)] = tries
		if isLychrel {
			lychrelCount++
		}
	}
	fmt.Printf("\nLychrel numbers found: %d\n", lychrelCount)
}

// determines if a given integer x is Lychrel and- if not- the number of
// reverse-and-add iterations needed to find a palindrome starting with x
// returns 0 if Lychrel
func testLychrel(x int) (bool, int) {
	n := big.NewInt(int64(x))

	nString := n.String()
	if isPalindrome(nString) {
		lychrelTestMap[nString] = 0
		return false, 0
	}

	for tries := 1; tries <= maxTries; tries++ {
		nReverse := reverse(n)
		sum := big.NewInt(0).Add(n, nReverse)
		fmt.Printf("\n  %2d.) %d + %d = %d", tries, n, nReverse, sum)

		sumString := sum.String()
		if itersToPalindrome, isKnown := lychrelTestMap[sumString]; isKnown {
			if itersToPalindrome == -1 {
				fmt.Printf("... known lychrel!")
				return true, -1
			}
			if itersToPalindrome+tries > maxTries {
				fmt.Printf("... new lychrel found, too many tries needed!")
				return true, -1
			}
			fmt.Printf("... known non-lychrel! %d + %d", tries, itersToPalindrome)
			return false, itersToPalindrome + tries

		}

		if isPalindrome := isPalindrome(sumString); isPalindrome {
			fmt.Printf("... new non-lychrel!")
			return false, tries
		}
		n = sum
	}
	fmt.Printf("... new lychrel found!")
	return true, -1
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
