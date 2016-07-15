package main

import (
	"fmt"
	"math/big"
)

func main() {
	val := big.NewInt(2)
	prod := pow(val, 1000)
	string := prod.String()
	runes := []rune(string)
	sum := 0
	for i := range runes {
		sum += int(runes[i] - '0')
	}
	fmt.Println("prod:", prod, "\nsum:", sum)
}

func pow(n *big.Int, exp int) *big.Int {
	prod := big.NewInt(1)
	for i := 1; i <= exp; i++ {
		prod.Mul(prod, n)
	}
	return prod
}
