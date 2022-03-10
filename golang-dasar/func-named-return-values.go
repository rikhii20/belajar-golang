package main

import "fmt"

// deklarasi variable di awal setelah parameter
func completeName() (firstName, middleName, lastName string) {
	firstName = "Rikhi"
	middleName = "Martin"
	lastName = "Sobari"

	return
}

func main() {
	namaDepan, namaTengah, NamaAkhir := completeName()
	fmt.Println(namaDepan, namaTengah, NamaAkhir)
}
