package main

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

var preamble []int
var preambleSize int = 25

func main() {
	defer timetrack.TimeTrack(time.Now())
	decipherXMAS()
}

func decipherXMAS() {
	scanner := bufio.NewScanner(os.Stdin)

	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Building up the preamble
		if i < preambleSize {
			val, _ := strconv.Atoi(line)
			preamble = append(preamble, val)
		} else {
			// This number should be the sum of two numbers in the preamble
			val, _ := strconv.Atoi(line)
			if !sumPreamble(val) {
				println(val)
				return
			}
			// Remove the first value from the preamble and insert this value
			preamble = preamble[1:]
			preamble = append(preamble, val)
		}
		i++
	}
}

// go over the preamble and check if two numbers sum to the given number
func sumPreamble(v int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := 0; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == v && i != j {
				return true
			}
		}
	}
	return false
}
