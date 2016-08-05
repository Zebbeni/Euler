package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		test(i, arrayify(i))
	}
}

func test(n int, digits []int) {
	for i := 2; i <= 6; i++ {
		if check(arrayify(n*i), digits) == false {
			return
		}
	}
	fmt.Println("found it:", n)
}

func check(digitsN, digits []int) bool {
	if len(digitsN) != len(digits) {
		return false
	}
	var digit int
	for len(digits) > 0 {
		digit = digits[0]
		digits = digits[1:]
		for j := 0; j < len(digitsN); j++ {
			if digit == digitsN[j] {
				digitsN = append(digitsN[:j], digitsN[j+1:]...)
				break
			}
		}
	}
	return len(digitsN) == len(digits)
}

func arrayify(n int) []int {
	digits := []int{}
	for val := n; val > 0; val = val / 10 {
		digits = append(digits, val%10)
	}
	return digits
}
