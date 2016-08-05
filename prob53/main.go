package main

import (
	"fmt"
	"math/big"
)

func main() {
	count := 0
	cutoff := big.NewInt(1000000)
	for n := 1; n <= 100; n++ {
		fmt.Println("n:", n)
		for r := 1; r <= n; r++ {
			if combinatorial(n, r).Cmp(cutoff) == 1 {
				count++
			}
		}
	}
	fmt.Println("count:", count)
}

func combinatorial(n, r int) *big.Int {
	return new(big.Int).Div(fact(n), new(big.Int).Mul(fact(r), fact(n-r)))
}

func fact(n int) *big.Int {
	prod := new(big.Int).MulRange(1, int64(n))
	return prod
}
