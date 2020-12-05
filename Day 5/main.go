package main

import (
	"bufio"
	"os"
)

var currentlyOccupiedSeats [1024]int

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		minRow := 0
		maxRow := 127
		minColumn := 0
		maxColumn := 7

		for _, c := range line {
			if c == 'F' {
				maxRow = maxRow - ((maxRow - minRow) / 2) - 1
			} else if c == 'B' {
				minRow = minRow + (maxRow-minRow)/2 + 1
			} else if c == 'R' {
				minColumn = minColumn + (maxColumn-minColumn)/2 + 1
			} else {
				maxColumn = maxColumn - ((maxColumn - minColumn) / 2) - 1
			}
		}

		seatId := minRow*8 + minColumn
		currentlyOccupiedSeats[seatId] = 1
	}
	for i, v := range currentlyOccupiedSeats {
		if i > 0 && i < len(currentlyOccupiedSeats) {
			if v == 0 && currentlyOccupiedSeats[i-1] == 1 && currentlyOccupiedSeats[i+1] == 1 {
				println(i)
			}
		}
	}
}
