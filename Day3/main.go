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
	for fileScanner.Scan() {
		line := fileScanner.Text()
		firstSlice := line[0 : len(line)/2]
		secondSlice := line[len(line)/2 : len(line)]
		repeatedValue = append(repeatedValue, findDuplicates(firstSlice, secondSlice))
	}
	readFile.Close()

	for _, val := range repeatedValue {
		fmt.Println(byte(val))
		acumulated += convertedValue(byte(val))
	}

	fmt.Printf("The final result: %d\n", acumulated)
}

func findDuplicates(firstSlice, secondSlice string) byte {
	chars1 := make(map[byte]int)
	chars2 := make(map[byte]int)

	for _, letter := range firstSlice {
		chars1[byte(letter)]++
	}

	for _, letter := range secondSlice {
		chars2[byte(letter)]++
	}

	for k := range chars1 {
		if _, ok := chars2[k]; ok {
			// fmt.Printf("From %s and %s is repeated %s\n\n", firstSlice, secondSlice, string(k))
			return k
		}
	}

	return 0
}
