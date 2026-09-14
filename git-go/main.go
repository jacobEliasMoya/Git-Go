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
		for _, arg := range args {
			fmt.Println(arg)
		}
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
