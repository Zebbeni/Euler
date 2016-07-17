package main

import (
	"fmt"
	"math"
)

const max = 10000

func main() {
	propSums := make([]int, max)
	amicable := []int{}
	for i := 0; i < max; i++ {
		propSums[i] = calcPropSum(i)
	}
	for p := range propSums {
		for n := range propSums {
			if p == propSums[n] && n == propSums[p] && p != n {
				if !inSlice(p, amicable) {
					amicable = append(amicable, p)
				}
				if !inSlice(n, amicable) {
					amicable = append(amicable, n)
				}
				break
			}
		}
	}
	sum := 0
	for a := range amicable {
		sum += amicable[a]
	}
	fmt.Println("sum amicables:", sum)
}

func calcPropSum(n int) int {
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

func inSlice(n int, slice []int) bool {
	for i := range slice {
		if n == slice[i] {
			return true
		}
	}
	return false
}
