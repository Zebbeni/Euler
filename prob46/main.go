package main

import (
	"fmt"
	"math"
)

var primes = []int{}

func main() {
	for i := 3; ; i += 2 {
		if !isPrime(i) {
			if !passTest(i) {
				fmt.Println(i, "cannot be done")
				return
			}
		} else {
			primes = append(primes, i)
		}
	}
}

func passTest(n int) bool {
	for i := 0; i < len(primes); i++ {
		remaining := n - primes[i]
		if remaining%2 == 0 && isSquare(remaining/2) {
			return true
		}
	}
	return false
}

func isSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func isPrime(n int) bool {
	return n > 1 && chop(n) == n
}

func chop(num int) int {
	for a := 2; a < int(math.Sqrt(float64(num)))+1; a++ {
		if num%a == 0 {
			return a
		}
	}
	return num
}
