package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"february",
		"march",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"december",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // capacity

	// jika merubah array, slice akan berubah juga
	// begitu juga sebaliknya
	// misal :
	// months[5] = "diubah"
	// fmt.Println(slice1) // [mei, diubah, juli]
}
