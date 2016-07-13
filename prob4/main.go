package main

import (
	"fmt"
	"strconv"
)

func main() {
	max, prod := 0, 0
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			prod = i * j
			if isPalindrome(prod) && prod > max {
				max = prod
			}
		}
	}
	fmt.Println(max)
}

func isPalindrome(num int) bool {
	string := strconv.Itoa(num)
	i, j := 0, len(string)-1
	for i <= j {
		if string[i] != string[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
