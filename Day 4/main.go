package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/set"
)

var numberOfValidPassports int = 0

func main() {

	requiredElements := set.New(set.NonThreadSafe)
	requiredElements.Add("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")

	elementsInThisPassport := set.New(set.NonThreadSafe)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		// Check if the line is empty, if so check if the passport is valid
		if len(line) == 0 {
			elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
			if elementsMissingFromPassport.IsEmpty() {
				numberOfValidPassports++
			}

			// Empty the set
			elementsInThisPassport.Clear()
		}

		// Add keys to the set
		splitOnSpace := strings.Split(line, " ")
		for _, s := range splitOnSpace {
			if len(s) > 1 {
				elementsInThisPassport.Add(strings.Split(s, ":")[0])
			}
		}
	}

	elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
	if elementsMissingFromPassport.IsEmpty() {
		numberOfValidPassports++
	}

	println(numberOfValidPassports)
}
