package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/reader"
	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
	"github.com/thijsheijden/advent_of_code/utils/maths"
)

func main() {
	lines := reader.ReadInput()
	earliestDepartureTime, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")

	var busIDs, timeAfterT []int
	for i, id := range busses {
		if id != "x" {
			intID, _ := strconv.Atoi(id)
			busIDs = append(busIDs, intID)
			timeAfterT = append(timeAfterT, i)
		}
	}

	part1(earliestDepartureTime, busIDs)
	part2(busIDs, timeAfterT)
}

func part1(earliestDepartureTime int, busIDs []int) {
	defer timetrack.TimeTrack(time.Now())

	var bestBusToTake int
	var bestDifference float64 = -1000

	for _, id := range busIDs {

		diff := getTimeDiff(float64(id), float64(earliestDepartureTime))

		if diff > bestDifference {
			bestBusToTake = id
			bestDifference = diff
		}
	}

	fmt.Printf("Part 1 output: %f", math.Abs(bestDifference*float64(bestBusToTake)))
	fmt.Println()
}

func part2(busIDs []int, timeAfterT []int) {
	defer timetrack.TimeTrack(time.Now())

	sum, dep := busIDs[0]+timeAfterT[0], busIDs[0]

	for j := 1; j < len(busIDs); j++ {
		for (sum+timeAfterT[j])%busIDs[j] != 0 {
			sum += dep
		}

		dep = (dep * busIDs[j]) / maths.Gcd(dep, busIDs[j])
	}

	fmt.Println(sum)
}

func getTimeDiff(id float64, t float64) float64 {
	d := float64(t) / id
	early := math.Floor(d)
	late := math.Ceil(d)
	early *= id
	late *= id

	earlyDiff := t - early
	lateDiff := t - late

	if earlyDiff <= 0 && earlyDiff < lateDiff {
		return earlyDiff
	} else if lateDiff <= 0 {
		return lateDiff
	}
	return 0
}
