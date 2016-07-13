package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	count, n, x := 0, 10001, 1
	for count < n {
		x++
		if x == chop(x) {
			fmt.Println(x)
			count++
		}
	}
	fmt.Println(time.Since(start))
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
