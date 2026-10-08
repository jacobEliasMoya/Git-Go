package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

type CandidateScore struct {
	Word  string
	Score int
}

func main() {
	// slice expression to just graba the items after the executable path
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Error: Missing args")
		os.Exit(1)
	}

	var fullMessage, gitType, gitScope, gitMessage string

	for i, item := range args {
		switch i {
		case 0:
			hasGitType, isGitType := commitType(item)

			if isGitType {
				gitType += hasGitType
			} else {
				os.Exit(0)
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

	argTypes := []string{"fix", "feat", "refactor", "docs", "test"}
	var filteredItem CandidateScore

	matchingStrings := stringSimilarityScore(argTypes, arg)

	if slices.Contains(argTypes, arg) {
		return arg, true
	}

	if len(matchingStrings) > 0 {
		filteredItem = bestPossibleMatch(matchingStrings)
		fmt.Printf("No Matching commit types, did you mean `%s`", filteredItem.Word)
	} else {
		fmt.Printf("No Matching commit types \nTry: fix, feat, refactor, docs, test")
	}

	return "", false
}

func stringSimilarityScore(arr []string, arg string) []CandidateScore {

	var possibleMatches []CandidateScore

	for _, arrayArgument := range arr {

		var arrChar, argChar []rune

		additionalScore := compareTotalChars(arrayArgument, arg)

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
				Score: score + additionalScore,
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

func compareScores(a, b CandidateScore) int {
	return cmp.Compare(b.Score, a.Score)
}

func bestPossibleMatch(matches []CandidateScore) CandidateScore {
	slices.SortFunc(matches, compareScores)
	return matches[0]
}

func compareTotalChars(possibleType string, commitType string) int {

	score := 0

	// take one string, parse chars, make map of that arr / build back into word for checks.
	possibleTypeMap := make(map[rune]struct{})

	for _, chars := range possibleType {
		possibleTypeMap[chars] = struct{}{}
	}

	for _, types := range commitType {
		if strings.ContainsRune(possibleType, types) {
			score++
		}
	}

	return score
}
