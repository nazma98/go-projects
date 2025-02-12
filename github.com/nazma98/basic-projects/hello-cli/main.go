package main

import "fmt"

func userName() string {
	fmt.Println("Enter you name : ")
	var name string
	fmt.Scanln(&name)
	return name
}

func greet(name string) {
	fmt.Println("Hello ", name, ", Welcome to learning Go!\n")
}
func main() {
	name := userName()
	greet(name)
}