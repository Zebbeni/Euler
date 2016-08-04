package main

import (
	"fmt"
	"math"
)

func main() {
	gap := 3330
	for i := 1000; i <= 9999; i++ {
		if isPrime(i) {
			if isPrime(i + gap) {
				if isPrime(i + 2*gap) {
					fmt.Println(i, i+gap, i+gap+gap)
				}
			}
		}
	}
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
