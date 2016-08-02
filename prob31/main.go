package main

import "fmt"

var values = []int{1, 2, 5, 10, 20, 50, 100, 200}

func main() {
	fmt.Println(waysToMake(200, len(values)-1))
}

func waysToMake(n, start int) int {
	numWays := 0
	for i := start; i >= 0; i-- {
		numWays += waysToMakeWith(n, i)
	}
	return numWays
}

func waysToMakeWith(n, i int) int {
	value := values[i]
	if value > n {
		return 0
	} else if value == n {
		return 1
	}
	return waysToMake(n-value, i)
}
