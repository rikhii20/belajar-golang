package main

import "fmt"

func main() {
	const firstName string = "Rikhi"
	const lastName = "Sobari"

	fmt.Println(firstName, lastName)

	// deklarasi multiple constant
	const (
		age    = 25
		height = 180
		weight = 72
	)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(weight)
}
