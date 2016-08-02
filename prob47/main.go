package main

import (
	"fmt"
	"math"
)

const numConsecutive = 4
const targetPrimes = 4

func main() {
	testResult := 0
	for i := 2; i < 9223372036854775807; i++ {
		testResult = doTest(i)
		if testResult == 1 {
			fmt.Println("Found it:", i)
			break
		} else {
			i = testResult
		}
	}
}

func doTest(n int) int {
	for i := n; i < n+numConsecutive; i++ {
		if hasTargetPrimes(i) == false {
			return i
		}
	}
	return 1
}

func hasTargetPrimes(i int) bool {
	factorMap := make(map[int]int)
	for n := i; n > 1; {
		factor := chop(n)
		factorMap[factor] = 1
		if len(factorMap) > targetPrimes {
			return false
		}
		n = n / factor
	}
	return len(factorMap) == targetPrimes
}

func chop(num int) int {
	a := 2
	for a < int(math.Sqrt(float64(num)))+1 {
		if num%a == 0 {
			return a
		}
		a += 1
	}
	return num
}
