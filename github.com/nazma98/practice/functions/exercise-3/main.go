package main

import "fmt"

func getNumbers(num1, num2 int) (int, int) {
	sum := num1 + num2
	product := num1 * num2
	return sum, product
}

func main() {
	a := 10
	b := 15
	sum, product := getNumbers(a, b)
	fmt.Println(sum, " ", product)
}
