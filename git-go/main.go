package main

import (
	"fmt"
	"os"
)

func main() {
	// slice expression to just graba the items after the executable path
	args:= os.Args[1:]
	fmt.Println("Main git-go ",args)
}
