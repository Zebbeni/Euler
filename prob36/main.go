package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0
	for i := 1; i < 1000000; i++ {
		if isDoublePalindrome(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isDoublePalindrome(n int) bool {
	decimal := strconv.Itoa(n)
	binary := strconv.FormatInt(int64(n), 2)
	return isPalindrome(decimal) && isPalindrome(binary)
}

func isPalindrome(s string) bool {
	runes := arraytize(s)
	length := len(runes)
	for i := 0; i < (length/2)+1; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func arraytize(s string) []rune {
	runes := []rune{}
	for _, runeVal := range s {
		runes = append(runes, runeVal)
	}
	return runes
}
