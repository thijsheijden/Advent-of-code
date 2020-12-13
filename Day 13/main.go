package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/reader"
	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

func main() {
	lines := reader.ReadInput()
	earliestDepartureTime, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")

	var busIDs []int
	for _, id := range busses {
		if id != "x" {
			intID, _ := strconv.Atoi(id)
			busIDs = append(busIDs, intID)
		}
	}

	part1(earliestDepartureTime, busIDs)
}

func part1(earliestDepartureTime int, busIDs []int) {
	defer timetrack.TimeTrack(time.Now())

	var bestBusToTake int
	var bestDifference float64 = -1000

	for _, id := range busIDs {
		floatID := float64(id)
		d := float64(earliestDepartureTime) / floatID
		early := math.Floor(d)
		late := math.Ceil(d)
		early *= floatID
		late *= floatID

		earlyDiff := float64(earliestDepartureTime) - early
		lateDiff := float64(earliestDepartureTime) - late

		var bestDiff float64
		if earlyDiff <= 0 && earlyDiff < lateDiff {
			bestDiff = earlyDiff
		} else if lateDiff <= 0 {
			bestDiff = lateDiff
		}

		if bestDiff > bestDifference {
			bestBusToTake = id
			bestDifference = bestDiff
		}
	}

	fmt.Printf("Part 1 output: %f", math.Abs(bestDifference*float64(bestBusToTake)))
	fmt.Println()
}
