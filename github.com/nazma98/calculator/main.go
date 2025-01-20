package main

import "fmt"

func getInput() (int, string, int) {
	fmt.Println("Enter first number : ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("Enter operator : ")
	var operator string
	fmt.Scanln(&operator)
	fmt.Println("Enter second number : ")
	var num2 int
	fmt.Scanln(&num2)

	return num1, operator, num2
}

func calculate(num1 int, operator string, num2 int) int {
	var result int
	if operator == "+" {
		result = (num1 + num2)
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "*" {
		result = num1 * num2
	} else if operator == "/" {
		result = num1 / num2
	} else if operator == "%" {
		result = num1 % num2
	}
	return result
}

func main() {
	num1, operator, num2 := getInput() 
	result := calculate(num1, operator, num2)
	fmt.Println(result)
}