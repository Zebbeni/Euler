package main

import "fmt"

func main() {
	maxDen, max, cycle := 1, 0, 1
	for i := 980; i < 990; i++ {
		cycle = calcCycle(i)
		if cycle > max {
			max = cycle
			maxDen = i
		}
	}
	fmt.Println("max:", maxDen, "cycle: ", max)
}

func calcCycle(den int) int {
	numMap := make(map[int]int)
	val, sub, count := 1, 1, 0
	for num := 10; ; count++ {
		if num == 0 {
			return 0
		}
		if lastSeen, ok := numMap[num]; ok {
			return count - lastSeen
		}
		numMap[num] = count
		val = int(num / den)
		sub = num - (val * den)
		num = sub * 10
	}
}
