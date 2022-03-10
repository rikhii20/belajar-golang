package main

import "fmt"

func main() {
	name := "Tommy maartanto"

	if name == "Rikhi" {
		fmt.Println("Hello", name)
	} else if name == "Tommy" {
		fmt.Println("Hello lagi", name)
	} else {
		fmt.Println("Kenalang dong", name)
	}

	if len(name) > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Sudah pas")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang short statement")
	} else {
		fmt.Println("Sudah pas short statement")
	}
}
