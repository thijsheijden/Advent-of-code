package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/thijsheijden/advent_of_code/cmd/reader"
	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

type vec2d struct {
	x int
	y int
}

func main() {
	lines := reader.ReadInput()
	part1(lines)
	part2(lines)
}

var heading, north, east, south, west int = 90, 0, 90, 180, 270
var northPos, eastPos int = 0, 0

func part1(lines []string) {
	defer timetrack.TimeTrack(time.Now())

	for _, l := range lines {
		instruction := l[0]
		value, _ := strconv.Atoi(trimFirstRune(l))

		if instruction == 'R' || instruction == 'L' {
			changeHeading(instruction, value)
		} else {
			move(instruction, value)
		}
	}

	manhattanDistance := math.Abs(float64(northPos)) + math.Abs(float64(eastPos))
	fmt.Printf("Part 1 output: %f\n", manhattanDistance)
	heading, north, east, south, west = 90, 0, 90, 180, 270
	northPos, eastPos = 0, 0
}

func changeHeading(rotate byte, amount int) {
	switch rotate {
	case 'R':
		heading += amount
	case 'L':
		heading -= amount
	}

	fixHeading()
}

func fixHeading() {
	if heading >= 360 {
		heading -= 360
	} else if heading < 0 {
		heading += 360
	}
}

func move(direction byte, value int) {
	if direction == 'F' {
		switch heading {
		case 0:
			northPos += value
		case 90:
			eastPos += value
		case 180:
			northPos -= value
		case 270:
			eastPos -= value
		}
	} else {
		switch direction {
		case 'N':
			northPos += value
		case 'E':
			eastPos += value
		case 'S':
			northPos -= value
		case 'W':
			eastPos -= value
		}
	}
}

func part2(lines []string) {
	defer timetrack.TimeTrack(time.Now())

	var waypoint vec2d = vec2d{10, 1}

	for _, l := range lines {
		instruction := l[0]
		value, _ := strconv.Atoi(trimFirstRune(l))

		if instruction == 'R' || instruction == 'L' {
			waypoint = changeWaypointPos(instruction, value, waypoint)
		} else {
			waypoint = moveToWaypoint(instruction, value, waypoint)
		}
	}

	manhattanDistance := math.Abs(float64(northPos)) + math.Abs(float64(eastPos))
	fmt.Printf("Part 2 output: %f\n", manhattanDistance)
}

func changeWaypointPos(instruction byte, value int, waypoint vec2d) vec2d {
	newWaypoint := waypoint
	switch instruction {
	case 'R':
		for i := 0; i < value/90; i++ {
			newWaypoint = vec2d{newWaypoint.y, -newWaypoint.x}
		}
	case 'L':
		for i := 0; i < value/90; i++ {
			newWaypoint = vec2d{-newWaypoint.y, newWaypoint.x}
		}
	}
	return newWaypoint
}

func moveToWaypoint(direction byte, value int, waypoint vec2d) vec2d {
	newWaypoint := waypoint
	if direction == 'F' {
		northPos += value * newWaypoint.y
		eastPos += value * newWaypoint.x
	} else {
		switch direction {
		case 'N':
			newWaypoint.y += value
		case 'E':
			newWaypoint.x += value
		case 'S':
			newWaypoint.y -= value
		case 'W':
			newWaypoint.x -= value
		}
	}
	return newWaypoint
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
