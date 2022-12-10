package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var stackList [10]Stack[string]

	initializeStacks(&stackList)

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		movePieces(fileScanner.Text(), &stackList)
	}
	readFile.Close()

	fmt.Println(stackList)
}

func movePieces(s string, stackList *[10]Stack[string]) {
	data := strings.Split(s, " ")
	pieces, _ := strconv.Atoi(data[1])
	from, _ := strconv.Atoi(data[3])
	to, _ := strconv.Atoi(data[5])

	var piece []string

	for i := 0; i < pieces; i++ {
		piece = append(piece, stackList[from].Pop())
	}

	for i := len(piece); i > 0; i-- {
		stackList[to].Push(piece[i-1])
	}
}

func initializeStacks(stackList *[10]Stack[string]) {
	stackList[1].List = []string{"V", "C", "D", "R", "Z", "G", "B", "W"}
	stackList[2].List = []string{"G", "W", "F", "C", "B", "S", "T", "V"}
	stackList[3].List = []string{"C", "B", "S", "N", "W"}
	stackList[4].List = []string{"Q", "G", "M", "N", "J", "V", "C", "P"}
	stackList[5].List = []string{"T", "S", "L", "F", "D", "H", "B"}
	stackList[6].List = []string{"J", "V", "T", "W", "M", "N"}
	stackList[7].List = []string{"P", "F", "L", "C", "S", "T", "G"}
	stackList[8].List = []string{"B", "D", "Z"}
	stackList[9].List = []string{"D", "N", "Z", "W"}

}

type Stack[T any] struct {
	List []T
}

func (stack *Stack[T]) Peek() T {
	if stack.Empty() {
		panic("Stack is empty")
	}
	return stack.List[len(stack.List)-1]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (stack *Stack[T]) Push(elm T) {
	stack.List = append(stack.List, elm)
}

func (stack *Stack[T]) Pop() (elm T) {
	elm, stack.List = stack.Peek(), stack.List[0:len(stack.List)-1]
	return elm
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.List) == 0
}
