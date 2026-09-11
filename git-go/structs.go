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
	p.Age++
}

// coming back to this, somewhat confusing without broader knowledge that should come shortly with my project
func (p Person) celebrate(results chan<- Person) {
	// working here soon
}

func changeName(person *Person, newName string) {
	person.Name = newName
}

func changeHair(person *Person, newHair string) {
	person.Hair = newHair
}

func changeAge(person *Person, newAge int) {
	person.Age = newAge
}
