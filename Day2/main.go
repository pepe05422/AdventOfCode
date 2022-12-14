package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type turn struct {
	hisSelection string
	mySelection  string
}

var lose = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var draw = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var win = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

func result(a, b string) int {

	selectionValue := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var duel int
	switch b {
	case "X":
		// Lose
		b = lose[a]
		break
	case "Y":
		// Draw
		b = draw[a]
		break
	case "Z":
		// Win
		b = win[a]
		break
	}

	switch (turn{hisSelection: a, mySelection: b}) {
	case turn{hisSelection: "A", mySelection: "X"},
		turn{hisSelection: "B", mySelection: "Y"},
		turn{hisSelection: "C", mySelection: "Z"}:
		duel = 3
		break

	case turn{hisSelection: "B", mySelection: "Z"},
		turn{hisSelection: "C", mySelection: "X"},
		turn{hisSelection: "A", mySelection: "Y"}:
		duel = 6
		break

	case turn{hisSelection: "C", mySelection: "Y"},
		turn{hisSelection: "A", mySelection: "Z"},
		turn{hisSelection: "B", mySelection: "X"}:
		duel = 0
		break
	}

	return duel + selectionValue[b]
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
		hands := strings.Split(fileScanner.Text(), " ")
		acumulative += result(hands[0], hands[1])
	}

	readFile.Close()

	fmt.Printf("The final result is: %d\n", acumulative)
}
