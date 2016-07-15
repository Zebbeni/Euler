package main

import (
	"fmt"
	"math/big"
)

var wid = int64(20)
var hei = int64(20)

func main() {
	widFact := new(big.Int).MulRange(1, wid)
	heiFact := new(big.Int).MulRange(1, hei)
	numerator := new(big.Int).MulRange(1, wid+hei)
	denominator := new(big.Int).Mul(widFact, heiFact)
	paths := new(big.Int).Div(numerator, denominator)
	fmt.Println(paths)
}
