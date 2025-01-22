package main

import "fmt"

func getNumbers() (int, int){
	var num1 int
	var num2 int
	fmt.Println("Enter first number : ")
	fmt.Scanln(&num1)
	fmt.Println("Enter first number : ")
	fmt.Scanln(&num2)
	return num1, num2
}

func findLargeNumber(num1 int, num2 int) int {
	max := num1
	if num2 > max {
		max = num2
	}
	return max
}

func main() {
	num1, num2 := getNumbers()
	max := findLargeNumber(num1, num2)
	fmt.Println("Large number : ", max)
}