package main

import (
	"fmt"
	"math"
)

const val = 1000

func main() {
	for a := 1; a <= val; a++ {
		for b := a + 1; b <= val; b++ {
			sqrt := math.Sqrt(float64(a*a + b*b))
			if isNaturalSquare(sqrt) {
				if a+b+int(sqrt) == val {
					fmt.Println("a:", a, "b:", b, "c:", int(sqrt))
					fmt.Println(a * b * int(sqrt))
				}
			}
		}
	}
}

func isNaturalSquare(sqrt float64) bool {
	return math.Ceil(sqrt) == math.Floor(sqrt)
}

func pow(n, x int) int {
	prod := 1
	for i := 0; i < x; i++ {
		prod *= n
	}
	return prod
}
