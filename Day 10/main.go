package main

import (
	"sort"
	"strconv"
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

func part1(lines []string) {
	defer timetrack.TimeTrack(time.Now())
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

	print("Solution part 1: ")
	println(oneJoltDif * threeJoltDif)
}

func part2(lines []string) {
	defer timetrack.TimeTrack(time.Now())
	var jolts []int
	for _, x := range lines {
		i, _ := strconv.Atoi(x)
		jolts = append(jolts, i)
	}

	possible := make(map[int][]int)
	possible[0] = []int{1, 2, 3}

	for _, element := range jolts {
		possible[element] = []int{element + 3, element + 2, element + 1}
	}

	result := connections(possible, make(map[int]int), slice.Max(jolts)+3, 0)
	println("Solution part 2:", result)
}

func connections(possible map[int][]int, memo map[int]int, target int, currPos int) int {
	if value, seen := memo[currPos]; seen {
		return value
	}

	value := 0
	for _, current := range possible[currPos] {
		if current != target {
			value += connections(possible, memo, target, current)
			continue
		}

		value++
	}

	memo[currPos] = value
	return value
}
