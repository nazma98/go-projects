package main

import "fmt"

type User struct{
	Name string
	Age int
}

func main() {
	var user1 User

	user1 = User {
		Name : "Nazma",
		Age : 20,
	}

	fmt.Println("Name : ", user1.Name)
	fmt.Println("Age : ", user1.Age)
}