package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := big.NewInt(1)
	j := big.NewInt(1)
	temp := big.NewInt(0)
	for c := 3; ; c++ {
		temp.Add(i, j)
		i.Set(j)
		j.Set(temp)
		if len(j.Text(10)) >= 1000 {
			fmt.Println("c:", c)
			break
		}
	}
}
