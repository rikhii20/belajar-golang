package main

import "fmt"

func main() {
	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke -", i)
	}

	fmt.Println("==================")

	// continue
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke -", j)
	}

}
