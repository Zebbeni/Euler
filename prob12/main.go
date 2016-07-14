package main

import (
	"fmt"
	"math"
)

const factors = 500

func main() {
	n := uint64(0)
	for i := uint64(1); ; i++ {
		n += i
		count := numFactors(n)
		if count >= factors {
			fmt.Println("found it:", n)
			break
		}
	}
}

func numFactors(n uint64) int {
	count := 0
	for i := uint64(1); i <= uint64(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			count++
			if n/i != i {
				count++
			}
		}
	}
	return count
}
