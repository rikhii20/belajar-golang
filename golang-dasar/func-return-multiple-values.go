package main

import "fmt"

func fullName() (string, string) {
	return "Rikhi", "Sobari"
}

func main() {
	// firstName, lastName := fullName()
	// fmt.Println(firstName, lastName)

	// ignore data yg tidak ingin ditampilkan (_)
	firstName, _ := fullName()
	fmt.Println(firstName)
}
