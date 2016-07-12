package main

import "fmt"

func main() {
	var1 := 1
	var2 := 2
	new, total := 0, 0
	n := 4000000
	for var1 <= n {
		if var1%2 == 0 {
			total += var1
		}
		new = var1 + var2
		var1 = var2
		var2 = new
	}
	fmt.Println(total)
}
