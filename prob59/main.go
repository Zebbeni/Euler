package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAX_ASCII = 256
const MIN_VAL = 32
const MAX_VAL = 256

func main() {
	values := readCipherFile("cipher.txt")

	//fmt.Println("A ->", charToInt('A'))
	//fmt.Println("* ->", charToInt('*'))
	//fmt.Println("107 ->", intToChar(107))
	//fmt.Println("A XOR * ->", intToChar(charToInt('A')^charToInt('*')))

	for a := 0; a <= MAX_ASCII; a++ {
		fmt.Println("a:", a)
		for b := 0; b <= MAX_ASCII; b++ {
			for c := 0; c <= MAX_ASCII; c++ {
				decipher(values, []int{a, b, c})
			}
		}
	}
	// okay we've got the basics down. Now we need to:
	// define and iterate through 'candidate' encryption trios, decrypting the message with each candidate
	// evaluate each decrypted message programmatically, searching for english words (a, an, the, ' ')
}

func decipher(values, cipher []int) {
	charVals := make([]rune, len(values))
	sum := 0
	for i, val := range values {
		xorVal := val ^ cipher[i%3]
		if xorVal < MIN_VAL || xorVal > MAX_VAL {
			return
		}
		charVals[i] = intToChar(xorVal)
		sum += xorVal
	}

	message := string(charVals)
	if strings.Contains(message, "a ") && strings.Contains(message, "an ") && strings.Contains(message, "the ") {
		fmt.Printf("\n\nCandidate cipher: %d, %d, %d:\nMessage: %s\n\nSum: %d\n", cipher[0], cipher[1], cipher[2], message, sum)
	}
}

func intToChar(c int) rune {
	return ([]rune(string(c)))[0]
}

func charToInt(c rune) int {
	return int(c)
}

func readCipherFile(filename string) []int {
	var intStrings []string
	in, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("there was an error reading the file", err)
		return nil
	}
	r := csv.NewReader(strings.NewReader(string(in)))
	intStrings, err = r.Read()
	if err == io.EOF {
		return stringsToInts(intStrings)
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return stringsToInts(intStrings)
}

func stringsToInts(intStrings []string) []int {
	ints := make([]int, len(intStrings))
	var err error
	for i, s := range intStrings {
		if ints[i], err = strconv.Atoi(s); err != nil {
			panic(err)
		}
	}
	return ints
}
