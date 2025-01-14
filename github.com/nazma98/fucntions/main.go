package main

import "fmt"

func printSomething() {
	fmt.Println("Education must be free & available!")
}

func sayHello(name string) {
	fmt.Println("Welcome to Golang course, ", name);
}

func main() {
	// print welcome message
	fmt.Println("Welcome to the application")

	// Get user name as input
	var name string
	fmt.Println("Enter your name -- ")
	fmt.Scanln(&name)

	var num1 int 
	var num2 int
	fmt.Println("Enter first number - ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number - ")
	fmt.Scanln(&num2)
	sum := num1 + num2

	// Display results
	fmt.Println("Hello, ", name)
	fmt.Println("The sum is - ", sum)

	fmt.Println("Thanks for using our applications!")
	fmt.Println("Goodbye")
}