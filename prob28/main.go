package main

import "fmt"

func main() {
	sum, last, toAdd := 1, 1, 0
	for i := 1; i <= 500; i++ {
		toAdd = i * 2
		for j := 0; j < 4; j++ {
			sum += (last + toAdd)
			last = last + toAdd
		}
	}
	fmt.Println("sum:", sum)
}
