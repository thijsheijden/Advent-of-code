package main

import (
	"fmt"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/reader"
	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
	"github.com/thijsheijden/advent_of_code/utils/slice"
)

func main() {
	lines := reader.ReadInput()
	part1(lines)
	part2(lines)
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
		change, newLayout := genNextLayout(layout, false)
		seatsChanged = change
		layout = newLayout
	}
	println(numberOfTakenSeats)
}

var numberOfTakenSeats int = 0

func genNextLayout(layout [][]byte, adjacent bool) (bool, [][]byte) {

	newLayout := make([][]byte, len(layout))
	for i := range newLayout {
		newLayout[i] = make([]byte, len(layout[i]))
	}

	seatsChanged := false
	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			if !adjacent {
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
			} else {
				updated, newSeat := updateSeatAdjacent(x, y, layout[y][x], layout)
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

func part2(l []string) {
	defer timetrack.TimeTrack(time.Now())

	var layout [][]byte

	for _, s := range l {
		bytes := []byte(s)
		layout = append(layout, bytes)
	}

	numberOfTakenSeats = 0
	seatsChanged := true
	for seatsChanged {
		change, newLayout := genNextLayout(layout, true)
		seatsChanged = change
		layout = newLayout
	}
	println(numberOfTakenSeats)
}

func updateSeatAdjacent(x int, y int, seat byte, layout [][]byte) (bool, byte) {

	if seat == '.' {
		return false, '.'
	}

	numOccupied := 0

	vertical := slice.ColumnToSlice(x, layout)
	horizontal := layout[y]

	// Vertical chairs above
	closest := -len(vertical)
	var closestSeat byte = '.'
	for dy := -1; y+dy >= 0; dy-- {
		if dy >= closest {
			if vertical[y+dy] != '.' {
				closestSeat = vertical[y+dy]
				closest = dy
			}
		}
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Vertical chairs below
	closest = len(vertical)
	closestSeat = '.'
	for dy := 1; y+dy < len(vertical); dy++ {
		if dy <= closest {
			if vertical[y+dy] != '.' {
				closestSeat = vertical[y+dy]
				closest = dy
			}
		}
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Horizontal chairs left
	closest = -len(horizontal)
	closestSeat = '.'
	for dx := -1; x+dx >= 0; dx-- {
		if dx >= closest {
			if horizontal[x+dx] != '.' {
				closestSeat = horizontal[x+dx]
				closest = dx
			}
		}
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Horizontal chairs right
	closest = len(horizontal)
	closestSeat = '.'
	for dx := 1; x+dx < len(horizontal); dx++ {
		if dx <= closest {
			if horizontal[x+dx] != '.' {
				closestSeat = horizontal[x+dx]
				closest = dx
			}
		}
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Vertical Chairs (x, y -> top left)
	closest = -len(horizontal)
	closestSeat = '.'
	var okay bool
	dx := -1
	dy := -1
	okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	for okay {
		if layout[y+dy][x+dx] != '.' {
			if dx >= closest {
				closest = dx
				closestSeat = layout[y+dy][x+dx]
			}
		}
		dy--
		dx--
		okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Vertical Chairs (x, y -> bottom right)
	closest = len(horizontal)
	closestSeat = '.'
	okay = false
	dx = 1
	dy = 1
	okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	for okay {
		if layout[y+dy][x+dx] != '.' {
			if dx <= closest {
				closest = dx
				closestSeat = layout[y+dy][x+dx]
			}
		}
		dy++
		dx++
		okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Vertical Chairs (x, y -> bottom left)
	closest = -len(horizontal)
	closestSeat = '.'
	okay = false
	dx = -1
	dy = 1
	okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	for okay {
		if layout[y+dy][x+dx] != '.' {
			if dx >= closest {
				closest = dx
				closestSeat = layout[y+dy][x+dx]
			}
		}
		dy++
		dx--
		okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	}

	if closestSeat == '#' {
		numOccupied++
	}

	// Vertical Chairs (x, y -> top right)
	closest = len(horizontal)
	closestSeat = '.'
	okay = false
	dx = 1
	dy = -1
	okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	for okay {
		if layout[y+dy][x+dx] != '.' {
			if dx <= closest {
				closest = dx
				closestSeat = layout[y+dy][x+dx]
			}
		}
		dy--
		dx++
		okay = !(x+dx < 0 || x+dx >= len(horizontal)) && !(y+dy < 0 || y+dy >= len(vertical))
	}

	if closestSeat == '#' {
		numOccupied++
	}

	if seat == '#' {
		if numOccupied >= 5 {
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
