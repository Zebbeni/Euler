package main

import (
	"fmt"
	"math"
)

var max = 0

func main() {
	for digits := 2; ; digits++ {
		findAllCombos(0, 1, 0, digits)
	}
}

// ex.
// num:     504006 (number to add to)
// blanks:   10100 (blanks value to be multiplied and added during checks)
// mult:   1000000 (value to multiply new digit by -> x1)
// remaining     1 (digits left to find -> -1 )
func findAllCombos(num, mult, blanks, remaining int) {
	if remaining == 0 {
		testCombo(num, blanks)
	} else {
		chooseNumber(num, mult, blanks, remaining)
		findAllCombos(num, mult*10, blanks+1*mult, remaining-1)
	}
}

func chooseNumber(num, mult, blanks, remaining int) {
	for i := 0; i < 10; i++ {
		if mult != 1 || (i%2 != 0 && i != 5) {
			findAllCombos(num+mult*i, mult*10, blanks, remaining-1)
		}
	}
}

func testCombo(num, blanks int) {
	vals := []int{}
	if blanks != 0 {
		count := 0
		for i := 0; i < 10; i++ {
			if isPrime(num + blanks*i) {
				vals = append(vals, num+blanks*i)
				count++
			}
		}
		if count > max {
			max = count
			fmt.Println("max:", max, "num", num, "blanks", blanks, "vals:", vals)
		}
	}
}

func isPrime(n int) bool {
	return n > 1 && chop(n) == n
}

func chop(num int) int {
	for a := 2; a < int(math.Sqrt(float64(num)))+1; a++ {
		if num%a == 0 {
			return a
		}
	}
	return num
}
