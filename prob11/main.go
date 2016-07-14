package main

import (
	"fmt"
	"math"
)

const source = `08022297381500400075040507785212507791084949994017811857608717409843694804566200814931735579142993714067538830034913366552709523046011426924685601325671370236912231167151676389419236542240402866331380244732609903450244753353783684203517125032988128642367102638406759547066183864706726206802621220956394396308409166499421245558056673992697177878968314883489637221362309750076442045351400613397343133957817532822753167159403800462161409535692163905429635314755588824001754243629855786560048357189070544443744602158515417581980816805944769287392138652177704895540045208839735991607975732162626793327986688366887576220720346336746551232639353690442167338253911249472180846293240627636206936417230238834629969826759857404361620733529783190017431497148868116235705540170547183515469169233486143520189196748`

const dim = 20
const goal = 4

var nums []uint8

func main() {
	max := 0
	nums = parseToIntSlice(source)
	for i := 0; i < len(nums); i++ {
		max = maxVal(max, getProd(i, 1, 0))  // 0
		max = maxVal(max, getProd(i, 1, -1)) // 45
		max = maxVal(max, getProd(i, 0, 1))  // 90
		max = maxVal(max, getProd(i, 1, 1))  // 315
	}
	fmt.Println("max:", max)
}

func getProd(n, xoff, yoff int) int {
	prod := int(nums[n])
	i := n
	if isOutsideTable(n, xoff, yoff) {
		return 0
	}
	for count := 0; count < goal-1; count++ {
		i += xoff
		i += (yoff * dim)
		prod *= int(nums[i])
	}
	return prod
}

func isOutsideTable(n, xoff, yoff int) bool {
	endCol := getCol(n + xoff*(goal-1))
	endRow := getRow(n + yoff*(goal-1)*dim)
	if endCol >= dim || endCol < 0 {
		return true
	}
	if math.Abs(float64(endCol-getCol(n))) >= goal {
		return true
	}
	if endRow >= dim || endRow < 0 {
		return true
	}
	if math.Abs(float64(endRow-getRow(n))) >= goal {
		return true
	}
	return false
}

func getRow(i int) int {
	if i < 0 {
		return -1
	}
	return int(i / dim)
}

func getCol(i int) int {
	return (i % dim)
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func parseToIntSlice(s string) []uint8 {
	singlenums := make([]uint8, len(s))
	doublenums := make([]uint8, len(s)/2)
	byteslice := []rune(s)
	for i := 0; i < len(s); i++ {
		singlenums[i] = uint8(byteslice[i] - '0')
	}
	// this is probably dumb, I'm sure there's a way to convert every 2 bytes to a single integer
	for i := 0; i < len(doublenums); i++ {
		doublenums[i] = 10*singlenums[i*2] + singlenums[i*2+1]
	}
	return doublenums
}
