package main

import "fmt"

func sayHello(name string) string {
	if name == "Rikhi" {
		return "Hello " + name
	} else {
		return "Hello bro"
	}
}

func main() {
	result := sayHello("Rikhi")
	fmt.Println(result)

	fmt.Println(sayHello(""))
}
