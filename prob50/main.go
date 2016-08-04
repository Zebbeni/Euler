package main

import (
	"fmt"
	"math"
)

var primes = []int{2}

func main() {
	sum, max, val := 2, 0, 0
	for i := 3; i < 500000; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
			sum += i
			val = sum
			for j := 0; j < len(primes); j++ {
				if val < 1000000 {
					if val > max {
						if isPrime(val) {
							max = val
							break
						} else {
							val -= primes[j]
						}
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println("final max:", max)
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
