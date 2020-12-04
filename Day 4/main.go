package main

import (
	"bufio"
	"os"
	"regexp"
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
				// Validate the value for this key
				kv := strings.Split(s, ":")
				if validateValue(kv[0], kv[1]) {
					elementsInThisPassport.Add(kv[0])
				}
			}
		}
	}

	elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
	if elementsMissingFromPassport.IsEmpty() {
		numberOfValidPassports++
	}

	println(numberOfValidPassports)
}

func validateValue(key string, value string) bool {
	switch key {
	case "byr":
		regex := regexp.MustCompile(`(19[2-8][0-9]|199[0-9]|200[0-2])`)
		return regex.Match([]byte(value))
	case "iyr":
		regex := regexp.MustCompile(`(201[0-9]|2020)`)
		return regex.Match([]byte(value))
	case "eyr":
		regex := regexp.MustCompile(`(202[0-9]|2030)`)
		return regex.Match([]byte(value))
	case "hgt":
		regex := regexp.MustCompile(`(1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in`)
		return regex.Match([]byte(value))
	case "hcl":
		regex := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
		return regex.Match([]byte(value))
	case "ecl":
		regex := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
		return regex.Match([]byte(value))
	case "pid":
		regex := regexp.MustCompile(`[0-9]{9}`)
		return regex.Match([]byte(value))
	case "cid":
		return true
	}
	return false
}
