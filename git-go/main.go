package main

import (
	"fmt"
	"os"
	"slices"
)

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

	stringSimilarityScore(argTypes, arg)

	if slices.Contains(argTypes, arg) {
		return arg, true
	}

	fmt.Println("No arg match existing types")
	return "", false
}

func stringSimilarityScore(arr []string, arg string) int {

	for _, arrayArgument := range arr {

		var arrChar, argChar []rune

		for _, arrayChar := range arrayArgument {
			arrChar = append(arrChar, arrayChar)
		}

		for _, char := range arg {
			argChar = append(argChar, char)
		}

		if slices.Equal(arrChar, argChar) {
			fmt.Printf("Matching: %s \n", string(argChar))
		} 



	}

	return 0
}

func returnString(arg string) string {

	if len(arg) > 0 {
		return arg
	}
	return ""
}
