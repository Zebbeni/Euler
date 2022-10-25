package main

import (
	"fmt"
	"math/big"
)

func main() {
	total := 0
	for n := int64(1); n <= 40; n++ {
		count := 0
		for x := int64(1); x < 1000; x++ {
			bigX := big.NewInt(x)
			p := bigX.Exp(bigX, big.NewInt(n), nil)
			digits := len(p.String())
			if int64(digits) == n {
				count++
			} else if int64(digits) > n {
				break
			}
		}
		fmt.Printf("\nn %d: %d", n, count)
		if count == 0 {
			break
		}
		total += count
	}
	fmt.Println("\n\ntotal:", total)
}
