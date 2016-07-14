package main

import (
	"fmt"
	"math"
)

const val = 2000000

func main() {
	sum := 0
	for i := 2; i < val; i++ {
		if chop(i) == i {
			sum += i
		}
	}
	fmt.Println(sum)
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
