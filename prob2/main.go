package main

import "fmt"

func main() {
	var1 := 1
	var2 := 2
	new := 0
	total := 0
	n := 4000000
	for var1 <= n {
		fmt.Println(var1)
		if var1%2 == 0 {
			fmt.Print("adding\n")
			total += var1
		}
		new = var1 + var2
		var1 = var2
		var2 = new
	}
	fmt.Println(total)
}
