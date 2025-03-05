package main

import "fmt"

type Student struct{
	ID int
	Name string
	Class string
	Marks int
}

func main() {
	var student1 Student

	student1 = Student {
		ID: 12,
		Name: "Aghd",
		Class: "Four",
		Marks: 89,
	}

	fmt.Println("Name of student : ", student1.Name)
	fmt.Println("ID : ", student1.ID)
	fmt.Println("Class : ", student1.Class)
	fmt.Println("Marks : ", student1.Marks)

}