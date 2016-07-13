package main

import "fmt"

func main() {
	n := 100
	sum := (n + 1) * n / 2
	sqsum := pow(sum, 2)
	sumsq := 0
	for i := 1; i <= n; i++ {
		sumsq += pow(i, 2)
	}
	fmt.Println(sqsum - sumsq)
}

func pow(n, x int) int {
	prod := 1
	for i := 0; i < x; i++ {
		prod *= n
	}
	return prod
}
