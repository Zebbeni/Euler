package main

import (
	"fmt"
	"math"
)

func main() {
	max, factor := 0, 0
	num := 600851475143
	for num != 1 {
		factor = chop(num)
		num = num / factor
		if factor > max {
			max = factor
		}
	}
	fmt.Print("Max:", max)
}

func chop(num int) int {
	a := 2
	for a < int(math.Sqrt(float64(num)))+1 {
		if num%a == 0 {
			return a
		}
		a += 1
	}
	return num
}
