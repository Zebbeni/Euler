package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	for i := 2; i < 1000000; i++ {
		if isCircular(i) {
			fmt.Println("i:", i)
			count++
		}
	}
	fmt.Println("count:", count)
}

func isCircular(n int) bool {
	vals := arraytize(n)
	var mult, num, count int
	for i := 0; i < len(vals); i++ {
		mult, num, count = 1, 0, 0
		for d := i; count < len(vals); count++ {
			num += vals[d] * mult
			mult *= 10
			d = (d + 1) % len(vals)
		}
		if !isPrime(num) {
			return false
		}
	}
	return true
}

func arraytize(n int) []int {
	vals := []int{}
	for val := n; val > 0; {
		vals = append(vals, val%10)
		val = int(val / 10)
	}
	return vals
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
