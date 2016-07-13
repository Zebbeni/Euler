package main

import "fmt"

func main() {
	count, n, x := 0, 10001, 1
	for count < n {
		x++
		if x == chop(x) {
			fmt.Println(x)
			count++
		}
	}
}

func chop(num int) int {
	a := 2
	for a < num {
		if num%a == 0 {
			return a
		}
		a += 1
	}
	return num
}
