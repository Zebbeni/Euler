package main

import (
	"fmt"
	"math/big"
)

var numMoreDigitsCount = 0

func main() {
	_, _ = calcNumDenom(1000)
	fmt.Println("numMoreDigitsCount:", numMoreDigitsCount)
}

func calcNumDenom(iterationsLeft int) (newNum *big.Int, newDenom *big.Int) {
	if iterationsLeft == 1 {
		newNum, newDenom = big.NewInt(1), big.NewInt(2)
		checkIteration(iterationsLeft, newNum, newDenom)
		return
	}

	num, denom := calcNumDenom(iterationsLeft - 1)
	newDenom = big.NewInt(0).Mul(denom, big.NewInt(2))
	newDenom = newDenom.Add(newDenom, num)
	newNum = denom

	checkIteration(iterationsLeft, newNum, newDenom)
	return
}

func checkIteration(i int, num, denom *big.Int) {
	newNum := big.NewInt(0).Add(num, denom)
	newDenom := denom
	if len(newNum.String()) > len(newDenom.String()) {
		numMoreDigitsCount++
	}

	fmt.Printf("\nIteration %d: 1 + %d/%d => %d/%d", i, num, denom, big.NewInt(0).Add(num, denom), denom)
}
