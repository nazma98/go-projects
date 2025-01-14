package main

import "fmt"

func printSomething() {
	fmt.Println("Education must be free & available!")
}

func sayHello(name string) {
	fmt.Println("Welcome to Golang course, ", name);
}

func main() {
	printSomething()
	sayHello("Nazma")
}