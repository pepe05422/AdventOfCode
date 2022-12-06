package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type distribution struct {
	from int
	to   int
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var acumulative int
	for fileScanner.Scan() {
		pairs := strings.Split(fileScanner.Text(), ",")
		elf1, elf2 := stripeValues(pairs)
		if overlaps(elf1, elf2) {
			acumulative++
		}

	}
	readFile.Close()

	fmt.Printf("The final result is: %d\n", acumulative)
}

func fullyContains(elf1, elf2 distribution) bool {
	return elf1.from <= elf2.from && elf1.to >= elf2.to || elf2.from <= elf1.from && elf2.to >= elf1.to
}

func overlaps(elf1, elf2 distribution) bool {
	return !(elf1.to < elf2.from) && !(elf2.to < elf1.from)
}

func stripeValues(pairs []string) (distribution, distribution) {
	val1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
	val2, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])
	val3, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
	val4, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

	elf1 := distribution{
		from: val1,
		to:   val2,
	}

	elf2 := distribution{
		from: val3,
		to:   val4,
	}

	return elf1, elf2
}
