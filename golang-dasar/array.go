package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rikhi"
	names[1] = "Sobari"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung
	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values)) // jumlah panjang array
	values[1] = 10           // mengubah nilai array
	fmt.Println(values)
}
