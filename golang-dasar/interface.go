package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHai(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	eko := Person{
		Name: "Rikhi",
	}
	sayHai(eko)
}
