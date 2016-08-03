package main

import (
	"fmt"
	"strconv"
)

func main() {
	decimals := []uint8{}
	for i := 1; i < 2000000; i++ {
		for _, val := range strconv.Itoa(i) {
			decimals = append(decimals, uint8(val-'0'))
		}
	}
	prod := uint8(1)
	for i := 1; i <= 1000000; i *= 10 {
		prod *= decimals[i-1]
	}
	fmt.Println("prod:", prod)
}
