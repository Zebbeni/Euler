package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	words := readWordsFile("words.txt")
	count := 0
	for _, word := range words {
		if isTriangleWord(word) {
			count++
		}
	}
	fmt.Println("count:", count)
}

func isTriangleWord(word string) bool {
	val := 0
	for i := range word {
		val += valueOf(word[i])
	}
	return isTriangle(val)
}

func isTriangle(n int) bool {
	quadVal := 8*n + 1
	sqrt := int(math.Sqrt(float64(quadVal)))
	return sqrt*sqrt == quadVal && (sqrt-1)%2 == 0
}

func valueOf(char byte) int {
	for a := range alphabet {
		if alphabet[a] == char {
			return a + 1
		}
	}
	return -1
}

func readWordsFile(filename string) []string {
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
