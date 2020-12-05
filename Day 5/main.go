package main

import (
	"bufio"
	"os"
)

var highestSeatId int = 0

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

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}
	println(highestSeatId)
}
