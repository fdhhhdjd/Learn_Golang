package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Printf("%s is eating.\n", a.Name)
}

type Dog struct {
	Animal
	Breed string
	IsPet bool
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
		IsPet:  true,
	}

	d.Eat()
}
