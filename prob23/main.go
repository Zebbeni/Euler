package main

import (
	"fmt"
	"math"
)

const max = 28123

var abundant = []int{}

func main() {
	for i := 0; i <= max; i++ {
		if sumDivisors(i) > i {
			abundant = append(abundant, i)
		}
	}
	// fmt.Println("abundant:", abundant)
	isSum := false
	sum, a1, a2 := 0, 0, 0
	for i := 0; i < max; i++ {
		isSum = false
		for a := range abundant {
			a1 = abundant[a]
			for b := a; b < len(abundant) && !isSum; b++ {
				a2 = abundant[b]
				if a1+a2 == i {
					fmt.Println("i:", i, "a1", a1, "a2", a2)
					isSum = true
				} else if a1+a2 > i {
					break
				}
			}
		}
		if !isSum {
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}

func sumDivisors(n int) int {
	sum := 0
	for i := 1; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 && n != i {
			sum += i
			if n/i != n && n/i != i {
				sum += n / i
			}
		}
	}
	return sum
}
