package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name -- ")
	fmt.Scanln(&name)
	return name
}

func getNumber() (int, int) {
	var num1 int 
	var num2 int
	fmt.Println("Enter first number - ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number - ")
	fmt.Scanln(&num2)

	return num1, num2
}

func greetUser(name string) {
	fmt.Println("Hello, ", name)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printSum(sum int) {
	fmt.Println("The sum is - ", sum)
}

func thankYouMessage() {
	fmt.Println("Thanks for using our applications!")
	fmt.Println("Goodbye")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getNumber()
	sum := add(num1, num2)
	greetUser(name)
	printSum(sum)
	thankYouMessage()
}