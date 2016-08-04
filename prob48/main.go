package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(0)
	for i := int64(1); i <= 1000; i++ {
		val := big.NewInt(i)
		sum.Add(sum, pow(val, i))
	}
	fmt.Println("sum:", sum.String())
}

func pow(n *big.Int, exp int64) *big.Int {
	prod := big.NewInt(1)
	for i := int64(1); i <= exp; i++ {
		prod.Mul(prod, n)
	}
	return prod
}
