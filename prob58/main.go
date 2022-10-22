package main

import (
	"fmt"
	"math"
)

var primes = 0
var total = 1

func main() {
	var x1, x2, x3, x4 int
	fmt.Printf("\n1: %10d", 1)
	for w := 1; w < 100000; w += 2 {
		x1 = (w * w) + 1 + w
		x2 = x1 + 1 + w
		x3 = x2 + 1 + w
		x4 = x3 + 1 + w

		if isPrime(x1) {
			primes++
		}
		if isPrime(x2) {
			primes++
		}
		if isPrime(x3) {
			primes++
		}
		if isPrime(x4) {
			primes++
		}
		total += 4

		ratio := float64(primes) / float64(total)
		fmt.Printf("\n%5d: %10d, %10d, %10d, %10d, %5d/%d: %10f", w+2, x1, x2, x3, x4, primes, total, ratio)

		if ratio < 0.1 {
			return
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
