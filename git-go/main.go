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

	var filteredItem CandidateScore
	var containsPotentialCandidate bool

	argTypes := []string{"fix", "feat", "refactor", "docs", "test"}

	matchingStrings := stringSimilarityScore(argTypes, arg)
	potentialScoredMatches, hasMatches := bestPossibleMatch(matchingStrings)

	fmt.Println("match info: ", potentialScoredMatches, hasMatches)

	fmt.Println(containsPotentialCandidate)
	// fmt.Println("matched: ", matchingStrings)

	if slices.Contains(argTypes, arg) {
		return arg, true
	}

	if len(matchingStrings) > 0 {
		filteredItem, containsPotentialCandidate = bestPossibleMatch(matchingStrings)
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

		finalScore := score + additionalScore

		possibleMatches = append(possibleMatches, CandidateScore{
			Word:  arrayArgument,
			Score: finalScore,
		})

	}

	return possibleMatches
}

func bestPossibleMatch(matches []CandidateScore) (CandidateScore, bool) {

	errCount := 0
	var noMatches bool

	if errCount == len(matches) {
		noMatches = true
	} else {
		noMatches = false
	}

	for _, match := range matches {
		if match.Score == 0 {
			errCount++
		}
	}

	if noMatches {
		return CandidateScore{}, false
	}

	slices.SortFunc(matches, compareScores)
	return matches[0], true
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

func compareScores(a, b CandidateScore) int {
	return cmp.Compare(b.Score, a.Score)
}

func returnString(arg string) string {

	if len(arg) > 0 {
		return arg
	}
	return ""
}

func displayHelpOptions() string {
	return ""
}
