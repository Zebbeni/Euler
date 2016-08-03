package main

import "fmt"

func main() {
	sum := 0
	for i := 3; i < 99999; i++ {
		if isSumOfFactorials(i) {
			sum += i
			fmt.Println("sum:", sum)
		}
	}
	fmt.Println(sum)
}

func isSumOfFactorials(n int) bool {
	sum := 0
	for val := n; val > 0; {
		sum += factorial(val % 10)
		if sum > n {
			return false
		}
		val = int(val / 10)
	}
	return sum == n
}

func factorial(n int) int {
	fact := 1
	for i := n; i > 0; i-- {
		fact *= i
	}
	return fact
}
