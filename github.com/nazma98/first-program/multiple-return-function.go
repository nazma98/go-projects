package main

import "fmt"

func getResult(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	mul := num1 * num2

	return sum, mul
}

func main() {
	a := 10
	b := 20

	sum, mul := getResult(a, b)

	fmt.Println(sum)
	fmt.Println(mul)
}