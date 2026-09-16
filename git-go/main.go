package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	// slice expression to just graba the items after the executable path
	args := os.Args[1:]
	hasArgs := returnArgs(args)

	// fullMessage := ""

	// gitType := args[0]
	// gitScope := args[1]
	// gitMessage := args[2]

	switch len(args) {
	case 1:
		fmt.Println(args[0])
	case 2:
		fmt.Println(args[0])
	case 3:
		fmt.Println(args[0])
	default:
		fmt.Println("bing bing")
	}

	if hasArgs {
	} else {
	}
}

func returnArgs(args []string) bool {
	if len(args) == 0 {
		return false
	}
	return true
}

func commitType(arg string) (string, bool) {

	// list of items that should be available
	// "fix", "feat", "refactor", "docs", "test" for now

	argTypes := []string{"fix", "feat", "refactor", "docs", "test"}

	if slices.Contains(argTypes, arg) {
		fmt.Println("arg matches type:", arg)
		return arg, true
	}

	fmt.Println("No arg matche existing types")
	return "", false
}

func commitScope(arg string) {

}

func commitMessage(arg string) {}
