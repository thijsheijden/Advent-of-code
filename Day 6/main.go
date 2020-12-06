package main

import (
	"bufio"
	"os"

	"github.com/fatih/set"
)

var totalQuestionsYes int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	questions := set.New(set.NonThreadSafe)
	questions.Add('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')

	questionsYesInThisGroup := set.New(set.NonThreadSafe)

	for scanner.Scan() {
		line := scanner.Text()

		// Newline, so tally up the number of yes votes
		if len(line) == 0 {
			tallyUpVotes(questions, &questionsYesInThisGroup)
			questionsYesInThisGroup.Clear()
		} else {
			// Line with votes, add them to total pool
			for _, c := range line {
				questionsYesInThisGroup.Add(c)
			}
		}
	}

	tallyUpVotes(questions, &questionsYesInThisGroup)
	questionsYesInThisGroup.Clear()

	println(totalQuestionsYes)
}

func tallyUpVotes(q set.Interface, qy *set.Interface) int {
	questionsNo := set.Difference(q, *qy)
	totalQuestionsYes += 26 - questionsNo.Size()
	return totalQuestionsYes
}
