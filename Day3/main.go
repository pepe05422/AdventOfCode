package main

import (
	"bufio"
	"fmt"
	"os"
)

func convertedValue(a byte) int {
	// Mayusculas
	if a >= 65 && a <= 90 {
		return 27 + (int(a) - 65)
	} else {
		return 1 + (int(a) - 97)
	}
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var repeatedValue []byte
	var acumulated int
	var slices []string
	var i = 1
	for fileScanner.Scan() {
		slices = append(slices, fileScanner.Text())
		if i%3 == 0 {
			// find repeated value
			fmt.Println(slices)
			repeatedValue = append(repeatedValue, findDuplicates(slices))
			slices = make([]string, 0)
		}
		i++

	}
	readFile.Close()

	for _, val := range repeatedValue {
		fmt.Println(byte(val))
		acumulated += convertedValue(byte(val))
	}

	fmt.Printf("The final result: %d\n", acumulated)
}

func findDuplicates(slices []string) byte {
	chars1 := make(map[byte]int)
	chars2 := make(map[byte]int)
	chars3 := make(map[byte]int)

	for _, letter := range slices[0] {
		chars1[byte(letter)]++
	}

	for _, letter := range slices[1] {
		chars2[byte(letter)]++
	}

	for _, letter := range slices[2] {
		chars3[byte(letter)]++
	}

	for k := range chars1 {
		if _, ok := chars2[k]; ok {
			// fmt.Printf("From %s and %s is repeated %s\n\n", firstSlice, secondSlice, string(k))
			if _, ok := chars3[k]; ok {
				return k
			}
		}
	}

	return 0
}
