package main

import (
	"fmt"
	"os"
)

func main() {

	// Addition of Person struct
	person := Person{
		Name: "Paul",
		Hair: "red",
		Age:  24,
	}

	person.birthday()

	person.displayPerson()

	// slice expression to just graba the items after the executable path
	args := os.Args[1:]
	displayArgs(args)

	changeName(&person, "jacob")
	changeHair(&person, "periwinkle")
	changeAge(&person, 32)

	fmt.Println(person)

}

func displayArgs(args []string) {

	if len(args) == 0 {
		fmt.Println("No args present")
		return
	}

	fmt.Println("args are ... ", args)

}
