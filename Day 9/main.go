package main

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/thijsheijden/advent_of_code/cmd/timetrack"
)

var preamble []int
var allValues []int
var preambleSize int = 25
var min, max int = 2147483647, 0

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
			allValues = append(allValues, val)
		} else {
			// This number should be the sum of two numbers in the preamble
			val, _ := strconv.Atoi(line)
			allValues = append(allValues, val)
			if !sumPreamble(val) {
				encryptionWeakness := findContiguosRange(val)
				println(encryptionWeakness)
				return
			}
			// Remove the first value from the preamble and insert this value
			preamble = preamble[1:]
			preamble = append(preamble, val)
		}
		i++
	}
}

// Find the contiguous range which sums to the invalid number and return the sum of lowest+highest
func findContiguosRange(v int) int {
	sum := 0
	for i := 0; i < len(allValues); i++ {
		sum = allValues[i]
		setMinMax(allValues[i])
		for j := i + 1; j < len(allValues); j++ {
			sum += allValues[j]
			setMinMax(allValues[j])
			if sum == v {
				return min + max
			} else if sum > v {
				break
			}
		}
		min, max = 2147483647, 0
	}
	return 0
}

func setMinMax(i int) {
	if i < min {
		min = i
	}
	if i > max {
		max = i
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
