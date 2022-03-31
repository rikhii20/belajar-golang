package main

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func (student Student) sayHai() {
	fmt.Println("hai saya", student.Name, "kelas", student.Grade)
}

func main() {
	rikhi := Student{
		Name:  "Rikhi",
		Grade: 7,
	}

	rikhi.sayHai()
}
