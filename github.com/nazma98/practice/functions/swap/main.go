package main

import "fmt"

func swap(num1, num2 int) (int, int) {
	num1, num2 = num2, num1
	return num1, num2
}

func main() {
	num1 := 4
	num2 := 8
	num1, num2 = swap(num1, num2)
	fmt.Println(num1, ", ", num2)
}
