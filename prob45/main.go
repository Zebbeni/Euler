package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 285 + 1; ; i++ {
		tI := genTri(i)
		if isPent(tI) && isHex(tI) {
			fmt.Println("tI:", tI, "i:", i)
			return
		}
	}
}

func genTri(n int) int {
	return n * (n + 1) / 2
}

func isPent(p int) bool {
	quad := 24*p + 1
	sqrt := int(math.Sqrt(float64(quad)))
	return sqrt*sqrt == quad && (sqrt+1)%6 == 0
}

func isHex(h int) bool {
	quad := 8*h + 1
	sqrt := int(math.Sqrt(float64(quad)))
	return sqrt*sqrt == quad && (sqrt+1)%4 == 0
}
