package main

import "fmt"

func main() {
	// for dengan statement :
	for count := 1; count <= 10; count++ {
		fmt.Println("perulangan ke -", count)
	}
	// init statement = count := 1, sebagai inisialisasi variabel
	// post statement = count++, sebagai kondisi yg akan selalu dieksekusi sampai perulangan selesai

	// =======================================

	// iterasi data di slice
	slice := []string{"Rikhi", "Rika", "Rian"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// iterasi dengan for range
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	// for range tanpa index
	for _, value := range slice {
		fmt.Println(value)
	}

	// ========================================

	// iterasi data di map
	person := make(map[string]string)
	person["name"] = "Rikhi"
	person["title"] = "Backend Developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
