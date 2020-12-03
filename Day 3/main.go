package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var treesEncountered int64 = 0
	var verticalDistance int = 0
	var horizontalDistance int = -3

	for scanner.Scan() {
		line := scanner.Text()

		// Move in the defined slope
		verticalDistance++
		horizontalDistance += 3

		// Check if the horizontal distance is larger than the line length
		if horizontalDistance >= len(line) {
			horizontalDistance = horizontalDistance % len(line)
		}

		println(line[horizontalDistance])

		// Determine if we encountered a tree
		if line[horizontalDistance] == '#' {
			treesEncountered++
		}
	}

	println(treesEncountered)
}
