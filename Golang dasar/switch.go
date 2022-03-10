package main

import "fmt"

func main() {
	name := "Rikhi"

	switch name {
	case "Rikhi":
		fmt.Println("Hello", name)
	case "Joko":
		fmt.Println("Hello", name)
	default:
		fmt.Println("kenalan dulu dong", name)
	}

	switch len(name) > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah pas")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang short statement")
	case false:
		fmt.Println("nama sudah pas short statement")
	}

	// bisa tanpa nama variabel
	switch {
	case len(name) > 10:
		fmt.Println("nama kepanjangan")
	case len(name) > 5:
		fmt.Println("nama kepanjangan")
	default:
		fmt.Println("nama sudah benar")
	}
}
