package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var maximum int
	var secondM int
	var thirdM int
	var prevMax int
	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		prevMax += number
		if strings.TrimRight(fileScanner.Text(), "\n") == "" {
			if prevMax > maximum {
				thirdM = secondM
				secondM = maximum
				maximum = prevMax
			} else if prevMax > secondM {
				thirdM = secondM
				secondM = prevMax
			} else if prevMax > thirdM {
				thirdM = prevMax
			}
			prevMax = 0
		}
	}
	readFile.Close()
	fmt.Println(maximum + secondM + thirdM)
}
