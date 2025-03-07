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

}