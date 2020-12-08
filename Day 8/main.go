package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

type operation struct {
	lineNumber int
	operation  string
	value      int
	executed   bool
}

var operations []operation

func main() {
	defer timetrack.TimeTrack(time.Now())
	loadOperations()
	findBadOperation()
}

func loadOperations() {
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
}

func executeProgram() (bool, int) {
	reachedPreviouslyExecutedOp := 0
	accumulator := 0
	for i := 0; i < len(operations); i++ {
		if operations[i].executed {
			reachedPreviouslyExecutedOp++
			if reachedPreviouslyExecutedOp > 300 {
				return false, accumulator
			}
		}
		operations[i].executed = true
		switch operations[i].operation {
		case "acc":
			accumulator += operations[i].value
		case "jmp":
			i += operations[i].value - 1
		}
	}
	return true, accumulator
}

// Find the bad operation by brute forcing
func findBadOperation() {
	highestAccumulatedValue := 0
	for i := 0; i < len(operations); i++ {
		switch operations[i].operation {
		case "jmp":
			// Try changing this value to nop and see if the program completes
			operations[i].operation = "nop"
			executed, accumulatedValue := executeProgram()
			if executed {
				if accumulatedValue > highestAccumulatedValue {
					highestAccumulatedValue = accumulatedValue
				}
			} else {
				// Change back operation
				operations[i].operation = "jmp"
			}
		}
	}
	println(highestAccumulatedValue)
}
