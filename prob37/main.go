package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 10; i < 1000000; i++ {
		if passTest(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func passTest(n int) bool {
	mult := 1
	for val := n; val > 0; mult *= 10 {
		if !isPrime(val) {
			return false
		}
		val = int(val / 10)
	}
	for val := n % mult; val > 0; {
		if !isPrime(val) {
			return false
		}
		mult /= 10
		val = val % mult
	}
	return true
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
	a := 2
	for a < int(math.Sqrt(float64(num)))+1 {
		if num%a == 0 {
			return a
		}
		a++
	}
	return num
}
