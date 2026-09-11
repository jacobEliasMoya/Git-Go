package main

import (
	"fmt"
	"os"
)

func main() {

	// Addition of Person struct
	paul := Person{
		Name: "paul",
		Hair: "red",
		Age:  24,
	}

	paul.birthday()

	paul.displayPerson()

	// slice expression to just graba the items after the executable path
	args := os.Args[1:]
	displayArgs(args)

	p := &paul

	fmt.Println(p)
}

func displayArgs(args []string) {

	if len(args) == 0 {
		fmt.Println("No args present")
		return
	}

	fmt.Println("args are ... ", args)

}
