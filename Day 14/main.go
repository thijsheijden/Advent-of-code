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
	part1(lines)
}

func part1(lines []string) {
	defer timetrack.TimeTrack(time.Now())

	var mask string
	var memory map[string]int = make(map[string]int)

	for _, line := range lines {
		l := strings.Split(line, " = ")
		if l[0] == "mask" {
			mask = l[1]
		} else {
			i, _ := strconv.Atoi(l[1])
			memory[l[0]] = applyMask(mask, []byte(intToBitString(int64(i))))
		}
	}

	// Sum up memory
	var res int = 0
	for _, v := range memory {
		res += v
	}

	fmt.Printf("Part 1 output: %d\n", res)
}

func part2() {

}

func applyMask(mask string, value []byte) int {
	for i, m := range mask {
		if m != 'X' {
			value[i] = byte(m)
		}
	}

	var r int = 0
	for i, b := range value {
		if b == '1' {
			r += int(math.Pow(2, float64(35-i)))
		}
	}

	return r
}

func intToBitString(i int64) string {
	str := strconv.FormatInt(i, 2)
	if len(str) < 36 {
		str = strings.Repeat("0", 36-len(str)) + str
	}
	return str
}
