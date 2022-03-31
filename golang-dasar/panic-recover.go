package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("error occured", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("jalan terus")
}

// recover untuk menangkap panicnya dan code terus dieksekusi
// simpan recover di defer func
