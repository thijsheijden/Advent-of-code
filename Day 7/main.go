package main

import (
	"bufio"
	"os"
	"strings"
)

var bags map[string][]string
var bagsAlreadyFound map[string]int
var bagsThatCanContainGold int

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
			values[i] = words[1] + words[2]
		}

		// Add the bags to the map
		bags[words[0]+words[1]] = values
	}

	bagsAlreadyFound = make(map[string]int)

	// Go over all bags in the dictionary
	for k, v := range bags {
		if containsGold(k, v) {
			bagsThatCanContainGold++
		}
	}

	println(bagsThatCanContainGold)
}

func containsGold(key string, values []string) bool {
	// Iterate over the bags in the values and see if they contain 'shiny gold'
	for _, bag := range values {
		if bag == "shinygold" || containsGold(bag, bags[bag]) {
			return true
		}
	}
	return false
}
