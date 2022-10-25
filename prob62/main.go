package main

import (
	"fmt"
)

// a map of stringified digit counts to their corresponding cube permutations
var cubes map[string][]int
var found bool

func main() {
	found = false
	cubes = make(map[string][]int)
	for i := 1; !found; i++ {
		x := i * i * i
		checkCube(x)
	}
}

func checkCube(cube int) {
	// count digit occurrence
	digitCounts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 10; i <= cube*10; i *= 10 {
		digit := ((cube % i) - (cube % (i / 10))) / (i / 10)
		digitCounts[digit]++
	}

	// stringify the count slice for use as a key to lookup all permutations
	permsKey := fmt.Sprintf("%v", digitCounts)
	others, ok := cubes[permsKey]
	if !ok {
		cubes[permsKey] = []int{cube}
		return
	}

	if len(others) == 4 {
		fmt.Printf("\n%d ... %s -> %d others: %v", cube, permsKey, len(others), others)
		found = true
	}

	cubes[permsKey] = append(cubes[permsKey], cube)
}
