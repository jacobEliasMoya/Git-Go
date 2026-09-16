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
	gitType := args[0]

	// gitScope := args[1]
	// gitMessage := args[2]

	if hasArgs {
		fmt.Println("Has Args: ", len(args))
		commitType(gitType)
	} else {
		fmt.Println("No Args")
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
