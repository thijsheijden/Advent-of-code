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
	letter        byte
	occurrenceOne int
	occurrenceTwo int
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
		oneTwo := strings.Split(occurrenceString, "-")
		one, err := strconv.Atoi(oneTwo[0])
		two, err2 := strconv.Atoi(oneTwo[1])

		if err != nil || err2 != nil {
			log.Fatalln(err)
			log.Fatalln(err2)
		}

		// Get the policy letter
		policyLetterSplit := strings.Split(policyLetterString, ":")[0]
		policyLetter := []byte(policyLetterSplit)[0]

		// Create the password objects
		passPolicy := passwordPolicy{occurrenceOne: one, occurrenceTwo: two, letter: policyLetter}
		pass := password{policy: passPolicy, password: passwordString}

		// Add to the list of passwords
		passwords = append(passwords, pass)
	}

	// Go over all passwords and determine how many of them adhere to their password policy
	var adhereToPasswordPolicy int = 0
	for _, pass := range passwords {
		// Check if the policy letter does not appear at position one AND two
		if !(pass.password[pass.policy.occurrenceOne-1] == pass.policy.letter && pass.password[pass.policy.occurrenceTwo-1] == pass.policy.letter) && (pass.password[pass.policy.occurrenceOne-1] == pass.policy.letter || pass.password[pass.policy.occurrenceTwo-1] == pass.policy.letter) {
			adhereToPasswordPolicy++
		}
	}

	// Return the number of passwords that adhere
	println(adhereToPasswordPolicy)
}
