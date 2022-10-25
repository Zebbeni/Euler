package main

import (
	"fmt"
	"math"
	"strconv"
)

const maxValue = 10000

var minSum = 9999999999
var bestSet []int

var primes []int

func main() {
	populatePrimes()
	for i, a := range primes {
		if a > minSum {
			continue
		}
		for j, b := range primes[i+1:] {
			if a+b > minSum {
				continue
			}
			if !testPair(a, b) {
				continue
			}
			for k, c := range primes[i+j+2:] {
				if a+b+c > minSum {
					continue
				}
				if !testPair(a, c) {
					continue
				}
				if !testPair(b, c) {
					continue
				}
				for l, d := range primes[i+j+k+3:] {
					if a+b+c+d > minSum {
						continue
					}
					if !testPair(a, d) {
						continue
					}
					if !testPair(b, d) {
						continue
					}
					if !testPair(c, d) {
						continue
					}
					for _, e := range primes[i+j+k+l+4:] {
						if a+b+c+d+e > minSum {
							continue
						}
						if !testPair(a, e) {
							continue
						}
						if !testPair(b, e) {
							continue
						}
						if !testPair(c, e) {
							continue
						}
						if !testPair(d, e) {
							continue
						}
						minSum = a + b + c + d + e
						bestSet = []int{a, b, c, d, e}
						fmt.Println("BestSet: ", bestSet, "sum: ", minSum)
					}
				}
			}
		}
	}
}

func populatePrimes() {
	primes = make([]int, 0)
	for i := 2; i < maxValue; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
}

func testPair(x, y int) bool {
	concatenated, _ := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))
	if !isPrime(concatenated) {
		return false
	}
	concatenated, _ = strconv.Atoi(strconv.Itoa(y) + strconv.Itoa(x))
	return isPrime(concatenated)
}

func isPrime(x int) bool {
	if x <= 3 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
