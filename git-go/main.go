package main

import (
	"fmt"
	"os"
)

func main() {
	// slice expression to just graba the items after the executable path
	args := os.Args[1:]
	displayArgs(args)
}

func displayArgs(args []string) {

	if len(args) == 0 {
		fmt.Println("No args present")
	} else {
		fmt.Println("args are ... ", args)
	}

}
