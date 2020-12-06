package main

import (
	"bufio"
	"os"

	"github.com/fatih/set"
)

var totalQuestionsYes int

// List of sets, one set for every person
var yesQuestionsPerPerson []set.Interface

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Newline, so look for the number of questions everyone voted yes for
		if len(line) == 0 {
			tallyUpVotes(yesQuestionsPerPerson)
		} else {
			// Line with votes, add every vote to a new set which we add to the list
			yesAnswers := set.New(set.NonThreadSafe)
			for _, c := range line {
				yesAnswers.Add(c)
			}
			yesQuestionsPerPerson = append(yesQuestionsPerPerson, yesAnswers)
		}
	}

	tallyUpVotes(yesQuestionsPerPerson)

	println(totalQuestionsYes)
}

func tallyUpVotes(qy []set.Interface) {
	allQuestions := set.New(set.NonThreadSafe)
	allQuestions.Add('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')

	for _, s := range qy {
		allQuestions = set.Intersection(allQuestions, s)
	}

	totalQuestionsYes += allQuestions.Size()
	allQuestions.Clear()
	yesQuestionsPerPerson = nil
}
