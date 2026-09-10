package main

import "fmt"

type Person struct {
	Name string
	Hair string
	Age  int
}

func (p Person) displayPerson() {
	fmt.Println("Name: ", p.Name)
	fmt.Println("Hair: ", p.Hair)
	fmt.Println("Age: ", p.Age)
}

func (p *Person) birthday() {
	p.Age ++
}

func (p Person) celebrate(results chan<- Person) {
	// working here
}
