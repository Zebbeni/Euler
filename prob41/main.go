package main

import (
	"fmt"
	"math"
)

func main() {
	max, mult, n := 0, 10, 1
	for i := 1; i < 9999999; i++ {
		if mult/i == 0 {
			mult *= 10
			n++
		}
		if isNPandigital(i, n) && isPrime(i) {
			max = i
		}
	}
	fmt.Println("max:", max)
}

func isNPandigital(i, n int) bool {
	valMap := make(map[int]bool)
	digits := scrapeDigits(i)
	for dig := 0; dig < len(digits); dig++ {
		val := digits[dig]
		if val == 0 || val > n || valMap[val] {
			return false
		}
		valMap[val] = true
	}
	return len(valMap) == n
}

func scrapeDigits(i int) []int {
	intList := []int{}
	for val := i; val > 0; {
		intList = append(intList, val%10)
		val = int(val / 10)
	}
	return intList
}

func isPrime(n int) bool {
	if n == 0 {
		return false
	}
	if n <= 1 {
		return false
	}
	if chop(n) == n {
		return true
	}
	return false
}

func chop(num int) int {
	for a := 2; a < int(math.Sqrt(float64(num)))+1; a++ {
		if num%a == 0 {
			return a
		}
	}
	return num
}
