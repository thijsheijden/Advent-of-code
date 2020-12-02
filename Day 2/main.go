package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	letter        string
	minOccurrence int
	maxOccurrence int
}

type password struct {
	policy   passwordPolicy
	password string
}

var passwords []password

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Load in all the passwords
	for scanner.Scan() {
		line := scanner.Text()

		var occurrenceString, passwordString, policyLetterString string
		fmt.Sscanf(line, "%s%s%s", &occurrenceString, &policyLetterString, &passwordString)

		// Get the min and max occurrence of the policy letter
		minMax := strings.Split(occurrenceString, "-")
		min, err := strconv.Atoi(minMax[0])
		max, err2 := strconv.Atoi(minMax[1])

		if err != nil || err2 != nil {
			log.Fatalln(err)
			log.Fatalln(err2)
		}

		// Get the policy letter
		policyLetter := strings.Split(policyLetterString, ":")[0]

		// Create the password objects
		passPolicy := passwordPolicy{minOccurrence: min, maxOccurrence: max, letter: policyLetter}
		pass := password{policy: passPolicy, password: passwordString}

		// Add to the list of passwords
		passwords = append(passwords, pass)
	}

	// Go over all passwords and determine how many of them adhere to their password policy
	var adhereToPasswordPolicy int = 0
	for _, pass := range passwords {
		// Count how many times the letter appears in the password
		count := strings.Count(pass.password, pass.policy.letter)

		// Check if count is no less than minOccurrence and no more than maxOccurrence
		if count <= pass.policy.maxOccurrence && count >= pass.policy.minOccurrence {
			adhereToPasswordPolicy++
		}
	}

	// Return the number of passwords that adhere
	println(adhereToPasswordPolicy)
}
