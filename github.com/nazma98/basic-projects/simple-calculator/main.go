package main

import "fmt"

func getNumbers() (int, int) {
	fmt.Println("Enter first number")
	var firstNumber int
	fmt.Scanln(&firstNumber)
	fmt.Println("Enter second number")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	return firstNumber, secondNumber
}

func add(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func subtract(firstNumber, secondNumber int) int {
	return firstNumber - secondNumber
}

func multiply(firstNumber, secondNumber int) int {
	return firstNumber * secondNumber
}

func divide(firstNumber, secondNumber int) int {
	if secondNumber == 0 {
		fmt.Println("Can't be divided by zero!")
		getNumbers()
	} else {
		return firstNumber / secondNumber
	}
}

func main() {
	firstNumber, secondNumber := getNumbers()
	sum := divide(firstNumber, secondNumber)
	fmt.Println(sum)
}