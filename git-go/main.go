package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	// slice expression to just graba the items after the executable path
	args := os.Args[1:]

	// fullMessage := ""

	// gitType := commitType(args[0])
	// gitScope := commitScope(args[1])
	// gitMessage := args[2]

	for i, item := range args {
		switch len(args[i]) {
		case 0:
			println("Nothing present for:", args[i])
		case 1:
			println("Nothing present for:", args[i])
		case 3:
			println("Nothing present for:", args[i])
		}

		fmt.Println("nah", item)
	}
}

func returnArgs(args []string) bool {
	if len(args) == 0 {
		return false
	}
	return true
}

func commitType(arg string) string {

	// list of items that should be available
	// "fix", "feat", "refactor", "docs", "test" for now

	argTypes := []string{"fix", "feat", "refactor", "docs", "test"}

	if slices.Contains(argTypes, arg) {
		fmt.Println("arg matches type:", arg)
		return arg
	}

	fmt.Println("No arg match existing types")
	return ""
}

func commitScope(arg string) string {

	if len(arg) > 0 {
		return arg
	}
	return ""
}

func commitMessage(arg string) {}
