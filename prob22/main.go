package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var names []string

func main() {
	names = readNamesFile("names.txt")
	sortNames(0, len(names)-1)
	sum := 0
	for i := range names {
		sum += score(i)
	}
	fmt.Println("sum:", sum)
}

func score(i int) int {
	sum := 0
	for n := range names[i] {
		sum += indexOf(names[i][n]) + 1
	}
	return sum * (i + 1)
}

//quicksort
func sortNames(lo, hi int) {
	if lo < hi {
		p := partitionNames(lo, hi)
		sortNames(lo, p-1)
		sortNames(p+1, hi)
	}
}

func partitionNames(lo, hi int) int {
	pivot, temp := names[hi], ""
	i := lo
	for j := lo; j < hi; j++ {
		if comesFirst(names[j], pivot) {
			temp = names[i]
			names[i] = names[j]
			names[j] = temp
			i++
		}
	}
	temp = names[i]
	names[i] = names[hi]
	names[hi] = temp
	return i
}

func indexOf(char byte) int {
	for a := range alphabet {
		if alphabet[a] == char {
			return a
		}
	}
	return -1
}

func comesFirst(name1, name2 string) bool {
	for n := range name1 {
		if n >= len(name2) {
			return false
		}
		name1Index := indexOf(name1[n])
		name2Index := indexOf(name2[n])
		if name1Index < name2Index {
			return true
		} else if name1Index > name2Index {
			return false
		}
	}
	return true
}

func readNamesFile(filename string) []string {
	var names []string
	in, err := (ioutil.ReadFile(filename))
	if err != nil {
		fmt.Println("there was an error reading the file")
		return names
	}
	r := csv.NewReader(strings.NewReader(string(in)))
	names, err = r.Read()
	if err == io.EOF {
		return names
	}
	if err != nil {
		log.Fatal(err)
		return names
	}
	return names
}
