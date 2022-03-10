package main

import "fmt"

// memasukkan function ke dalam variable sebagai value
func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Rikhi"))
}
