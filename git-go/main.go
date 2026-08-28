package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// saving args
	args := os.Args[1:]
	fmt.Println(filepath.Base(os.Args[0]))
	fmt.Println(args[0])
	fmt.Println(args[1])

}
