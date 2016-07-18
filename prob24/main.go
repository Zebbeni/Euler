package main

import "fmt"

var digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var solution = make([]int, 10)
var sequence = 1000000

func main() {
	for d := 0; len(digits) > 0; d++ {
		k := factorial(len(digits)) / len(digits)
		i := (sequence - 1) / k
		sequence = sequence - (i * k)
		solution[d] = digits[i]
		digits = append(digits[:i], digits[i+1:]...)
	}
	fmt.Println("solution:", solution)
}

func factorial(n int) int {
	fact := 1
	for i := n; i > 0; i-- {
		fact *= i
	}
	return fact
}
