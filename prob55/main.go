package main

import (
	"fmt"
	"math/big"
)

const (
	maxNum   = 10000
	minNum   = 1
	maxTries = 50
)

func main() {
	lychrelCount := 0
	for i := minNum; i <= maxNum; i++ {
		isLychrel := testLychrel(i)
		if isLychrel {
			lychrelCount++
		}
	}
	fmt.Println("lychrelCount:", lychrelCount)
}

// determines if a given integer x is Lychrel and- if not- the number of
// reverse-and-add iterations needed to find a palindrome starting with x
// returns 0 if Lychrel
func testLychrel(x int) bool {
	n := big.NewInt(int64(x))
	for tries := 1; tries <= maxTries; tries++ {
		nReverse := reverse(n)
		sum := big.NewInt(0).Add(n, nReverse)
		if isPalindrome := isPalindrome(sum.String()); isPalindrome {
			return false
		}
		n = sum
	}
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
