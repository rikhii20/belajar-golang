package main

import "fmt"

func sum(number ...int) int {
	count := 0
	for _, value := range number {
		count += value
	}
	return count
}

func main() {
	total := sum(10, 10, 10)
	fmt.Println(total)

	// passing data slice ke variabel argument
	slice := []int{10, 20, 30, 40}
	total = sum(slice...)
	fmt.Println(total)
}
