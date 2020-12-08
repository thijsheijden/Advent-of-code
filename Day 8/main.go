package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	lineNumber int
	operation  string
	value      int
	executed   bool
}

var operations []operation
var accumulator int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		operationAndValue := strings.Split(line, " ")
		val, _ := strconv.Atoi(operationAndValue[1])

		operations = append(operations, operation{lineNumber: i,
			operation: operationAndValue[0],
			value:     val,
			executed:  false})

		i++
	}

	for i := 0; i < len(operations); i++ {
		if operations[i].executed {
			println(accumulator)
			return
		}
		operations[i].executed = true
		switch operations[i].operation {
		case "acc":
			accumulator += operations[i].value
		case "jmp":
			i += operations[i].value - 1
		}
	}
}
