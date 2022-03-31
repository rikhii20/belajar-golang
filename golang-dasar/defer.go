package main

import "fmt"

func logging() {
	fmt.Println(("selesai memanggil function"))
}

func runApps(value int) {
	defer logging()
	fmt.Println("run apps")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApps(5)
}
