package main

import "fmt"

const n = 120

func main() {
	maxSolutions, maxN := 0, 0
	for n := 1; n <= 1000; n++ {
		solutions := 0
		for a := 1; a < n; a++ {
			for b := a; b < (n - a); b++ {
				c := n - a - b
				if a*a+b*b == c*c {
					solutions++
				}
			}
		}
		if solutions > maxSolutions {
			maxN = n
			maxSolutions = solutions
		}
	}
	fmt.Println("max solutions:", maxSolutions, "maxN:", maxN)
}
