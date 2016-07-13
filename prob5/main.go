package main

import (
	"fmt"
	"math"
)

func main() {
	end, factor, prod := 20, 0, 1
	factormax := make([]int, end)

	for n := 2; n <= end; n++ {
		num := n
		factornum := make([]int, end)
		for num != 1 {
			factor = chop(num)
			factornum[factor-1]++
			num = num / factor
		}
		for i := 0; i < end; i++ {
			factormax[i] = max(factornum[i], factormax[i])
		}
	}
	for i := 0; i < end; i++ {
		prod *= pow(i+1, factormax[i])
	}
	fmt.Println(prod)
}

func chop(num int) int {
	a := 2
	for a < int(math.Sqrt(float64(num)))+1 {
		if num%a == 0 {
			return a
		}
		a++
	}
	return num
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func pow(n, exp int) int {
	prod := 1
	for i := 0; i < exp; i++ {
		prod *= n
	}
	return prod
}
