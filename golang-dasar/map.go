package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rikhi",
		"address": "Kudus",
	}

	//menambah data
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Golang"
	book["author"] = "Rikhi"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
