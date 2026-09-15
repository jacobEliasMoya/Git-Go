package main

import (
	"fmt"
	"os"
)

func main() {
	// slice expression to just graba the items after the executable path
	args := os.Args[1:]
	hasArgs := returnArgs(args)

	if hasArgs {
		fmt.Println("Has Args:")
		commitType(args[0])
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
	argTypes := []string{"fix", "feat", "refactor", "docs", "test"}

	for _, t := range argTypes {
		if t == arg {
			fmt.Println("arg matches type:", arg)
			return arg, true
		}
	}

	fmt.Println("No arg matche existing types")
	return "", false
}

func commitScope(arg string) {

}

func commitMessage(arg string) {}
