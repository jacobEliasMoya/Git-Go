package main

import (
	"fmt"
	"os"
)

func main() {
	// saving args
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Review: Multiple Args Present")
		return
	}
	printArgs(args[0])
}


