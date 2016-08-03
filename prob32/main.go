package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 0; i <= 99999; i++ {
		if isProdPandigital(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isProdPandigital(n int) bool {
	for a := 2; a < int(math.Sqrt(float64(n)))+1; a++ {
		if n%a == 0 {
			if isPandigital(a, n/a, n) {
				fmt.Println("a:", a, "b:", n/a, "c:", n)
				return true
			}
		}
	}
	return false
}

// TODO: turn copy/pasted code into a function
func isPandigital(a, b, c int) bool {
	valMap := make(map[int]bool)
	digits := scrapeDigits(a)
	for digA := 0; digA < len(digits); digA++ {
		if digits[digA] == 0 || valMap[digits[digA]] {
			return false
		}
		valMap[digits[digA]] = true
	}
	digits = scrapeDigits(b)
	for digB := 0; digB < len(digits); digB++ {
		if digits[digB] == 0 || valMap[digits[digB]] {
			return false
		}
		valMap[digits[digB]] = true
	}
	digits = scrapeDigits(c)
	for digC := 0; digC < len(digits); digC++ {
		if digits[digC] == 0 || valMap[digits[digC]] {
			return false
		}
		valMap[digits[digC]] = true
	}
	return len(valMap) == 9
}

func scrapeDigits(i int) []int {
	intList := []int{}
	for val := i; val > 0; {
		intList = append(intList, val%10)
		val = int(val / 10)
	}
	return intList
}
