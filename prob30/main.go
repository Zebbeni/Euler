package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i <= 999999; i++ {
		if i == fifthPowerSum(i) {
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}

func fifthPowerSum(i int) int {
	sum, mult, digit := 0, 10, 0
	for val := i; val > 0; {
		digit = val % mult
		sum += pow(digit, 5)
		val = int(val / mult)
	}
	return sum
}

func pow(n, exp int) int {
	prod := 1
	for i := 0; i < exp; i++ {
		prod *= n
	}
	return prod
}
