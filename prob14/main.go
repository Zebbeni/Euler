package main

import "fmt"

func main() {
	max, winner := 0, 0
	for i := 1; i < 1000000; i++ {
		terms := collatz(i)
		if terms > max {
			max = terms
			winner = i
		}
	}
	fmt.Println("max:", max, "winner:", winner)
}

func collatz(i int) int {
	n, count := i, 1
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	return count
}
