package main

import "fmt"

type Customer struct {
	Name    string
	Age     int
	married bool
}

func main() {
	Rikhi := Customer{
		Name:    "Rikhi",
		Age:     20,
		married: false,
	}
	fmt.Println(Rikhi)
}
