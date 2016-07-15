package main

import "fmt"

const source = `759564174782183587102004824765190123750334880277730763679965042806167092414126568340807033414872334732371694295371446525439152975114701133287773177839681757917152381714914358502729486366046889536730731669874031046298272309709873933853600423`
const rows = 15

var tree [][]int

func main() {
	readSource()
	fmt.Println(tree)
	solve()
	fmt.Println(tree)
}

func solve() {
	for r := rows - 2; r >= 0; r-- {
		for i := range tree[r] {
			tree[r][i] = tree[r][i] + max(tree[r+1][i], tree[r+1][i+1])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Everything below here is just for parsing the source string
// into the tree. Gotta learn go's string parsing functions
func readSource() {
	singlenums := make([]int, len(source))
	doublenums := make([]int, len(source)/2)
	byteslice := []rune(source)
	for i := 0; i < len(source); i++ {
		singlenums[i] = int(byteslice[i] - '0')
	}
	for i := 0; i < len(doublenums); i++ {
		doublenums[i] = 10*singlenums[i*2] + singlenums[i*2+1]
	}
	makeTree(doublenums)
}

func makeTree(nums []int) {
	tree = make([][]int, rows)
	w := 1
	count := 0
	for r := 0; r < rows; r++ {
		tree[r] = make([]int, w)
		for i := 0; i < w; i++ {
			tree[r][i] = nums[count]
			count++
		}
		w++
	}
}
