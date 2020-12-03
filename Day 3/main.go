package main

import (
	"bufio"
	"os"
)

type slope struct {
	horizontal       int
	vertical         int
	x                int
	y                int
	treesEncountered int64
}

var slopes []slope

func main() {

	var verticalDistance int = 0

	slopes = append(slopes, slope{horizontal: 1, vertical: 1, x: 0, y: 0, treesEncountered: 0},
		slope{horizontal: 3, vertical: 1, x: 0, y: 0, treesEncountered: 0},
		slope{horizontal: 5, vertical: 1, x: 0, y: 0, treesEncountered: 0},
		slope{horizontal: 7, vertical: 1, x: 0, y: 0, treesEncountered: 0},
		slope{horizontal: 1, vertical: 2, x: 0, y: 0, treesEncountered: 0})

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(slopes); i++ {
			slope := &slopes[i]

			if verticalDistance%slope.vertical == 0 {
				// Determine if we need to wrap our horizontal distance
				if slope.x >= len(line) {
					slope.x = slope.x % len(line)
				}

				// Determine if we encounter a tree
				if line[slope.x] == '#' {
					slope.treesEncountered++
				}

				// Update slope
				slope.x += slope.horizontal
				slope.y += slope.vertical
			}
		}

		// Update vertical distance by one
		verticalDistance++
	}

	println(slopes[0].treesEncountered)

	var totalEncounteredTrees int64 = 1

	// Multiply all encountered trees
	for _, slope := range slopes {
		totalEncounteredTrees *= slope.treesEncountered
	}

	println(totalEncounteredTrees)
}
