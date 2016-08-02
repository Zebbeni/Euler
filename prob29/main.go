package main

import (
	"fmt"
	"math/big"
)

func main() {
	valMap := make(map[string]bool)
	var val string
	for a := 2; a <= 1000; a++ {
		for b := 2; b <= 1000; b++ {
			val = pow(a, b).String()
			if !valMap[val] {
				valMap[val] = true
			}
		}
	}

	fmt.Println("number of Terms:", len(valMap))
}

func pow(n, exp int) *big.Int {
	prod := big.NewInt(1)
	bigN := big.NewInt(int64(n))
	for i := 1; i <= exp; i++ {
		prod.Mul(prod, bigN)
	}
	return prod
}
