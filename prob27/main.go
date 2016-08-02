package main

import (
	"fmt"
	"math"
)

func main() {
	var max, maxA, maxB, value, count int
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			count = 0
			for n := 0; ; n++ {
				value = (n * n) + (a * n) + b
				if isPrime(value) {
					count++
				} else {
					if count > max {
						maxA = a
						maxB = b
						max = count
					}
					break
				}
			}
		}
	}
	fmt.Println("a:", maxA, "b:", maxB)
	fmt.Println("product:", maxA*maxB)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if chop(n) == n {
		return true
	}
	return false
}

func chop(num int) int {
	a := 2
	for a < int(math.Sqrt(float64(num)))+1 {
		if num%a == 0 {
			return a
		}
		a++
	}
	return num
}
