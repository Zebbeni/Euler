package main

import "fmt"

var primes = []int{2, 3, 5, 7, 11, 13, 17}
var sum = 0

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(digits); i++ {
		if digits[i] != 0 {
			dTemp := make([]int, len(digits))
			copy(dTemp, digits)
			chosen := []int{dTemp[i]}
			choice := append(dTemp[:i], dTemp[i+1:]...)
			findCombos(chosen, choice)
		}
	}
	fmt.Println("sum:", sum)
}

func findCombos(chosen, choice []int) {
	if len(choice) == 1 {
		val := append(chosen, choice...)
		if isPrimeWalkable(val) {
			num, mult := 0, 1
			for i := len(val) - 1; i >= 0; i-- {
				num += val[i] * mult
				mult *= 10
			}
			sum += num
		}
	} else {
		for i := 0; i < len(choice); i++ {
			choiceTemp := make([]int, len(choice))
			copy(choiceTemp, choice)
			newChosen := append(chosen, choiceTemp[i])
			newChoice := append(choiceTemp[:i], choiceTemp[i+1:]...)
			findCombos(newChosen, newChoice)
		}
	}
}

func isPrimeWalkable(intList []int) bool {
	var val, mult int
	primeIndex := 0
	for i := 1; i < len(intList)-2; i++ {
		val, mult = 0, 100
		for j := i; j < i+3; j++ {
			val += intList[j] * mult
			mult /= 10
		}
		if val%primes[primeIndex] != 0 {
			return false
		}
		primeIndex++
	}
	return true
}
