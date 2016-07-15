package main

import "fmt"

var ones = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var teens = [...]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = [...]string{"", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func main() {
	string := ""
	for i := 1; i <= 1000; i++ {
		string += buildString(i)
	}
	fmt.Println(string)
	fmt.Println("length:", len(string))
}

func buildString(value int) string {
	isBig := false
	s := ""
	n := value
	if n >= 1000 {
		s += "onethousand"
		n -= 1000
		isBig = true
	}
	if n >= 100 {
		num := int(n / 100)
		s += ones[num-1] + "hundred"
		n -= num * 100
		isBig = true
	}
	if n > 0 && isBig {
		s += "and"
	}
	if n >= 20 {
		num := int(n / 10)
		n -= num * 10
		s += tens[num-1]
	} else if n >= 10 {
		s += teens[n-10]
		return s
	}
	if n > 0 {
		s += ones[int(n-1)]
	}
	return s
}
