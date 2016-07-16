package main

import "fmt"

var daysPerMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	startingSundays, dayOfWeek := 0, 1
	for year := 1900; year <= 2000; year++ {
		setFebruaryDays(year)
		for month := 0; month < 12; month++ {
			if dayOfWeek%7 == 0 && year >= 1901 {
				startingSundays++
			}
			dayOfWeek += daysPerMonth[month]
		}
	}
	fmt.Println("startingSundays:", startingSundays)
}

func setFebruaryDays(year int) {
	if year%100 == 0 && year%400 == 0 {
		daysPerMonth[1] = 29
	} else if year%100 != 0 && year%4 == 0 {
		daysPerMonth[1] = 29
	} else {
		daysPerMonth[1] = 28
	}
}
