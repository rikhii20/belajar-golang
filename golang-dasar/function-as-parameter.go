package main

import "fmt"

type Filter func(string) string // alias

func sayHelloFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Rikhi", spamFilter)
	sayHelloFilter("anjing", spamFilter)
}
