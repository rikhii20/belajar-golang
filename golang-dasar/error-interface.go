package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak bisa nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 20)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}
