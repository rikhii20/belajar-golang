package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 5
	var b = 10
	var c = a + b
	fmt.Println(c)

	// augmented operator
	var d = 10
	d += 10 // d = d + 10
	fmt.Println(d)

	// unary operator
	var e = 5
	e++ // e = e + 1
	fmt.Println(e)

	var negative = -100
	var positive = 100
	fmt.Println(negative)
	fmt.Println(positive)

}
