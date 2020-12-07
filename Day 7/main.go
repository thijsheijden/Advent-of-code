package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var bags map[string][]string
var numberOfContainedBags int

func main() {
	bags = make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Remove dot at end of line
		line = line[:len(line)-1]

		// Split at 'contain'
		keyValue := strings.Split(line, "contain")

		// Split the right side on ','
		values := strings.Split(keyValue[1], ",")

		// Split the key into words to only grab the color
		words := strings.Fields(keyValue[0])

		// Split the value into words and only grab the color
		for i := 0; i < len(values); i++ {
			words := strings.Fields(values[i])
			values[i] = words[0] + " " + words[1] + words[2]
		}

		// Add the bags to the map
		bags[words[0]+words[1]] = values
	}

	// Go over all bags in the dictionary
	for _, bag := range bags["shinygold"] {
		// Determine the number of bags inside each of those bags
		numberOfContainedBags += containedBags(bag)
	}

	println(numberOfContainedBags)
}

func containedBags(key string) int {

	var bagsContainedInThisBage int = 0

	words := strings.Fields(key)

	i, _ := strconv.Atoi(words[0])

	bagsContainedInThisBage += i

	// Iterate over the bags in the values and see if they contain 'shiny gold'
	for _, bag := range bags[words[1]] {
		bagsContainedInThisBage += i * containedBags(bag)
	}

	return bagsContainedInThisBage
}
