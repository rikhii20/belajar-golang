package main

import "fmt"

func main() {
	// menyebutkan tipe datanya dulu
	var name string

	name = "Rikhi Sobari"
	fmt.Println(name)

	name = "Rika Sobari"
	fmt.Println(name)

	// lgsg assign ke variablenya, otomatiis go akan detect data types nya
	var friend = "Budi"
	fmt.Println(friend)

	// tanpa deklarasi kata var
	negara := "Indonesia"
	fmt.Println(negara)

	// multiple variable
	var (
		firstName = "Rikhi"
		lastName  = "Sobari"
	)

	fmt.Println(firstName, lastName)

}
