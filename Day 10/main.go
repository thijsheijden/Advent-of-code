package main

import (
	"sort"
	"strconv"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/reader"
	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

func main() {
	defer timetrack.TimeTrack(time.Now())
	lines := reader.ReadInput()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var jolts []int
	for _, x := range lines {
		i, _ := strconv.Atoi(x)
		jolts = append(jolts, i)
	}

	// Sort the list
	sort.Slice(jolts, func(i, j int) bool {
		return jolts[i] < jolts[j]
	})

	var oneJoltDif, twoJoltDif, threeJoltDif, currentJoltage int = 0, 0, 1, 0

	for i := 0; i < len(jolts); i++ {
		diff := jolts[i] - currentJoltage
		switch diff {
		case 1:
			oneJoltDif++
			currentJoltage++
		case 2:
			twoJoltDif++
			currentJoltage += 2
		case 3:
			threeJoltDif++
			currentJoltage += 3
		}
	}

	println(oneJoltDif * threeJoltDif)
}

func part2(lines []string) {

}
