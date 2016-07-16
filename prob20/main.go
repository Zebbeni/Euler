package main

import (
	"fmt"
	"math/big"
)

func main() {
	factorial := new(big.Int).MulRange(1, 100)
	bytes := factorial.Append([]byte{}, 10)
	sum := 0
	for b := range bytes {
		sum += int(bytes[b] - '0')
	}
	fmt.Println("sum:", sum)
}
