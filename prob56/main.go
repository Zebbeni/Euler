package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	maxDigitalSum := 0
	for a := 1; a < 100; a++ {
		bigA := big.NewInt(int64(a)) // stay consistent
		prodA := bigA
		for b := 2; b < 100; b++ {
			prodA = big.NewInt(0).Mul(prodA, bigA)
			runes := []rune(prodA.String())

			sum := 0
			for _, thisChar := range runes {
				val, _ := strconv.Atoi(string(thisChar))
				sum += val
			}

			fmt.Printf("\n%d^%d = %s, sum: %d", a, b, prodA.String(), sum)

			if sum > maxDigitalSum {
				maxDigitalSum = sum
				fmt.Print(" ...new maxDigitalSum:", maxDigitalSum)
			}
		}
	}
	fmt.Println("\n\nFinal:", maxDigitalSum)
}
