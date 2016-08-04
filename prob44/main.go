package main

import (
	"fmt"
	"math"
)

func main() {
	minD := 999999999
	for k := 1; k < 99999999; k++ {
		pK := genPentagonal(k)
		for j := k - 1; j > 0; j-- {
			pJ := genPentagonal(j)
			if pK-pJ < minD {
				if isPentagonal(pK-pJ) && isPentagonal(pK+pJ) {
					minD = pK - pJ
					fmt.Println("pK:", pK, "pJ", pJ, "minD:", minD)
					break
				}
			} else {
				break
			}
		}
	}
	fmt.Println("minD:", minD)
}

func genPentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func isPentagonal(p int) bool {
	quad := 24*p + 1
	sqrt := int(math.Sqrt(float64(quad)))
	return sqrt*sqrt == quad && (sqrt+1)%6 == 0
}
