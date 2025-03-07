package main

import "fmt"

type Student struct{
	ID int
	Name string
	Class string
}

func (std Student) printDetails() {
	fmt.Println("Name : ", std.Name)
	fmt.Println("Class : ", std.Class)
}

func main() {
	var student1 Student

	student1 = Student{
		ID: 12,
		Name: "Shimul",
		Class: "Six",
	}

	student1.printDetails()
}