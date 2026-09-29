package main

import (
	"fmt"
	"os"
	"slices"
)

type CandidateScore struct {
	Word  string
	Score int
}

func main() {
	// slice expression to just graba the items after the executable path
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Error: Missing args <sc>")
	}

	var fullMessage, gitType, gitScope, gitMessage string

	for i, item := range args {
		switch i {
		case 0:
			hasGitType, isGitType := commitType(item)

			if isGitType {
				gitType += hasGitType
			} else {
				return
			}

		case 1:
			gitScope += returnString(item)
		case 2:
			gitMessage += ": " + returnString(item)
		}
	}

	fullMessage += fmt.Sprintf("%s (%s)%s", gitType, gitScope, gitMessage)

	fmt.Println(fullMessage)
}

func returnArgs(args []string) bool {
	if len(args) == 0 {
		return false
	}
	return true
}

func commitType(arg string) (string, bool) {

	// list of items that should be available
	// "fix", "feat", "refactor", "docs", "test", for now

	argTypes := []string{"fix", "feat", "refactor", "docs", "test"}

	matchingStrings := stringSimilarityScore(argTypes, arg)

	if slices.Contains(argTypes, arg) {
		return arg, true
	}

	fmt.Println("No arg match existing types")
	return "", false
}

func stringSimilarityScore(arr []string, arg string) []CandidateScore {

	var possibleMatches []CandidateScore

	for _, arrayArgument := range arr {

		var arrChar, argChar []rune

		for _, arrayChar := range arrayArgument {
			arrChar = append(arrChar, arrayChar)
		}

		for _, char := range arg {
			argChar = append(argChar, char)
		}

		score := 0

		for i, char := range arrChar {

			if i >= len(argChar) {
				break
			}

			if char == argChar[i] {
				score++
			}

		}

		if score > 0 {
			possibleMatches = append(possibleMatches, CandidateScore{
				Word:  arrayArgument,
				Score: score,
			})
		}

	}

	return possibleMatches
}

func returnString(arg string) string {

	if len(arg) > 0 {
		return arg
	}
	return ""
}

func bestPossibleMatch(matches []CandidateScore) string {

}
