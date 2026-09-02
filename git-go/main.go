package main

import (
	"fmt"
	"os"
)

func main() {
	// saving args
	args := os.Args[1:]
	osStuff, err := os.ReadFile("progress.md")

	if err != nil {
		fmt.Println("This is the error %w",err)
	}

	fmt.Printf("osStuff: %v\n", osStuff)

	if len(args) != 1 {
		fmt.Println("Review: Multiple Args Present")
		return 
	}

	printArgs(args[0])
}
