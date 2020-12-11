package main

import (
	"fmt"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/reader"
	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

func main() {
	lines := reader.ReadInput()
	part1(lines)
}

func part1(l []string) {
	defer timetrack.TimeTrack(time.Now())

	var layout [][]byte

	for _, s := range l {
		bytes := []byte(s)
		layout = append(layout, bytes)
	}

	seatsChanged := true
	for seatsChanged {
		change, newLayout := genNextLayout(layout)
		seatsChanged = change
		layout = newLayout
	}
	println(numberOfTakenSeats)
}

var numberOfTakenSeats int = 0

func genNextLayout(layout [][]byte) (bool, [][]byte) {
	newLayout := make([][]byte, len(layout))
	for i := range newLayout {
		newLayout[i] = make([]byte, len(layout[i]))
	}

	seatsChanged := false
	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			updated, newSeat := updateSeat(x, y, layout[y][x], layout)
			if updated {
				if newSeat == '#' {
					numberOfTakenSeats++
				} else {
					numberOfTakenSeats--
				}
				seatsChanged = true
			}
			newLayout[y][x] = newSeat
		}
	}
	return seatsChanged, newLayout
}

func updateSeat(x int, y int, seat byte, layout [][]byte) (bool, byte) {

	if seat == '.' {
		return false, '.'
	}

	numOccupied := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// Not out of bounds
			if !(x+dx < 0 || x+dx >= len(layout[0])) && !(y+dy < 0 || y+dy >= len(layout)) && !(dx == 0 && dy == 0) {
				if layout[y+dy][x+dx] == '#' {
					numOccupied++
				}
			}
		}
	}

	if seat == '#' {
		if numOccupied >= 4 {
			return true, 'L'
		}
	}

	if seat == 'L' {
		if numOccupied == 0 {
			return true, '#'
		}
	}

	return false, seat
}

func printLayout(layout [][]byte) {
	for _, r := range layout {
		fmt.Printf("%s", r)
		fmt.Println()
	}
	fmt.Println()
}
