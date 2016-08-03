package main

import "fmt"

func main() {
	num, den := 1, 1
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++ {
			for c := 1; c < 10; c++ {
				if c != b && c != a && isSpecial(a, b, c) {
					num *= (a*10 + c)
					den *= (c*10 + b)
				}
			}
		}
	}
	fmt.Println("num:", num)
	fmt.Println("den:", den)
} // doesn't actually print the answer, but it's very easy to get here

func isSpecial(a, b, c int) bool {
	num := (a*10 + c)
	den := (c*10 + b)
	div := float32(num) / float32(den)
	if div < 1.0 && float32(a)/float32(b) == div {
		return true
	}
	return false
}
